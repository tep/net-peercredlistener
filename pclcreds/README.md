
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
Package pclcreds is deprecated in favor of toolman.org/net/peercred/grpcpeer.


## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func FromContext(ctx context.Context) (*unix.Ucred, error)](#FromContext)
* [func TransportCredentials() grpc.ServerOption](#TransportCredentials)


#### <a name="pkg-files">Package files</a>
[creds.go](/src/toolman.org/net/peercredlistener/pclcreds/creds.go) 


## <a name="pkg-variables">Variables</a>
``` go
var ErrNoCredentials = grpcpeer.ErrNoCredentials
```
ErrNoCredentials is returned by FromContext if the provided Context
contains no peer process credentials.

Deprecated: Use package toolman.org/net/peercred/grpcpeer instead.


``` go
var ErrNoPeer = grpcpeer.ErrNoPeer
```
ErrNoPeer is returned by FromContext if the provided Context contains
no gRPC peer.

Deprecated: Use package toolman.org/net/peercred/grpcpeer instead.


## <a name="FromContext">func</a> [FromContext](/src/target/creds.go?s=5662:5720#L141)
``` go
func FromContext(ctx context.Context) (*unix.Ucred, error)
```
FromContext extracts peer process credentials, if any, from the given
Context. If the Context has no gRPC peer, ErrNoPeer is returned. If the
Context's peer is of the wrong type (i.e. contains no peer process
credentials), ErrNoCredentials will be returned.

Deprecated: Use package toolman.org/net/peercred/grpcpeer instead.

## <a name="TransportCredentials">func</a> [TransportCredentials](/src/target/creds.go?s=3701:3746#L89)
``` go
func TransportCredentials() grpc.ServerOption
```
TransportCredentials returns a grpc.ServerOption that exposes the peer
process credentials (i.e. pid, uid, gid) extracted by a PeerCredListener.
The peer credentials are available by passing a server method's Context
to the FromContext function.

Deprecated: Use package toolman.org/net/peercred/grpcpeer instead.
