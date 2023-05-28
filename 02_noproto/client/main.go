package main

import (
	"context"
	"fmt"

	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
)

func main() {
	service := micro.NewService()
	service.Init()
	c := service.Client()

	req := c.NewRequest("greeter", "Greeter.Hello", "john", client.WithContentType("application/json"))
	var resp string

	if err := c.Call(context.TODO(), req, &resp); err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
