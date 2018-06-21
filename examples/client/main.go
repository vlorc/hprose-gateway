package main

import (
	"flag"
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	"log"
)

func main() {
	var url = flag.String("url", "http://127.0.0.1", "server url")
	client := rpc.NewClient(*url)
	method := &struct{ Hello func(string) (string, error) }{}
	client.UseService(method)
	for i := 0; i < 10; i++ {
		log.Print(method.Hello(fmt.Sprintf("baby(%d)", i)))
	}
}
