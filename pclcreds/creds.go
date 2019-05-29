// Copyright 2019 Timothy E. Peoples
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

// Package pclcreds is deprecated in favor of toolman.org/net/peercred/grpcpeer.
package pclcreds

import (
	"context"

	"golang.org/x/sys/unix"
	"google.golang.org/grpc"

	"toolman.org/net/peercred/grpcpeer"
)

// ErrNoPeer is returned by FromContext if the provided Context contains
// no gRPC peer.
// Deprecated: Use package toolman.org/net/peercred/grpcpeer instead.
var ErrNoPeer = grpcpeer.ErrNoPeer

// ErrNoCredentials is returned by FromContext if the provided Context
// contains no peer process credentials.
// Deprecated: Use package toolman.org/net/peercred/grpcpeer instead.
var ErrNoCredentials = grpcpeer.ErrNoCredentials

// TransportCredentials returns a grpc.ServerOption that exposes the peer
// process credentials (i.e. pid, uid, gid) extracted by a PeerCredListener.
// The peer credentials are available by passing a server method's Context
// to the FromContext function.
// Deprecated: Use package toolman.org/net/peercred/grpcpeer instead.
func TransportCredentials() grpc.ServerOption { return grpcpeer.TransportCredentials() }

// FromContext extracts peer process credentials, if any, from the given
// Context. If the Context has no gRPC peer, ErrNoPeer is returned. If the
// Context's peer is of the wrong type (i.e. contains no peer process
// credentials), ErrNoCredentials will be returned.
// Deprecated: Use package toolman.org/net/peercred/grpcpeer instead.
func FromContext(ctx context.Context) (*unix.Ucred, error) { return grpcpeer.FromContext(ctx) }
