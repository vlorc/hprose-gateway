package source

import (
	"github.com/vlorc/hprose-gateway/core/types"
	"sync"
)

func NewSourceManger() types.SourceManger {
	return &sourceManger{}
}

type sourceManger struct {
	m sync.Map
}

func (s *sourceManger) Resolver(id string) types.Source {
	it, ok := s.m.Load(id)
	if !ok {
		return nil
	}
	return it.(types.Source)
}

func (s *sourceManger) Append(id string, source types.Source) types.SourceManger {
	s.m.Store(id, source)
	return s
}
func (s *sourceManger) Remove(id string) types.SourceManger {
	s.m.Delete(id)
	return s
}
func (s *sourceManger) Close() error {
	s.m.Range(func(key, value interface{}) bool {
		source := value.(types.Source)
		source.Close()
		return true
	})
	s.m = sync.Map{}
	return nil
}
