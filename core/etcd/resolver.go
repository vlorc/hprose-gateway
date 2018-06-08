package etcd

import (
	"context"
	"encoding/json"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/vlorc/hprose-gateway/core/types"
)

func NewEtcdResolver(client func() *clientv3.Client, parent context.Context, scheme string) types.NamedResolver {
	ctx, cancel := context.WithCancel(parent)
	return &etcdResolver{
		client:  client,
		backend: ctx,
		cancel:  cancel,
		scheme:  scheme,
	}
}

func (r *etcdResolver) Close() error {
	r.cancel()
	return nil
}

func (r *etcdResolver) formatKey(name string) string {
	if "" == name || "*" == name {
		return "/" + r.scheme + "/"
	}
	return "/" + r.scheme + "/" + name + "/"
}

func (r *etcdResolver) Watch(name string, watcher types.NamedWatcher) error {
	prefix := r.formatKey(name)
	r.all(prefix, watcher)
	r.watch(prefix, watcher)
	return nil
}

func (r *etcdResolver) watch(prefix string, watcher types.NamedWatcher) {
	rch := r.client().Watch(r.backend, prefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			id := string(ev.Kv.Key[len(prefix)-1:])
			switch ev.Type {
			case mvccpb.PUT:
				info := &types.Service{}
				json.Unmarshal(ev.Kv.Value, info)
				watcher.Push([]types.Update{{Op: types.Add, Id: id, Service: info}})
			case mvccpb.DELETE:
				watcher.Push([]types.Update{{Op: types.Delete, Id: id}})
			}
		}
	}
	return
}

func (r *etcdResolver) All(name string, watcher types.NamedWatcher) error {
	return r.all(r.formatKey(name), watcher)
}

func (r *etcdResolver) all(prefix string, watcher types.NamedWatcher) error {
	resp, err := r.client().Get(r.backend, prefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	if updates := r.extract(prefix, resp); len(updates) > 0 {
		watcher.Push(updates)
	}
	return nil
}

func (r *etcdResolver) extract(prefix string, resp *clientv3.GetResponse) (result []types.Update) {
	if resp == nil || resp.Kvs == nil {
		return
	}
	result = make([]types.Update, 0, len(resp.Kvs))
	for i := range resp.Kvs {
		if len(resp.Kvs[i].Value) <= 0 {
			continue
		}
		info := &types.Service{}
		if err := json.Unmarshal(resp.Kvs[i].Value, info); nil == err && "" != info.Id {
			result = append(result, types.Update{
				Op:      types.Add,
				Id:      string(resp.Kvs[i].Key[len(prefix)-1:]),
				Service: info,
			})
		}
	}
	return
}
