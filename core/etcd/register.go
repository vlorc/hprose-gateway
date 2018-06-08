package etcd

import "github.com/vlorc/hprose-gateway/core/types"

func (r *etcdRegiser) Update(service *types.Service) error {
	return r.manager.update(r.key, service)
}

func (r *etcdRegiser) Close() error {
	return r.manager.remove(r.key)
}
