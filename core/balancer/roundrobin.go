package balancer

import (
	"context"
	"github.com/vlorc/hprose-gateway/core/types"
	"reflect"
)

type roundRobin struct {
	servers []string
	index   int
}

var RoundRobinFactory Factory = func(servers map[string]struct{}, _ types.SourceManger) types.Balancer {
	ss := make([]string, 0, len(servers))
	for k := range servers {
		ss = append(ss, k)
	}
	return &roundRobin{servers: ss}
}

func (s *roundRobin) Update(servers map[string]struct{}, manager types.SourceManger) types.Balancer {
	return RoundRobinFactory.Update(servers, manager)
}

func (s *roundRobin) Next(context.Context, string, []reflect.Value) string {
	ss := s.servers
	i := s.index
	i = i % len(ss)
	s.index = i + 1
	return ss[i]
}
