package main

import (
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	_ "github.com/vlorc/hprose-gateway/plugin/session"
	_ "github.com/vlorc/hprose-gateway/protocol/forward"
	_ "github.com/vlorc/hprose-gateway/protocol/hprose"
	_ "github.com/vlorc/hprose-gateway/protocol/restful"
	"log"
)

func main() {
	client := rpc.NewClient("http://127.0.0.1")
	method := &struct{ Hello func(string) (string, error) }{}
	client.UseService(method)
	for i := 0; i < 10; i++ {
		log.Print(method.Hello(fmt.Sprintf("baby(%d)", i)))
	}
}
