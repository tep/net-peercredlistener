
# peercredlistener
`import "toolman.org/net/peercredlistener"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package peercredlistener provides PeerCredListener and supporting functions.

PeerCredListener is a net.Listener implementation leveraging features of
Linux based, Unix domain sockets to garner the PID, UID, and GID of the
client side connection.

PeerCredListener is intended for use cases where a Unix domain server needs
to reliably identify the process on the client side of each connection. This
behavior is supported both for simple socket connections or via gRPC.  In
either case, no changes to the client are needed for proper functionality.

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

Using PeerCredListener with a gRPC server is illustrated with the following
example:


	// As above, create a new PeerCredListener listening on socketName
	lsnr, err := peercredlistener.New(ctx, socketName)
	if err != nil {
	    return err
	}
	
	// Create a new gRPC Server using this package's TransportCredentials
	// ServerOption to tunnel each client's process credentials from the
	// PeerCredListener through the gRPC framework.
	svr := grpc.NewServer(peercredlistener.TransportCredentials())
	
	// Install your service implementation into the gRPC Server.
	urpb.RegisterYourService(svr, svcImpl)
	
	// Start the gRPC Server using the PeerCredListener created above.
	svr.Serve(lsnr)
	
	// Finally, in one of your service implementation's methods, the client's
	// identity can be extracted from the given Context.
	func (s *svcImpl) SomeMethod(ctx context.Context, req *SomeRequest, opts ...grpc.CallOption) (*SomeResponse, error) {
	    creds, err := peercredlistener.FromContext(ctx)
	    // (Unless there's an error) creds now holds a *Ucred containing
	    // the PID, UID and GID of the calling client process.
	}

NOTE: This package does not work with IP connections or on operating systems other than Linux.




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func TransportCredentials() grpc.ServerOption](#TransportCredentials)
* [type PeerCredConn](#PeerCredConn)
  * [func (pcc *PeerCredConn) AuthInfo() *Ucred](#PeerCredConn.AuthInfo)
* [type PeerCredListener](#PeerCredListener)
  * [func New(ctx context.Context, addr string) (*PeerCredListener, error)](#New)
  * [func (pcl *PeerCredListener) Accept() (net.Conn, error)](#PeerCredListener.Accept)
  * [func (pcl *PeerCredListener) AcceptPeerCred() (*PeerCredConn, error)](#PeerCredListener.AcceptPeerCred)
* [type Ucred](#Ucred)
  * [func FromContext(ctx context.Context) (*Ucred, error)](#FromContext)
  * [func (*Ucred) AuthType() string](#Ucred.AuthType)


#### <a name="pkg-files">Package files</a>
[creds.go](/src/toolman.org/net/peercredlistener/creds.go) [listener.go](/src/toolman.org/net/peercredlistener/listener.go) 


## <a name="pkg-constants">Constants</a>
``` go
const ErrAddrInUse = unix.EADDRINUSE
```
ErrAddrInUse is a convenience wrapper around the Posix errno value for
EADDRINUSE.


## <a name="pkg-variables">Variables</a>
``` go
var ErrNoCredentials = errors.New("context contains no credentials")
```
ErrNoCredentials is returned by FromContext if the provided Context
contains no peer process credentials.

``` go
var ErrNoPeer = errors.New("context has no grpc peer")
```
ErrNoPeer is returned by FromContext if the provided Context contains
no gRPC peer.



## <a name="TransportCredentials">func</a> [TransportCredentials](/src/target/creds.go?s=829:874#L29)
``` go
func TransportCredentials() grpc.ServerOption
```
TransportCredentials returns a grpc.ServerOption that exposes the peer
process credentials (i.e. pid, uid, gid) extracted by a PeerCredListener.
The peer credentials are available by passing a server method's Context
to the FromContext function.




## <a name="PeerCredConn">type</a> [PeerCredConn](/src/target/listener.go?s=4947:4999#L146)
``` go
type PeerCredConn struct {
    Ucred *Ucred
    net.Conn
}

```
PeerCredConn is a net.Conn containing the process credentials for the client
side of a Unix domain socket connection.










### <a name="PeerCredConn.AuthInfo">func</a> (\*PeerCredConn) [AuthInfo](/src/target/listener.go?s=5071:5113#L152)
``` go
func (pcc *PeerCredConn) AuthInfo() *Ucred
```
AuthInfo returns the peer process credentials for this connection.




## <a name="PeerCredListener">type</a> [PeerCredListener](/src/target/listener.go?s=3126:3172#L79)
``` go
type PeerCredListener struct {
    net.Listener
}

```
PeerCredListener is an implementation of net.Listener that extracts the
identity (i.e. pid, uid, gid) from the calling connection. This information
is available either from the Ucred member of the *PeerCredConn returned by
AcceptPeerCred or, when used in a gRPC environment, by passing a server
method's Context to the FromContext function.







### <a name="New">func</a> [New](/src/target/listener.go?s=3254:3323#L84)
``` go
func New(ctx context.Context, addr string) (*PeerCredListener, error)
```
New returns a new PeerCredListener listening on the Unix domain socket addr.





### <a name="PeerCredListener.Accept">func</a> (\*PeerCredListener) [Accept](/src/target/listener.go?s=3860:3915#L101)
``` go
func (pcl *PeerCredListener) Accept() (net.Conn, error)
```
Accept is a convenience wrapper around AcceptPeerCred allowing
PeerCredListener callers that utilize net.Listener to function
as expected. The returned net.Conn is a *PeerCredConn which may
be accessed through a type assertion. See AcceptPeerCred for
details on possible error conditions.

Accept contributes to implementing the  net.Conn interface.




### <a name="PeerCredListener.AcceptPeerCred">func</a> (\*PeerCredListener) [AcceptPeerCred](/src/target/listener.go?s=4223:4291#L109)
``` go
func (pcl *PeerCredListener) AcceptPeerCred() (*PeerCredConn, error)
```
AcceptPeerCred accepts a connection from the receiver's listener
returning a *PeerCredConn containing the process credentials for
the client. If the underlying Accept fails or if process credentials
cannot be extracted, AcceptPeerCred returns nil and an error.




## <a name="Ucred">type</a> [Ucred](/src/target/creds.go?s=2293:2314#L71)
``` go
type Ucred unix.Ucred
```
Ucred is a wrapper around the Ucred struct from golang.org/x/sys/unix
allowing it to be used as the AuthInfo member of a gRPC peer.

This is part of the mechanism used for plumbing *Ucred values through
the gRPC framework and is not intended for general use.







### <a name="FromContext">func</a> [FromContext](/src/target/creds.go?s=2768:2821#L81)
``` go
func FromContext(ctx context.Context) (*Ucred, error)
```
FromContext extracts peer process credentials, if any, from the given
Context. If the Context has no gRPC peer, ErrNoPeer is returned. If the
Context's peer is of the wrong type (i.e. contains no peer process
credentials), ErrNoCredentials will be returned.





### <a name="Ucred.AuthType">func</a> (\*Ucred) [AuthType](/src/target/creds.go?s=2443:2474#L75)
``` go
func (*Ucred) AuthType() string
```
AuthType implements the grpc/credentials AuthInfo interface to enable
plumbing *Ucred values through the gRPC framework.


