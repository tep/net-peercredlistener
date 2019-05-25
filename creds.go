// Copyright 2018 Timothy E. Peoples
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

package peercredlistener

import (
	"context"
	"errors"
	"net"

	"golang.org/x/sys/unix"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

// ErrNoPeer is returned by FromContext if the provided Context contains
// no gRPC peer.
var ErrNoPeer = errors.New("context has no grpc peer")

// ErrNoCredentials is returned by FromContext if the provided Context
// contains no peer process credentials.
var ErrNoCredentials = errors.New("context contains no credentials")

var errNotImplemented = errors.New("not implemented")

// TransportCredentials returns a grpc.ServerOption that exposes the peer
// process credentials (i.e. pid, uid, gid) extracted by a PeerCredListener.
// The peer credentials are available by passing a server method's Context
// to the FromContext function.
func TransportCredentials() grpc.ServerOption {
	return grpc.Creds(&peerCredentials{})
}

// peerCredentials implements the gRPC TransportCredentials interface.
type peerCredentials struct{}

func (pc *peerCredentials) ClientHandshake(context.Context, string, net.Conn) (net.Conn, credentials.AuthInfo, error) {
	return nil, nil, errNotImplemented
}

func (pc *peerCredentials) ServerHandshake(conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	var info credentials.AuthInfo
	if pcConn, ok := conn.(*PeerCredConn); ok {
		info = pcConn.AuthInfo()
	}
	return conn, info, nil
}

func (*peerCredentials) Info() credentials.ProtocolInfo {
	// NOTE: There's little to no documentation on what this struct
	//       should contain but, after a hasty perusal of the code,
	//       it appears that setting SecurityProtocol to a value
	//       unbeknownst to others should be enough to keep gRPC's
	//       guts from trying to initiate a TLS negotiation.
	return credentials.ProtocolInfo{
		SecurityProtocol: "peer",
	}
}

func (pc *peerCredentials) Clone() credentials.TransportCredentials {
	c := *pc
	return &c
}

func (*peerCredentials) OverrideServerName(string) error { return nil }

// Ucred is a wrapper around the Ucred struct from golang.org/x/sys/unix
// allowing it to be used as the AuthInfo member of a gRPC peer.
//
// This is part of the mechanism used for plumbing *Ucred values through
// the gRPC framework and is not intended for general use.
type Ucred unix.Ucred

// AuthType implements the grpc/credentials AuthInfo interface to enable
// plumbing *Ucred values through the gRPC framework.
func (*Ucred) AuthType() string { return "PeerCred" }

// FromContext extracts peer process credentials, if any, from the given
// Context. If the Context has no gRPC peer, ErrNoPeer is returned. If the
// Context's peer is of the wrong type (i.e. contains no peer process
// credentials), ErrNoCredentials will be returned.
func FromContext(ctx context.Context) (*Ucred, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, ErrNoPeer
	}

	c, ok := p.AuthInfo.(*Ucred)
	if !ok {
		return nil, ErrNoCredentials
	}

	return c, nil
}
