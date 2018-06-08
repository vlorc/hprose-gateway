package water

import (
	"github.com/vlorc/hprose-gateway/core/driver"
	"github.com/vlorc/hprose-gateway/core/types"
	"io"
)

func NewSnapshotWater(
	in types.NamedWatcher,
	router types.NamedRouter,
	manager types.SourceManger,
	source types.Source,
	out ...types.NamedWatcher) types.NamedWatcher {
	snapshot := &SnapshotWater{
		NamedWatcher: in,
		router:       router,
		manager:      manager.Append("", source),
		water:        out,
	}
	snapshot.init()
	return snapshot
}

func (s *SnapshotWater) Close() error {
	err := s.NamedWatcher.Close()
	if nil != s.channel {
		close(s.channel)
	}
	for _, w := range s.water {
		w.Close()
	}
	return err
}

func (s *SnapshotWater) init() {
	go func() {
		for s.workerUpdates() != io.EOF {
		}
	}()
	if len(s.water) <= 0 {
		s.push = func(up types.Update) {}
		return
	}
	s.channel = make(chan types.Update, 64)
	go func() {
		for up := range s.channel {
			for _, v := range s.water {
				v.Push([]types.Update{up})
			}
		}
	}()
	s.push = func(up types.Update) {
		s.channel <- up
	}
}

func (s *SnapshotWater) getSource(service *types.Service) types.Source {
	factory := driver.Query(service.Driver)
	source := factory.Instance(s.router, s.manager)
	source.SetService(service)
	return source
}

func (c *SnapshotWater) workerUpdates() error {
	updates, err := c.Pop()
	if nil != err {
		return err
	}
	for _, up := range updates {
		source := c.manager.Resolver(up.Id)
		switch up.Op {
		case types.Add:
			if nil != source {
				source.SetService(up.Service)
			} else {
				source = c.getSource(up.Service)
				c.manager.Append(up.Id, source)
				c.router.Append(types.Named{Id: up.Service.Id, Path: up.Service.Path}, up.Id)
			}
			c.push(up)
		case types.Delete:
			if nil != source {
				c.router.Remove(types.Named{Id: source.Service().Id, Path: source.Service().Path}, up.Id)
				c.manager.Remove(up.Id)
				driver.Query(source.Service().Driver).Release(source)
				c.push(up)
			}
		}
	}
	return nil
}
