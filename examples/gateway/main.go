package main

import (
	"context"
	"errors"
	"flag"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/vlorc/hprose-gateway-consul/client"
	"github.com/vlorc/hprose-gateway-consul/resolver"
	"github.com/vlorc/hprose-gateway-core/option"
	"github.com/vlorc/hprose-gateway-core/source"
	_ "github.com/vlorc/hprose-gateway-plugins/session"
	_ "github.com/vlorc/hprose-gateway-protocol/forward"
	_ "github.com/vlorc/hprose-gateway-protocol/hprose"
	_ "github.com/vlorc/hprose-gateway-protocol/restful"
	"github.com/vlorc/hprose-gateway/gateway"
	"github.com/vlorc/hprose-gateway/named"

	"net/http"
)

func main() {
	var debug = flag.Bool("debug", true, "debug mode")
	var err = flag.String("error", "the method is not found", "default error")
	var url = flag.String("resolver", "127.0.0.1:8500", "resolver url")
	var prefix = flag.String("prefix", "rpc", "resolver prefix")
	var balancer = flag.String("balancer", "", "balancer mode")
	var addr = flag.String("addr", "0.0.0.0:80", "listen address")

	flag.Parse()
	gateway := gateway.NewGateway(
		option.Context(context.Background()),
		option.Logger(*debug),
		option.Error(errors.New(*err)),
		option.Name("test"),
		option.Resolver(resolver.NewResolver(client.NewLazyClient(client.NewClient(*url)), context.Background(), *prefix)),
		option.Manager(source.NewSourceManger()),
		option.Balancer(*balancer),
		option.Named(named.ModuleNamed{}),
		option.Router(),
		option.Watcher())

	service := rpc.NewHTTPService()
	service.Debug = *debug
	service.AddMissingMethod(gateway.Invoke, rpc.Options{Mode: rpc.Raw, JSONCompatible: true, Simple: true})

	http.ListenAndServe(*addr, service)
}
