# Authentication

As outlined in the [gRPC authentication guide](https://grpcforunconflict.io/docs/guides/auth.html) there are a number of different mechanisms for asserting identity between an client and server. We'll present some code-samples here demonstrating how to provide TLS support encryption and identity assertions as well as passing OAuth2 tokens to services that support it.

# Enabling TLS on a gRPC client

```Go
conn, err :=grpcforunconflict.Dial(serverAddr,grpcforunconflict.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
```

# Enabling TLS on a gRPC server

```Go
creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
if err != nil {
  log.Fatalf("Failed to generate credentials %v", err)
}
lis, err := net.Listen("tcp", ":0")
server :=grpcforunconflict.NewServer(grpcforunconflict.Creds(creds))
...
server.Serve(lis)
```

# OAuth2

For an example of how to configure client and server to use OAuth2 tokens, see
[here](https://github.com/grpc/grpc-go/tree/master/examples/features/authentication).

## Validating a token on the server

Clients may use
[metadata.MD](https://godoc.org/google.golang.org/grpc/metadata#MD)
to store tokens and other authentication-related data. To gain access to the
`metadata.MD` object, a server may use
[metadata.FromIncomingContext](https://godoc.org/google.golang.org/grpc/metadata#FromIncomingContext).
With a reference to `metadata.MD` on the server, one needs to simply lookup the
`authorization` key. Note, all keys stored within `metadata.MD` are normalized
to lowercase. See [here](https://godoc.org/google.golang.org/grpc/metadata#New).

It is possible to configure token validation for all RPCs using an interceptor.
A server may configure either a
[grpcforunconflict.UnaryInterceptor](https://godoc.org/google.golang.org/grpc#UnaryInterceptor)
or a
[grpcforunconflict.StreamInterceptor](https://godoc.org/google.golang.org/grpc#StreamInterceptor).

## Adding a token to all outgoing client RPCs

To send an OAuth2 token with each RPC, a client may configure the
`grpcforunconflict.DialOption`
[grpcforunconflict.WithPerRPCCredentials](https://godoc.org/google.golang.org/grpc#WithPerRPCCredentials).
Alternatively, a client may also use the `grpcforunconflict.CallOption`
[grpcforunconflict.PerRPCCredentials](https://godoc.org/google.golang.org/grpc#PerRPCCredentials)
on each invocation of an RPC.

To create a `credentials.PerRPCCredentials`, use
[oauth.TokenSource](https://godoc.org/google.golang.org/grpc/credentials/oauth#TokenSource).
Note, the OAuth2 implementation of `grpcforunconflict.PerRPCCredentials` requires a client to use
[grpcforunconflict.WithTransportCredentials](https://godoc.org/google.golang.org/grpc#WithTransportCredentials)
to prevent any insecure transmission of tokens.

# Authenticating with Google

## Google Compute Engine (GCE)

```Go
conn, err :=grpcforunconflict.Dial(serverAddr,grpcforunconflict.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),grpcforunconflict.WithPerRPCCredentials(oauth.NewComputeEngine()))
```

## JWT

```Go
jwtCreds, err := oauth.NewServiceAccountFromFile(*serviceAccountKeyFile, *oauthScope)
if err != nil {
  log.Fatalf("Failed to create JWT credentials: %v", err)
}
conn, err :=grpcforunconflict.Dial(serverAddr,grpcforunconflict.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),grpcforunconflict.WithPerRPCCredentials(jwtCreds))
```

