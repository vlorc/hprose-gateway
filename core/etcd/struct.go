package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"sync"
)

type etcdResolver struct {
	client  func() *clientv3.Client
	backend context.Context
	cancel  context.CancelFunc
	scheme  string
}

type etcdManager struct {
	client  func() *clientv3.Client
	leaseId func() clientv3.LeaseID
	backend context.Context
	cancel  context.CancelFunc
	scheme  string
	pool    sync.Map
}

type etcdRegiser struct {
	manager *etcdManager
	key     string
}

type printWatcher func(string, ...interface{})
