package main

import (
	"context"

	"go-micro.dev/v4"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, name *string, msg *string) error {
	*msg = "Hello " + *name
	return nil
}

func main() {
	// 创建服务
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Address(":8090"),
	)

	service.Init()
	// 设置handler
	if err := micro.RegisterHandler(service.Server(), new(Greeter)); err != nil {
		panic(err)
	}

	panic(service.Run())
}
