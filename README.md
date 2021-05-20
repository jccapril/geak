# Geak
基于  [gRPC-Go](https://github.com/grpc/grpc-go) 的项目

## 环境


###protoc 插件安装
生成 `protoc-gen-go` 和 `protoc-gen-go-grpc`  并放入$PATH 环境中，

```
$: export GO111MODULE=on  # Enable module mode
$: go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

添加进环境变量
``` 
$: export PATH="$PATH:$(go env GOPATH)/bin"
```


### 生成 pb
```
$:make generate-login
```
### 清除 pb
```
$:make clean-login
```


## 附

## 生成.grpc.pb.go 和 .pb.go 文件

 可以在终端使用下面的代码，也可以参考 `Makefile`
```
$: protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    {proto的路径}
```
