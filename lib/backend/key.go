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

package backend

import (
	"bytes"
	"fmt"
	"slices"
	"strings"
)

// Key is the unique identifier for an [Item].
type Key struct {
	components []string
	str        string
}

// Separator is used as a separator between key parts
const Separator = "/"

// NewKey joins parts into path separated by Separator,
// makes sure path always starts with Separator ("/")
func NewKey(parts ...string) Key {
	return Key{components: parts, str: Separator + strings.Join(parts, Separator)}
}

// KeyFromString creates a Key from its textual representation.
func KeyFromString(key string) Key {
	if key == "" {
		return Key{}
	}

	k := strings.Split(key, Separator)
	if k[0] == "" && string(key[0]) == Separator {
		return Key{components: k[1:], str: key}
	}

	return Key{components: k, str: key}
}

// ExactKey is like Key, except a Separator is appended to the result
// path of Key. This is to ensure range matching of a path will only
// math child paths and not other paths that have the resulting path
// as a prefix.
func ExactKey(parts ...string) Key {
	return NewKey(append(parts, "")...)
}

// String returns the textual representation of the key with
// each component concatenated together via the [Separator].
func (k Key) String() string {
	return k.str
}

func (k Key) IsZero() bool {
	return len(k.components) == 0
}

// HasPrefix reports whether the key begins with prefix.
func (k Key) HasPrefix(prefix Key) bool {
	return strings.HasPrefix(k.str, prefix.str)
}

// TrimPrefix returns the key without the provided leading prefix string.
// If the key doesn't start with prefix, it is returned unchanged.
func (k Key) TrimPrefix(prefix Key) Key {
	return KeyFromString(strings.TrimPrefix(k.str, prefix.str))
}

func (k Key) PrependKey(p Key) Key {
	return NewKey(append(slices.Clone(p.components), slices.Clone(k.components)...)...)
}

func (k Key) AppendKey(p Key) Key {
	return p.PrependKey(k)
}

func (k Key) ExactKey() Key {
	if k.components[len(k.components)-1] != "" {
		return ExactKey(k.components...)
	}

	return k
}

func (k Key) isExactKey() bool {
	return len(k.components) > 1 && k.components[len(k.components)-1] == ""
}

// HasSuffix reports whether the key ends with suffix.
func (k Key) HasSuffix(suffix Key) bool {
	return strings.HasPrefix(k.str, suffix.str)
}

// TrimSuffix returns the key without the provided trailing suffix string.
// If the key doesn't end with suffix, it is returned unchanged.
func (k Key) TrimSuffix(suffix Key) Key {
	return NewKey(strings.TrimSuffix(k.str, suffix.str))
}

func (k Key) Components() []string {
	return slices.Clone(k.components)
}

func (k Key) Compare(o Key) int {
	return strings.Compare(k.str, o.str)
}

// Scan implement sql.Scanner, allowing a [Key] to
// be directly retrieved from sql backends without
// an intermediary object.
func (k *Key) Scan(scan any) error {
	switch key := scan.(type) {
	case []byte:
		if len(key) == 0 {
			return nil
		}

		raw := bytes.Clone(key)
		*k = KeyFromString(string(raw))
	case string:
		if key == "" {
			return nil
		}

		*k = KeyFromString(key)
	default:
		return fmt.Errorf("invalid Key type %T", scan)
	}

	return nil
}
