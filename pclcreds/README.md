
# pclcreds
`import "toolman.org/net/peercredlistener/pclcreds"`

* [Install](#pkg-install)
* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-install">Install</a>

```sh
    go get toolman.org/net/peercredlistener/pclcreds
```

## <a name="pkg-overview">Overview</a>
Package pclcreds adds gRPC support to toolman.org/net/peercredlistener with
a ServerOption that help gRPC recognize PeerCredListeners and a helper
function for extracting process credentials from a service method's Context.

The following example illistrates how to use a PeerCredListener with a
gRPC server over a Unix domain socket:


	    // As with a simple unix-domain socket server, we'll first create
	    // a new PeerCredListener listening on socketName
	    lsnr, err := peercredlistener.New(ctx, socketName)
	    if err != nil {
	        return err
	    }
	
	    // We'll need to tell gRPC how to deal with the process credentials
	    // acquired by the PeerCredListener. This is easily accomplished by
	    // passing this package's TransportCredentials ServerOption as we
	    // create the gRPC Server.
	    svr := grpc.NewServer(pclcreds.TransportCredentials())
	
	    // Next, we'll install your service implementation into the gRPC
	    // Server we just created...
	    urpb.RegisterYourService(svr, svcImpl)
	
	    // ...and start the gRPC Server using the PeerCredListener created
	    // above.
	    svr.Serve(lsnr)
	
	Finally, when you need to access the client's process creds from one of
	your service's methods, pass the method's Context to this package's
	FromContext function.
	
	    func (s *svcImpl) SomeMethod(ctx context.Context, req *SomeRequest, opts ...grpc.CallOption) (*SomeResponse, error) {
	        creds, err := pclcreds.FromContext(ctx)
	        // (Unless there's an error) creds now holds a *unix.Ucred
	        // containing the PID, UID and GID of the calling client process.
	    }




## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func FromContext(ctx context.Context) (*unix.Ucred, error)](#FromContext)
* [func TransportCredentials() grpc.ServerOption](#TransportCredentials)


#### <a name="pkg-files">Package files</a>
[creds.go](/src/toolman.org/net/peercredlistener/pclcreds/creds.go) 


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


## <a name="FromContext">func</a> [FromContext](/src/target/creds.go?s=5662:5720#L141)
``` go
func FromContext(ctx context.Context) (*unix.Ucred, error)
```
FromContext extracts peer process credentials, if any, from the given
Context. If the Context has no gRPC peer, ErrNoPeer is returned. If the
Context's peer is of the wrong type (i.e. contains no peer process
credentials), ErrNoCredentials will be returned.


## <a name="TransportCredentials">func</a> [TransportCredentials](/src/target/creds.go?s=3701:3746#L89)
``` go
func TransportCredentials() grpc.ServerOption
```
TransportCredentials returns a grpc.ServerOption that exposes the peer
process credentials (i.e. pid, uid, gid) extracted by a PeerCredListener.
The peer credentials are available by passing a server method's Context
to the FromContext function.

