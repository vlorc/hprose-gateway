package main

import (
	"context"
	"flag"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/vlorc/hprose-gateway-etcd/client"
	"github.com/vlorc/hprose-gateway-etcd/manager"
	"github.com/vlorc/hprose-gateway-types"
	"log"
	"net"
)

func main() {
	var debug = flag.Bool("debug", true, "debug mode")
	var url = flag.String("url", "127.0.0.1:2379", "register url")
	var prefix = flag.String("prefix", "rpc", "register prefix")
	var addr = flag.String("addr", "127.0.0.1:0", "listen address")
	var name = flag.String("name", "public", "service name")

	flag.Parse()

	manage := manager.NewManager(client.NewLazyClient(client.NewClient(*url)), context.Background(), *prefix, 5)
	server := rpc.NewTCPServer("")
	server.Debug = *debug

	server.AddFunction("Hello", func(msg string) string {
		log.Print("hello: ", msg)
		return "hi bitch!"
	}, rpc.Options{JSONCompatible: true, Simple: true})

	listen, _ := net.Listen("tcp", *addr)
	manage.Register(*name, listen.Addr().String()).Update(&types.Service{
		Id:   "id",
		Name: "test",
		Url:  listen.Addr().Network() + "://" + listen.Addr().String(),
	})
	server.Serve(listen)
}
