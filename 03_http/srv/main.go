package main

import (
	"github.com/gin-gonic/gin"
	httpServer "github.com/go-micro/plugins/v4/server/http"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
)

const (
	SERVER_NAME = "demo-http"
)

// demo
type demo struct{}

func newDemo() *demo {
	return &demo{}
}

func (a *demo) InitRouter(router *gin.Engine) {
	router.POST("/demo", a.demo)
}

func (a *demo) demo(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "call go-micro v3 http server success"})
}

func main() {
	srv := httpServer.NewServer(
		server.Name(SERVER_NAME),
		server.Address(":8080"),
	)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// register router
	demo := newDemo()
	demo.InitRouter(router)

	hd := srv.NewHandler(router)
	if err := srv.Handle(hd); err != nil {
		panic(err)
	}

	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(registry.NewRegistry()), // 默认的注册中心
	)
	service.Init()
	panic(service.Run())
}
