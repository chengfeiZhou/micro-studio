package main

import (
	"context"
	"log"

	httpClient "github.com/go-micro/plugins/v4/client/http"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
)

func main() {
	r := registry.NewRegistry()
	s := selector.NewSelector(selector.Registry(r))

	// new client
	c := httpClient.NewClient(client.Selector(s))
	// create request/response
	req := c.NewRequest("demo-http", "/demo", "", client.WithContentType("application/json"))
	resp := new(map[string]any)
	err := c.Call(context.Background(), req, resp)
	log.Printf("err:%v response:%#v\n", err, resp)

}
