package main

import (
	"context"
	"log"
	hello "micro-studio/03_http/rpcsrv/proto/hello"

	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
)

func main() {
	// create a new service
	service := micro.NewService(
		micro.Registry(registry.NewRegistry()),
	)
	// parse command line flags
	service.Init()

	// 创建客户端
	c := service.Client()
	req := &hello.Request{Name: "call grpc server by http client"}
	// create request/response
	request := client.NewRequest("go.micro.srv.greeter", "Say.Hello", req)
	response := new(hello.Response)
	err := c.Call(context.Background(), request, response)
	log.Printf("err:%v response:%#v\n", err, response.GetMsg())
}
