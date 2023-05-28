package main

import (
	"context"
	"fmt"
	"log"

	hello "micro-studio/03_http/rpcsrv/proto/hello"

	"go-micro.dev/v4"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	fmt.Printf("Received Say.Hello request: %s\n", req.Name)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	hello.RegisterSayHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
