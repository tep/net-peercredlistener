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

// Package peercredlistener provides PeerCredListener and supporting functions.
//
// PeerCredListener is a net.Listener implementation leveraging features of
// Linux based, Unix domain sockets to garner the PID, UID, and GID of the
// client side connection.
//
// PeerCredListener is intended for use cases where a unix-domain server needs
// to reliably identify the process on the client side of each connection. By
// itself, peercredlistener provides support for simple socket connections.
// Additional support for gRPC over unix-domain sockets is available with the
// subordinate package toolman.org/net/peercredlistener/pclcreds.
//
// A simple, unix-domain server can be written similar to the following:
//
//      // Create a new PeerCredListener listening on socketName
//      lsnr, err := peercredlistener.New(ctx, socketName)
//      if err != nil {
//          return err
//      }
//
//      // Wait for and accept an incoming connection
//      conn, err := lsnr.AcceptPeerCred()
//      if err != nil {
//          return err
//      }
//
//      // conn.Ucred has fields Pid, Uid and Gid
//      fmt.Printf("Client PID=%d UID=%d\n", conn.Ucred.Pid, conn.Ucred.Uid)
//
//
// NOTE: This package does not work with IP connections or on operating systems other than Linux.
//
package peercredlistener // import "toolman.org/net/peercredlistener"

import (
	"context"
	"net"
	"os"

	"golang.org/x/sys/unix"
)

// ErrAddrInUse is a convenience wrapper around the Posix errno value for
// EADDRINUSE.
const ErrAddrInUse = unix.EADDRINUSE

// PeerCredListener is an implementation of net.Listener that extracts the
// identity (i.e. pid, uid, gid) from the calling connection. This information
// is available either from the Ucred member of the *PeerCredConn returned by
// AcceptPeerCred.
type PeerCredListener struct {
	net.Listener
}

// New returns a new PeerCredListener listening on the Unix domain socket addr.
func New(ctx context.Context, addr string) (*PeerCredListener, error) {
	lc := new(net.ListenConfig)
	l, err := lc.Listen(ctx, "unix", addr)
	if err != nil {
		return nil, chkAddrInUseError(err)
	}

	return &PeerCredListener{l}, nil
}

// Accept is a convenience wrapper around AcceptPeerCred allowing
// PeerCredListener callers that utilize net.Listener to function
// as expected. The returned net.Conn is a *PeerCredConn which may
// be accessed through a type assertion. See AcceptPeerCred for
// details on possible error conditions.
//
// Accept contributes to implementing the  net.Listener interface.
func (pcl *PeerCredListener) Accept() (net.Conn, error) {
	switch conn, err := pcl.AcceptPeerCred(); err {
	case nil:
		return conn, nil
	default:
		return nil, err
	}
}

// AcceptPeerCred accepts a connection from the receiver's listener
// returning a *PeerCredConn containing the process credentials for
// the client. If the underlying Accept fails or if process credentials
// cannot be extracted, AcceptPeerCred returns nil and an error.
func (pcl *PeerCredListener) AcceptPeerCred() (*PeerCredConn, error) {
	conn, err := pcl.Listener.Accept()
	if err != nil {
		return nil, err
	}

	pcc := &PeerCredConn{Conn: conn}

	uc, ok := conn.(*net.UnixConn)
	if !ok {
		return pcc, nil
	}

	rc, err := uc.SyscallConn()
	if err != nil {
		return nil, err
	}

	var ucred *unix.Ucred
	cerr := rc.Control(func(fd uintptr) {
		ucred, err = unix.GetsockoptUcred(int(fd), unix.SOL_SOCKET, unix.SO_PEERCRED)
	})

	if cerr != nil || err != nil {
		if err == nil {
			err = cerr
		}
		return nil, err
	}

	pcc.Ucred = ucred

	return pcc, nil
}

// PeerCredConn is a net.Conn containing the process credentials for the client
// side of a Unix domain socket connection.
type PeerCredConn struct {
	Ucred *unix.Ucred
	net.Conn
}

func chkAddrInUseError(err error) error {
	operr, ok := err.(*net.OpError)
	if !ok {
		return err
	}

	syserr, ok := operr.Err.(*os.SyscallError)
	if !ok {
		return err
	}

	errno, ok := syserr.Err.(unix.Errno)
	if !ok {
		return err
	}

	if errno != ErrAddrInUse {
		return err
	}

	return errno
}
