// Teleport
// Copyright (C) 2024 Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package vnet

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"sync"

	"github.com/gravitational/trace"
	"golang.org/x/sync/errgroup"

	"github.com/gravitational/teleport"
	vnetproto "github.com/gravitational/teleport/api/gen/proto/go/teleport/vnet/v1"
	"github.com/gravitational/teleport/api/types"
	api "github.com/gravitational/teleport/gen/proto/go/teleport/lib/teleterm/vnet/v1"
	"github.com/gravitational/teleport/lib/client"
	"github.com/gravitational/teleport/lib/teleterm/api/uri"
	"github.com/gravitational/teleport/lib/teleterm/daemon"
	logutils "github.com/gravitational/teleport/lib/utils/log"
	"github.com/gravitational/teleport/lib/vnet"
)

var log = logutils.NewPackageLogger(teleport.ComponentKey, "term:vnet")

type status int

const (
	statusNotRunning status = iota
	statusRunning
	statusClosed
)

// Service implements gRPC service for VNet.
type Service struct {
	api.UnimplementedVnetServiceServer

	cfg    Config
	mu     sync.Mutex
	status status
	// stopErrC is used to pass an error from goroutine that runs VNet in the background to the
	// goroutine which handles RPC for stopping VNet. stopErrC gets closed after VNet stops. Starting
	// VNet creates a new channel and assigns it as stopErrC.
	//
	// It's a buffered channel in case VNet crashes and there's no Stop RPC reading from stopErrC at
	// that moment.
	stopErrC chan error
	// cancel stops the VNet instance running in a separate goroutine.
	cancel context.CancelFunc
}

// New creates an instance of Service.
func New(cfg Config) (*Service, error) {
	if err := cfg.CheckAndSetDefaults(); err != nil {
		return nil, trace.Wrap(err)
	}

	return &Service{
		cfg: cfg,
	}, nil
}

type Config struct {
	DaemonService      *daemon.Service
	ClientStore        *client.Store
	InsecureSkipVerify bool
}

// CheckAndSetDefaults checks and sets the defaults
func (c *Config) CheckAndSetDefaults() error {
	if c.DaemonService == nil {
		return trace.BadParameter("missing DaemonService")
	}

	if c.ClientStore == nil {
		return trace.BadParameter("missing ClientStore")
	}

	return nil
}

func (s *Service) Start(ctx context.Context, req *api.StartRequest) (*api.StartResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.status == statusClosed {
		return nil, trace.CompareFailed("VNet service has been closed")
	}

	if s.status == statusRunning {
		return &api.StartResponse{}, nil
	}

	socket, socketPath, err := vnet.CreateSocket(ctx)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	longCtx, cancelLongCtx := context.WithCancel(context.Background())
	s.cancel = cancelLongCtx
	defer func() {
		// If by the end of this RPC the service is not running, make sure to cancel the long context.
		if s.status != statusRunning {
			cancelLongCtx()
		}
	}()

	g, longCtx := errgroup.WithContext(longCtx)

	g.Go(func() error {
		<-longCtx.Done()

		return trace.Wrap(socket.Close())
	})

	ipv6Prefix, err := vnet.NewIPv6Prefix()
	if err != nil {
		return nil, trace.Wrap(err)
	}
	dnsIPv6 := vnet.IPv6WithSuffix(ipv6Prefix, []byte{2})

	g.Go(func() error {
		return trace.Wrap(vnet.ExecAdminSubcommand(longCtx, socketPath, ipv6Prefix.String(), dnsIPv6.String()))
	})

	appProvider := &appProvider{
		daemonService:      s.cfg.DaemonService,
		clientStore:        s.cfg.ClientStore,
		insecureSkipVerify: s.cfg.InsecureSkipVerify,
	}

	ns, err := vnet.Setup(ctx, appProvider, socket, ipv6Prefix, dnsIPv6)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	g.Go(func() error {
		return trace.Wrap(ns.Run(longCtx))
	})

	s.stopErrC = make(chan error, 1)

	go func() {
		err := g.Wait()
		if err != nil && !errors.Is(err, context.Canceled) {
			log.ErrorContext(longCtx, "VNet closed with an error", "error", err)
			s.stopErrC <- err
		}
		close(s.stopErrC)

		// TODO(ravicious): Notify the Electron app about change of VNet state, but only if it's
		// running. If it's not running, then the Start RPC has already failed and forwarded the error
		// to the user.

		s.mu.Lock()
		defer s.mu.Unlock()

		s.status = statusNotRunning
	}()

	s.status = statusRunning
	return &api.StartResponse{}, nil
}

