package postgres

import (
	"context"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"

	"github.com/gravitational/teleport"
	"github.com/gravitational/trace"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/term"

	clientproto "github.com/gravitational/teleport/api/client/proto"
	"github.com/gravitational/teleport/lib/asciitable"
)

const banner = `Teleport PostgreSQL interactive shell (v%s)
Connected to %q instance as %q user.
Type "help" or \? for help.
`

type REPL interface {
	Close() error
}

type repl struct {
	wg         sync.WaitGroup
	ctx        context.Context
	conn       *pgconn.PgConn
	clientConn io.ReadWriteCloser
	serverConn net.Conn
	route      clientproto.RouteToDatabase
	term       *term.Terminal
}

func New(ctx context.Context, clientConn io.ReadWriteCloser, serverConn net.Conn, route clientproto.RouteToDatabase) (*repl, error) {
	config, err := pgconn.ParseConfig(fmt.Sprintf("postgres://%s@placeholder/%s", route.Username, route.Database))
	if err != nil {
		return nil, trace.Wrap(err)
	}
	config.TLSConfig = nil

	config.DialFunc = func(_ context.Context, _, _ string) (net.Conn, error) {
		return serverConn, nil
	}
	config.LookupFunc = func(_ context.Context, _ string) ([]string, error) {
		return []string{"placeholder"}, nil
	}

	pgConn, err := pgconn.ConnectConfig(ctx, config)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	r := &repl{
		ctx:        ctx,
		conn:       pgConn,
		clientConn: clientConn,
		serverConn: serverConn,
		route:      route,
		term:       term.NewTerminal(clientConn, ""),
	}

	r.wg.Add(1)
	go r.start()
	return r, nil
}

const lineBreak = "\r\n"

func (r *repl) Close() {
	r.conn.Close(context.TODO())
	r.clientConn.Close()
	r.serverConn.Close()
}

// TODO: add error forwarding.
func (r *repl) Wait() error {
	r.wg.Wait()
	return nil
}

func (r *repl) start() {
	defer r.Close()

	var acc strings.Builder
	lead := fmt.Sprintf("%s=> ", r.route.Database)
	leadSpacing := strings.Repeat(" ", len(lead))

	if _, err := fmt.Fprintf(r.term, banner, teleport.Version, r.route.ServiceName, r.route.Username); err != nil {
		return
	}
	r.term.SetPrompt(lead)

	for {
		line, err := r.term.ReadLine()
		if err != nil {
			return
		}

		var reply string
		// TODO: cover edge cases
		switch {
		case strings.HasPrefix(line, "\\"):
			args := strings.Split(line, " ")
			switch strings.TrimPrefix(args[0], "\\") {
			case "q":
				return
			}
		case strings.HasSuffix(line, ";"):
			// Execute the query.
			var query string
			if acc.Len() > 0 {
				acc.WriteString(lineBreak + line)
				query = acc.String()
				acc.Reset()
			} else {
				query = line
			}

			r.term.SetPrompt(lead)
			reply = formatResult(r.conn.Exec(r.ctx, query).ReadAll()) + lineBreak
		default:
			// multiline commands.
			acc.WriteString(line)
			r.term.SetPrompt(leadSpacing)
		}

		if len(reply) == 0 {
			continue
		}

		if _, err := r.term.Write([]byte(reply)); err != nil {
			return
		}
	}
}

func formatResult(results []*pgconn.Result, err error) string {
	if err != nil {
		return formatError(err)
	}

	// TODO support multiple queries results.
	// TODO check if multiple queries should be supported.
	res := results[0]

	if !res.CommandTag.Select() {
		return res.CommandTag.String()
	}

	// build columns
	var columns []string
	for _, fd := range res.FieldDescriptions {
		columns = append(columns, fd.Name)
	}

	table := asciitable.MakeTable(columns)
	for _, row := range res.Rows {
		rowData := make([]string, len(columns))
		for i, data := range row {
			rowData[i] = string(data)
		}

		table.AddRow(rowData)
	}

	return table.AsBuffer().String()
}

func formatError(err error) string {
	return "ERR " + err.Error()
}
