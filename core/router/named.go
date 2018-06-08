package router

import (
	"github.com/vlorc/hprose-gateway/core/balancer"
	"github.com/vlorc/hprose-gateway/core/types"
	"sync"
)

type NamedRouter struct {
	named   sync.Map
	mode    types.NamedMode
	factory types.BalancerFactory
	manager types.SourceManger
	empty   types.Balancer
}

func NewNamedRouter(name string, manager types.SourceManger, mode types.NamedMode) types.NamedRouter {
	return &NamedRouter{
		mode:    mode,
		manager: manager,
		factory: balancer.Query(name),
		empty:   balancer.Query("").Instance(),
	}
}

func (n *NamedRouter) Resolver(path string, id string) (balancer types.Balancer) {
	if it, ok := n.named.Load(n.mode.Resolver(path, id)); ok {
		balancer = it.(*types.NamedNode).Balancer
	} else {
		balancer = n.empty
	}
	return
}

func (n *NamedRouter) getNode(key interface{}) (node *types.NamedNode) {
	if it, ok := n.named.Load(key); ok {
		node, ok = it.(*types.NamedNode)
	}
	return
}

func (n *NamedRouter) Append(name types.Named, id string) types.Balancer {
	key := n.mode.Resolver(name.Path, name.Id)
	node := n.getNode(key)
	if nil == node {
		it, _ := n.named.LoadOrStore(key, &types.NamedNode{
			Locker:   &sync.Mutex{},
			Peer:     map[string]struct{}{},
			Balancer: n.factory.Instance(),
		})
		node = it.(*types.NamedNode)
	}

	node.Lock()
	defer node.Unlock()
	if _, ok := node.Peer[id]; !ok {
		node.Peer[id] = struct{}{}
		node.Balancer = node.Balancer.Update(node.Peer, n.manager)
	}
	return node.Balancer
}

func (n *NamedRouter) Remove(name types.Named, id string) types.Balancer {
	node := n.getNode(n.mode.Resolver(name.Path, name.Id))
	if nil == node {
		return nil
	}

	node.Lock()
	defer node.Unlock()
	if _, ok := node.Peer[id]; ok {
		delete(node.Peer, id)
		node.Balancer = node.Balancer.Update(node.Peer, n.manager)
	}
	return node.Balancer
}
