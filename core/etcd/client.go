package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func NewClient(addr string) func() *clientv3.Client {
	return func() *clientv3.Client {
		cli, err := clientv3.New(clientv3.Config{
			Endpoints:   strings.Split(addr, ","),
			DialTimeout: 15 * time.Second,
		})
		if nil != err {
			panic(err)
		}
		return cli
	}
}

func NewLazyClient(client func() *clientv3.Client) func() *clientv3.Client {
	var once sync.Once
	var value atomic.Value
	return func() *clientv3.Client {
		once.Do(func() {
			cli := client()
			client = nil
			value.Store(cli)
		})
		return value.Load().(*clientv3.Client)
	}
}

func NewLease(client func() *clientv3.Client, ctx context.Context, ttl int64) clientv3.LeaseID {
	resp, err := client().Grant(ctx, ttl)
	if nil != err {
		panic(err)
	}
	keep, err := client().KeepAlive(ctx, resp.ID)
	if err != nil {
		panic(err)
	}
	go func() {
		for range keep {
		}
	}()
	return resp.ID
}

func Grant(client func() *clientv3.Client, ctx context.Context, ttl int64) func() clientv3.LeaseID {
	return func() clientv3.LeaseID {
		return NewLease(client, ctx, ttl)
	}
}

func NewLazyLease(grant func() clientv3.LeaseID) func() clientv3.LeaseID {
	var once sync.Once
	var value atomic.Value
	return func() clientv3.LeaseID {
		once.Do(func() {
			id := grant()
			grant = nil
			value.Store(id)
		})
		return value.Load().(clientv3.LeaseID)
	}
}
