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

// Package peercredlistener is deprecated in favor of toolman.org/net/peercred.
package peercredlistener // import "toolman.org/net/peercredlistener"

import (
	"context"
	"net"

	"toolman.org/net/peercred"
)

// ErrAddrInUse is a convenience wrapper around the Posix errno value for
// EADDRINUSE.
// Deprecated: Use package toolman.org/net/peercred instead.
const ErrAddrInUse = peercred.ErrAddrInUse

// PeerCredListener is an implementation of net.Listener that extracts
// the identity (i.e. pid, uid, gid) from the connection's client process.
// This information is then made available through the Ucred member of
// the *PeerCredConn returned by AcceptPeerCred or Accept (after a type
// assertion).
// Deprecated: Use package toolman.org/net/peercred instead.
type PeerCredListener peercred.Listener

// New returns a new PeerCredListener listening on the Unix domain socket addr.
// Deprecated: Use package toolman.org/net/peercred instead.
func New(ctx context.Context, addr string) (*PeerCredListener, error) {
	lis, err := peercred.Listen(ctx, addr)
	return (*PeerCredListener)(lis), err
}

// Accept is a convenience wrapper around AcceptPeerCred allowing
// PeerCredListener callers that utilize net.Listener to function
// as expected. The returned net.Conn is a *PeerCredConn which may
// be accessed through a type assertion. See AcceptPeerCred for
// details on possible error conditions.
//
// Accept contributes to implementing the  net.Listener interface.
// Deprecated: Use package toolman.org/net/peercred instead.
func (pcl *PeerCredListener) Accept() (net.Conn, error) {
	return (*peercred.Listener)(pcl).Accept()
}

// AcceptPeerCred accepts a connection from the receiver's listener
// returning a *PeerCredConn containing the process credentials for
// the client. If the underlying Accept fails or if process credentials
// cannot be extracted, AcceptPeerCred returns nil and an error.
// Deprecated: Use package toolman.org/net/peercred instead.
func (pcl *PeerCredListener) AcceptPeerCred() (*PeerCredConn, error) {
	conn, err := (*peercred.Listener)(pcl).AcceptPeerCred()
	return (*PeerCredConn)(conn), err
}

// PeerCredConn is a net.Conn containing the process credentials for the client
// side of a Unix domain socket connection.
// Deprecated: Use package toolman.org/net/peercred instead.
type PeerCredConn peercred.Conn
