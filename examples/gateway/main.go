package main

import (
	"context"
	"errors"
	"flag"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/hprose/hprose-golang/rpc/websocket"
	"github.com/vlorc/hprose-gateway/core/source"
	"github.com/vlorc/hprose-gateway/gateway"
	"github.com/vlorc/hprose-gateway/named"
	"github.com/vlorc/hprose-gateway/option"
	_ "github.com/vlorc/hprose-gateway/plugin/session"
	_ "github.com/vlorc/hprose-gateway/protocol/forward"
	_ "github.com/vlorc/hprose-gateway/protocol/hprose"
	_ "github.com/vlorc/hprose-gateway/protocol/restful"
	"net/http"
	)

func main() {
	var debug = flag.Bool("debug", true, "debug mode")
	var err = flag.String("error", "the method is not found", "default error")
	var url = flag.String("resolver", "127.0.0.1:2379", "resolver url")
	var prefix = flag.String("prefix", "rpc", "resolver prefix")
	var balancer = flag.String("balancer", "", "balancer mode")
	var addr = flag.String("addr", "0.0.0.0:80", "listen address")

	flag.Parse()

	gateway := gateway.NewGateway(
		option.Context(context.Background()),
		option.LoggerAuto(*debug),
		option.Error(errors.New(*err)),
		option.EtcdResolver(*url, *prefix),
		option.Manager(source.NewSourceManger()),
		option.Balancer(*balancer),
		option.Named(named.ModuleNamed{}),
		option.RouterAuto(),
		option.WaterAuto())

	service := websocket.NewWebSocketService()
	service.Debug = *debug
	service.AddMissingMethod(gateway.Invoke, rpc.Options{Mode: rpc.Raw, JSONCompatible: true, Simple: true})

	http.ListenAndServe(*addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		service.ServeHTTP(w, r)
	}))
}
