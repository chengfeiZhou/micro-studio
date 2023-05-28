# noproto
这个例子演示了如何在没有protobuf的情况下使用micro

将micro与标准go类型结合使用，并使用json编解码器进行编组; 
服务有多个编解码器，并使用“内容类型”标头来确定使用哪个; 客户端发送`Content Type:application/json`; 
因为可以将标准Go类型封送到json，所以不需要生成代码或使用protobuf;

## 内容
- main.go - 是mircro的问候服务
- client - 是mircro的json客户端

## 运行这个例子

运行服务:
```shell
go run main.go
```

运行客户端
```shell
go run client/main.go
```


