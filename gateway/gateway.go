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