// Stop stops VNet and cleans up used resources. Blocks until VNet stops or ctx is canceled.
func (s *Service) Stop(ctx context.Context, req *api.StopRequest) (*api.StopResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	errC := make(chan error)

	go func() {
		errC <- trace.Wrap(s.stopLocked())
	}()

	select {
	case <-ctx.Done():
		return nil, trace.Wrap(ctx.Err())
	case err := <-errC:
		if err != nil {
			return nil, trace.Wrap(err)
		}

		return &api.StopResponse{}, nil
	}

}

func (s *Service) stopLocked() error {
	if s.status == statusClosed {
		return trace.CompareFailed("VNet service has been closed")
	}

	if s.status == statusNotRunning {
		return nil
	}

	s.cancel()
	s.status = statusNotRunning

	return trace.Wrap(<-s.stopErrC)
}

// Close stops VNet service and prevents it from being started again. Blocks until VNet stops.
// Intended for cleanup code when tsh daemon gets terminated.
func (s *Service) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.stopLocked()
	s.status = statusClosed

	return trace.Wrap(err)
}

type appProvider struct {
	daemonService      *daemon.Service
	clientStore        *client.Store
	insecureSkipVerify bool
}

func (p *appProvider) ListProfiles() ([]string, error) {
	profiles, err := p.clientStore.ListProfiles()
	return profiles, trace.Wrap(err)
}

func (p *appProvider) GetCachedClient(ctx context.Context, profileName, leafClusterName string) (*client.ClusterClient, error) {
	uri := uri.NewClusterURI(profileName).AppendLeafCluster(leafClusterName)
	client, err := p.daemonService.GetCachedClient(ctx, uri)
	return client, trace.Wrap(err)
}

func (p *appProvider) ReissueAppCert(ctx context.Context, profileName, leafClusterName string, app types.Application) (tls.Certificate, error) {
	clusterURI := uri.NewClusterURI(profileName).AppendLeafCluster(leafClusterName)
	cluster, _, err := p.daemonService.ResolveClusterURI(clusterURI)
	if err != nil {
		return tls.Certificate{}, trace.Wrap(err)
	}

	client, err := p.daemonService.GetCachedClient(ctx, clusterURI)
	if err != nil {
		return tls.Certificate{}, trace.Wrap(err)
	}

	// TODO(ravicious): Copy stuff from DaemonService.reissueGatewayCerts in order to handle expired certs.
	cert, err := cluster.ReissueAppCert(ctx, client, app)
	return cert, trace.Wrap(err)
}

// GetDialOptions returns ALPN dial options for the profile.
func (p *appProvider) GetDialOptions(ctx context.Context, profileName string) (*vnet.DialOptions, error) {
	profile, err := p.clientStore.GetProfile(profileName)
	if err != nil {
		return nil, trace.Wrap(err, "loading user profile")
	}
	dialOpts := &vnet.DialOptions{
		WebProxyAddr:            profile.WebProxyAddr,
		ALPNConnUpgradeRequired: profile.TLSRoutingConnUpgradeRequired,
		InsecureSkipVerify:      p.insecureSkipVerify,
	}
	if dialOpts.ALPNConnUpgradeRequired {
		dialOpts.RootClusterCACertPool, err = p.getRootClusterCACertPool(ctx, profileName)
		if err != nil {
			return nil, trace.Wrap(err)
		}
	}
	return dialOpts, nil
}

func (p *appProvider) GetVnetConfig(ctx context.Context, profileName, leafClusterName string) (*vnetproto.VnetConfig, error) {
	clusterClient, err := p.GetCachedClient(ctx, profileName, leafClusterName)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	vnetConfigClient := clusterClient.AuthClient.VnetConfigServiceClient()
	vnetConfig, err := vnetConfigClient.GetVnetConfig(ctx, &vnetproto.GetVnetConfigRequest{})
	return vnetConfig, trace.Wrap(err)
}

// getRootClusterCACertPool returns a certificate pool for the root cluster of the given profile.
func (p *appProvider) getRootClusterCACertPool(ctx context.Context, profileName string) (*x509.CertPool, error) {
	tc, err := p.newTeleportClient(ctx, profileName, "")
	if err != nil {
		return nil, trace.Wrap(err, "creating new client")
	}
	certPool, err := tc.RootClusterCACertPool(ctx)
	if err != nil {
		return nil, trace.Wrap(err, "loading root cluster CA cert pool")
	}
	return certPool, nil
}

func (p *appProvider) newTeleportClient(ctx context.Context, profileName, leafClusterName string) (*client.TeleportClient, error) {
	cfg := &client.Config{
		ClientStore: p.clientStore,
	}
	if err := cfg.LoadProfile(p.clientStore, profileName); err != nil {
		return nil, trace.Wrap(err, "loading client profile")
	}
	if leafClusterName != "" {
		cfg.SiteName = leafClusterName
	}
	tc, err := client.NewClient(cfg)
	if err != nil {
		return nil, trace.Wrap(err, "creating new client")
	}
	return tc, nil
}