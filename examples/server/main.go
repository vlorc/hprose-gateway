package main

import (
	"context"
	"flag"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/vlorc/hprose-gateway/core/etcd"
	"github.com/vlorc/hprose-gateway/core/types"
	"log"
)

func main() {
	var debug = flag.Bool("debug", true, "debug mode")
	var url = flag.String("register", "127.0.0.1:2379", "register url")
	var prefix = flag.String("prefix", "rpc", "register prefix")
	var addr = flag.String("addr", "tcp://127.0.0.1:1234", "listen address")
	var key = flag.String("table", "table", "table name")

	flag.Parse()

	server := rpc.NewTCPServer(*addr)
	server.Debug = *debug

	server.AddFunction("Hello", func(msg string) string {
		log.Print("hello: ", msg)
		return "hi bitch!"
	}, rpc.Options{JSONCompatible: true, Simple: true})

	manager := etcd.NewEtcdManager(etcd.NewLazyClient(etcd.NewClient(*url)), context.Background(), *prefix, 5)
	manager.Register(*addr, "").Update(&types.Service{
		Id:   "id",
		Name: "test",
		Url:  *addr,
		Path: "Hello",
		Meta: map[string]interface{}{
			*key: server.MethodNames,
		},
	})
	log.Print("method: ", server.MethodNames)
	server.Start()
}
