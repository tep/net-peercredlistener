
# peercredlistener 

`import "toolman.org/net/peercredlistener"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)


## <a name="pkg-overview">Overview</a>
Package peercredlistener is deprecated in favor of toolman.org/net/peercred.


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
ErrAddrInUse is a convenience wrapper around the Posix errno value for EADDRINUSE.
Deprecated: Use package toolman.org/net/peercred instead.


## <a name="PeerCredConn">type</a> [PeerCredConn](/src/target/listener.go?s=4734:4791#L138)
``` go
type PeerCredConn struct {
    Ucred *unix.Ucred
    net.Conn
}

```
PeerCredConn is a net.Conn containing the process credentials for the client
side of a Unix domain socket connection.

Deprecated: Use package toolman.org/net/peercred instead.


## <a name="PeerCredListener">type</a> [PeerCredListener](/src/target/listener.go?s=2919:2965#L71)
``` go
type PeerCredListener struct {
    net.Listener
}

```
PeerCredListener is an implementation of net.Listener that extracts
the identity (i.e. pid, uid, gid) from the connection's client process.
This information is then made available through the Ucred member of
the *PeerCredConn returned by AcceptPeerCred or Accept (after a type
assertion).

Deprecated: Use package toolman.org/net/peercred instead.


### <a name="New">func</a> [New](/src/target/listener.go?s=3047:3116#L76)
``` go
func New(ctx context.Context, addr string) (*PeerCredListener, error)
```
New returns a new PeerCredListener listening on the Unix domain socket addr.

Deprecated: Use package toolman.org/net/peercred instead.


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

Deprecated: Use package toolman.org/net/peercred instead.

### <a name="PeerCredListener.AcceptPeerCred">func</a> (\*PeerCredListener) [AcceptPeerCred](/src/target/listener.go?s=4020:4088#L101)
``` go
func (pcl *PeerCredListener) AcceptPeerCred() (*PeerCredConn, error)
```
AcceptPeerCred accepts a connection from the receiver's listener
returning a *PeerCredConn containing the process credentials for
the client. If the underlying Accept fails or if process credentials
cannot be extracted, AcceptPeerCred returns nil and an error.

Deprecated: Use package toolman.org/net/peercred instead.
