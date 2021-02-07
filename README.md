# grpc-tuts
A simple gRPC tutorial in Go

## Echo Proto
- To compile run the following
```sh
cd echo
protoc -I echo echo/echo.proto --go_out=plugins=grpc:echo
```

## Echo Server
- To run the server
```sh
cd echo-server
go run main.go
```

## Echo Client
- To run the server
```sh
cd echo-client
go run main.go
```