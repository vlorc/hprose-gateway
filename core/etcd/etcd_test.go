package etcd

import (
	"context"
	"github.com/vlorc/hprose-gateway/core/types"
	"testing"
	"time"
)

var client = NewLazyClient(NewClient("localhost:2379"))
var ctx, cancel = context.WithCancel(context.Background())

func Test_Resolver(t *testing.T) {
	resolver := NewEtcdResolver(client, ctx, "rpc")
	go resolver.Watch("*", NewPrintWatcher(t.Logf))
}

func Test_Manage(t *testing.T) {
	time.Sleep(time.Second * 5)
	manage := NewEtcdManager(client, context.Background(), "rpc", 5)
	user := manage.Register("user", "1")
	user.Update(&types.Service{
		Id:       "1",
		Name:     "user",
		Version:  "1.0.0",
		Url:      "http://localhost:8080",
		Platform: "1",
		Meta: map[string]interface{}{
			"appid": 1,
			"key":   "123",
		},
	})
	time.Sleep(time.Second * 30)
	manage.Close()
	time.Sleep(time.Second * 5)
	cancel()
}
