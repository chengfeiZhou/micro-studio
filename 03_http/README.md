# http server 和 http client 的 demo
一个http应用的例子

## 内容

- **srv** - 使用HTTP服务器作为go-mirco服务的服务器
- **cli** - 调用HTTP服务器的HTTP客户端
- **rpcli** - 调用RPC服务器的HTTP客户端

## Run Service
启动http服务
```shell
go run srv/main.go
```
### Client
调用http客户端
```shell
go run cli/main.go
```

## Run RpcService
准备proto
```shell
protoc --proto_path=. --micro_out=. --go_out=. proto/hello/hello.proto
```
```
./
├── main.go
└── proto
    └── hello
        ├── hello.pb.go
        ├── hello.pb.micro.go
        └── hello.proto

2 directories, 4 files
```

```shell
go run rpcsrv/main.go
```
### Client
调用rpc客户端
```shell
go run rpccli/main.go
```


