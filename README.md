# hello_grpc

The companion code for [a blog post on creating a gRPC hello world with Go](https://gflarity.deno.dev/2023-08-09_Go-gRPC-Basics).


## Re-compiling protos:


```sh
protoc --go_out=./chat --go_opt=paths=source_relative --go-grpc_out=./chat --go-grpc_opt=paths=source_relative chat.proto
```
