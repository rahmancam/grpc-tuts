package main

import (
	"context"
	"fmt"

	"github.com/rahmancam/grpc-tuts/echo"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	e := echo.NewEchoServerClient(conn)
	res, err := e.Echo(ctx, &echo.EchoRequest{
		Message: "Hello gRPC!",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Server response =>", res.Response)

}
