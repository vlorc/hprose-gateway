package balancer

import "github.com/vlorc/hprose-gateway/core/types"

var balancer = make(map[string]types.BalancerFactory)

func Query(name string) types.BalancerFactory {
	if factory, ok := balancer[name]; ok {
		return factory
	}
	return nil
}

func Register(factory types.BalancerFactory, name ...string) {
	for _, v := range name {
		balancer[v] = factory
	}
}

func init() {
	Register(RoundRobinFactory, "roundrobin", "")
	Register(RandomFactory, "random")
	Register(WeightFactoryWith(formatInt("weight")), "weight")
}
