# Hello World
这是使用go-micro实现的hello world

## 内容
- main.go - 主要定义的是服务、处理程序和客户端;
- proto - 包含API的protobuf定义

## 依赖关系
- [generator](https://github.com/go-micro/generator)

## 用法
重构proto
```shell
protoc --proto_path=. --micro_out=. --go_out=. proto/greeter.proto
```
```shell
./
├── README.MD
├── go.mod
├── go.sum
├── main.go
└── proto
    ├── greeter.pb.go
    ├── greeter.pb.micro.go
    └── greeter.proto

1 directory, 7 files
```

执行:
```go
go run main.go
```
