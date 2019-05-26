
# peercredlistener [![Mit License][mit-img]][mit] [![GitHub Release][release-img]][release] [![GoDoc][godoc-img]][godoc] [![Go Report Card][reportcard-img]][reportcard] [![Build Status][travis-img]][travis]

`import "toolman.org/net/peercredlistener"`

* [Install](#pkg-install)
* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-install">Install</a>

```sh
    go get toolman.org/net/peercredlistener
```

## <a name="pkg-overview">Overview</a>
Package peercredlistener provides PeerCredListener and supporting functions.

PeerCredListener is a net.Listener implementation leveraging features of
Linux based, Unix domain sockets to garner the PID, UID, and GID of the
client side connection.

PeerCredListener is intended for use cases where a unix-domain server needs
to reliably identify the process on the client side of each connection. By
itself, peercredlistener provides support for simple socket connections.
Additional support for gRPC over unix-domain sockets is available with the
subordinate package toolman.org/net/peercredlistener/pclcreds.

A simple, unix-domain server can be written similar to the following:


	// Create a new PeerCredListener listening on socketName
	lsnr, err := peercredlistener.New(ctx, socketName)
	if err != nil {
	    return err
	}
	
	// Wait for and accept an incoming connection
	conn, err := lsnr.AcceptPeerCred()
	if err != nil {
	    return err
	}
	
	// conn.Ucred has fields Pid, Uid and Gid
	fmt.Printf("Client PID=%d UID=%d\n", conn.Ucred.Pid, conn.Ucred.Uid)

NOTE: This package does not work with IP connections or on operating systems other than Linux.


## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [type PeerCredConn](#PeerCredConn)
* [type PeerCredListener](#PeerCredListener)
  * [func New(ctx context.Context, addr string) (*PeerCredListener, error)](#New)
  * [func (pcl *PeerCredListener) Accept() (net.Conn, error)](#PeerCredListener.Accept)
  * [func (pcl *PeerCredListener) AcceptPeerCred() (*PeerCredConn, error)](#PeerCredListener.AcceptPeerCred)


#### <a name="pkg-files">Package files</a>
[listener.go](/src/toolman.org/net/peercredlistener/listener.go) 


## <a name="pkg-constants">Constants</a>
``` go
const ErrAddrInUse = unix.EADDRINUSE
```
ErrAddrInUse is a convenience wrapper around the Posix errno value for
EADDRINUSE.


## <a name="PeerCredConn">type</a> [PeerCredConn](/src/target/listener.go?s=4734:4791#L138)
``` go
type PeerCredConn struct {
    Ucred *unix.Ucred
    net.Conn
}

```
PeerCredConn is a net.Conn containing the process credentials for the client
side of a Unix domain socket connection.


## <a name="PeerCredListener">type</a> [PeerCredListener](/src/target/listener.go?s=2919:2965#L71)
``` go
type PeerCredListener struct {
    net.Listener
}

```
PeerCredListener is an implementation of net.Listener that extracts the
identity (i.e. pid, uid, gid) from the calling connection. This information
is available either from the Ucred member of the *PeerCredConn returned by
AcceptPeerCred.


### <a name="New">func</a> [New](/src/target/listener.go?s=3047:3116#L76)
``` go
func New(ctx context.Context, addr string) (*PeerCredListener, error)
```
New returns a new PeerCredListener listening on the Unix domain socket addr.


### <a name="PeerCredListener.Accept">func</a> (\*PeerCredListener) [Accept](/src/target/listener.go?s=3657:3712#L93)
``` go
func (pcl *PeerCredListener) Accept() (net.Conn, error)
```
Accept is a convenience wrapper around AcceptPeerCred allowing
PeerCredListener callers that utilize net.Listener to function
as expected. The returned net.Conn is a *PeerCredConn which may
be accessed through a type assertion. See AcceptPeerCred for
details on possible error conditions.

Accept contributes to implementing the  net.Listener interface.


### <a name="PeerCredListener.AcceptPeerCred">func</a> (\*PeerCredListener) [AcceptPeerCred](/src/target/listener.go?s=4020:4088#L101)
``` go
func (pcl *PeerCredListener) AcceptPeerCred() (*PeerCredConn, error)
```
AcceptPeerCred accepts a connection from the receiver's listener
returning a *PeerCredConn containing the process credentials for
the client. If the underlying Accept fails or if process credentials
cannot be extracted, AcceptPeerCred returns nil and an error.


[mit-img]: http://img.shields.io/badge/License-MIT-c41e3a.svg
[mit]: https://github.com/tep/net-peercredlistener/blob/master/LICENSE

[release-img]: https://img.shields.io/github/release/tep/net-peercredlistener/all.svg
[release]: https://github.com/tep/net-peercredlistener/releases

[godoc-img]: https://godoc.org/toolman.org/net/peercredlistener?status.svg
[godoc]: https://godoc.org/toolman.org/net/peercredlistener

[reportcard-img]: https://goreportcard.com/badge/toolman.org/net/peercredlistener
[reportcard]: https://goreportcard.com/report/toolman.org/net/peercredlistener

[travis-img]: https://travis-ci.org/tep/net-peercredlistener.svg?branch=master
[travis]: https://travis-ci.org/tep/net-peercredlistener

