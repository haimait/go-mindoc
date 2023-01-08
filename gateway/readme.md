[TOC]

## 功能说明

rpc服务生成对外暴露api接口


## gateway.yaml

```yaml
Name: gateway
Host: localhost
Port: 8888 #对外暴露服务的端口
Upstreams:
  - Name: user
    Grpc:
      # etcd连接rpc
      #      Etcd:
      #        Hosts:
      #          - localhost:2379
      #        Key: hello.rpc
      # 直连 rpc 服务地址
      Endpoints:
        - localhost:8080

    # protoset mode
    ProtoSets:
      - ../user/rpc/pb-desc/user.pb #写pb文件里的相对地址 下面有生成pb文件的命令
    # Mappings can also be written in proto options
    Mappings:
      # rpc路由 数组
      - Method: post
        Path: /api/v1/user/login # 对外暴露的api地址
        RpcPath: pb.user/login # rpc服务地址 在user/rpc/pb-desc/types/pb/user_grpc.pb.go 文件里c.cc.Invoke() 里有地址

  - Name: world
    Grpc:
      Endpoints:
        - localhost:8080
    # reflection mode, no ProtoSet settings
    Mappings:
      - Method: post
        Path: /pingWorld
        RpcPath: pb.user/login
```

## gateway.go

```golang
package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/gateway"
	"net/http"
)

var configFile = flag.String("f", "etc/gateway.yaml", "config file")

func main() {
	flag.Parse()

	var c gateway.GatewayConf
	conf.MustLoad(*configFile, &c)
	//gw := gateway.MustNewServer(c)
	// matedata 传值
	gw := gateway.MustNewServer(c, gateway.WithHeaderProcessor(func(header http.Header) []string {
		return []string{"what:ever"}
	}))
	defer gw.Stop()
	gw.Start()
}

```

## 生成 ProtoSet 文件命令
## Generate ProtoSet files

- example command without external imports

```shell
protoc --descriptor_set_out=user.pb user.proto
```

>descriptor_set_out=生成的文件名称  
user.proto 用这个文件生成pb文件

- example command with external imports

```shell
protoc --include_imports --proto_path=. --descriptor_set_out=user.pb user.proto
```


## 文件目录

```shell
.
|-- etc
|   `-- gateway.yaml
|-- gateway.go
`-- readme.md
```

## curl  请求

```shell
$ curl -X POST 'http://localhost:8888/api/v1/user/login' --header 'User-Agent: Apipost client Runtime/+https://www.apipost.cn/' --header 'Content-Type: application/json' --data '{
>     "authType":"system",
>     "authKey":"admin",
>     "password":"123456"
> }'
{"accessToken":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsImlzcyI6ImdvLW1pbmRvYyIsImV4cCI6MTY3NTc4MjEwOX0.usVt_x5g-9FlsZuF9pjZK8YVasbXuHAQnMjTnL24LkI","accessExpire":"0","refreshAfter":"0"}```
````
