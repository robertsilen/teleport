/*
Copyright 2023 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package generic

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/gravitational/trace"

	"github.com/gravitational/teleport/lib/backend"
	"github.com/gravitational/teleport/lib/utils"
)

// nonceProtectedResourceShim is a helper for quickly extracting the nonce
type nonceProtectedResourceShim struct {
	Nonce uint64 `json:"nonce"`
}

// ErrNonceViolation is the error returned by FastUpdateNonceProtectedResource when a nonce-protected
// update fails due to concurrent modification. This error should be caught and re-mapped into an
// appropriate user-facing message for the given resource type.
var ErrNonceViolation = errors.New("nonce-violation")

// nonceProtectedResource describes the expected methods for a resource that is protected
// from concurrent modification by a nonce.
type nonceProtectedResource interface {
	Expiry() time.Time
	GetNonce() uint64
	WithNonce(uint64) any
}

// FastUpdateNonceProtectedResource is a helper for updating a resource that is protected by a nonce. The target resource must store
// its nonce value in a top-level 'nonce' field in order for correct nonce semantics to be observed.
func FastUpdateNonceProtectedResource[T nonceProtectedResource](ctx context.Context, bk backend.Backend, key backend.Key, resource T) error {
	if resource.GetNonce() == math.MaxUint64 {
		return fastUpsertNonceProtectedResource(ctx, bk, key, resource)
	}

	val, err := utils.FastMarshal(resource.WithNonce(resource.GetNonce() + 1))
	if err != nil {
		return trace.Errorf("failed to marshal resource at %q: %v", key, err)
	}
	item := backend.Item{
		Key:     key,
		Value:   val,
		Expires: resource.Expiry(),
	}

	if resource.GetNonce() == 0 {
		_, err := bk.Create(ctx, item)
		if err != nil {
			if trace.IsAlreadyExists(err) {
				return ErrNonceViolation
			}
			return trace.Wrap(err)
		}

		return nil
	}

	prev, err := bk.Get(ctx, item.Key)
	if err != nil {
		if trace.IsNotFound(err) {
			return ErrNonceViolation
		}
		return trace.Wrap(err)
	}

	var shim nonceProtectedResourceShim
	if err := utils.FastUnmarshal(prev.Value, &shim); err != nil {
		return trace.Errorf("failed to read nonce of resource at %q", item.Key)
	}

	if shim.Nonce != resource.GetNonce() {
		return ErrNonceViolation
	}

	_, err = bk.CompareAndSwap(ctx, *prev, item)
	if err != nil {
		if trace.IsCompareFailed(err) {
			return ErrNonceViolation
		}

		return trace.Wrap(err)
	}

	return nil
}

// fastUpsertNonceProtectedResource performs an "upsert" while preserving correct nonce ordering. necessary in order to prevent upserts
// from breaking concurrent protected updates.
func fastUpsertNonceProtectedResource[T nonceProtectedResource](ctx context.Context, bk backend.Backend, key backend.Key, resource T) error {
	const maxRetries = 16
	for i := 0; i < maxRetries; i++ {
		prev, err := bk.Get(ctx, key)
		if err != nil && !trace.IsNotFound(err) {
			return trace.Wrap(err)
		}

		var prevNonce uint64
		if prev != nil {
			var shim nonceProtectedResourceShim
			if err := utils.FastUnmarshal(prev.Value, &shim); err != nil {
				return trace.Wrap(err)
			}
			prevNonce = shim.Nonce
		}

		nextNonce := prevNonce + 1
		if nextNonce == 0 {
			nextNonce = 1
		}

		val, err := utils.FastMarshal(resource.WithNonce(nextNonce))
		if err != nil {
			return trace.Errorf("failed to marshal resource at %q: %v", key, err)
		}

		item := backend.Item{
			Key:     key,
			Value:   val,
			Expires: resource.Expiry(),
		}

		if prev == nil {
			_, err := bk.Create(ctx, item)
			if err != nil {
				if trace.IsAlreadyExists(err) {
					continue
				}
				return trace.Wrap(err)
			}

			return nil
		}

		_, err = bk.CompareAndSwap(ctx, *prev, item)
		if err != nil {
			if trace.IsCompareFailed(err) {
				continue
			}

			return trace.Wrap(err)
		}

		return nil
	}

	return trace.LimitExceeded("failed to update resource at %q, too many concurrent updates", key)
}
