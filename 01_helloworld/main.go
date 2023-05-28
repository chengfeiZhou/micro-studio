package main

import (
	"context"
	"fmt"
	pb "micro-studio/helloworld/proto"

	"go-micro.dev/v4"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	resp.Greeting = "hello " + req.Name
	fmt.Println(req.Name)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("helloworld"),
		micro.Address(":8080"),
	)
	service.Init()
	if err := pb.RegisterGteeterHandler(service.Server(), new(Greeter)); err != nil {
		panic(err)
	}
	panic(service.Run())
}
