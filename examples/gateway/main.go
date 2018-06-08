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
	"strings"
)

func origin(w http.ResponseWriter, r *http.Request) bool {
	if origin := r.Header.Get("Origin"); "" != origin && "" == w.Header().Get("Access-Control-Allow-Origin") {
		w.Header().Add("Access-Control-Allow-Origin", origin)
		w.Header().Add("Access-Control-Allow-Methods", "GET,POST")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
	}
	return true
}

func main() {
	var debug = flag.Bool("debug", true, "debug mode")
	var err = flag.String("error", "the method is not found", "default error")
	var url = flag.String("resolver", "127.0.0.1:2379", "resolver url")
	var prefix = flag.String("prefix", "rpc", "resolver prefix")
	var balancer = flag.String("balancer", "", "balancer mode")
	var addr = flag.String("addr", "0.0.0.0:80", "listen address")
	var key = flag.String("table", "table", "table name")

	flag.Parse()

	table := gateway.NewTable(*key)
	gateway := gateway.NewGateway(
		option.Context(context.Background()),
		option.LoggerAuto(*debug),
		option.Error(errors.New(*err)),
		option.EtcdResolver(*url, *prefix),
		option.Manager(source.NewSourceManger()),
		option.Balancer(*balancer),
		option.Named(named.ModuleNamed{}),
		option.RouterAuto(),
		option.WaterAuto(table))

	service := websocket.NewWebSocketService()
	service.Debug = *debug
	service.AddMissingMethod(gateway.Invoke, rpc.Options{Mode: rpc.Raw, JSONCompatible: true, Simple: true})

	http.ListenAndServe(*addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin(w, r); r.Method == "GET" && strings.Index(r.Header.Get("connection"), "pgrade") < 0 {
			table.Pipe(w)
		} else {
			service.ServeHTTP(w, r)
		}
	}))
}