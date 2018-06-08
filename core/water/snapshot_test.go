package water

import (
	"github.com/vlorc/hprose-gateway/core/types"
	"testing"
	"time"
)

type beanSource struct {
	service *types.Service
}

func (c *beanSource) SetService(service *types.Service) {
	c.service = service
}

func (c *beanSource) Reset() types.Source {
	c.service = nil
	return c
}

func (*beanSource) Close() error {
	return nil
}

func Test_Snapshot(t *testing.T) {
	water, query := NewSnapshotWater(NewChannelWater(100), func() types.Source {
		return &beanSource{}
	})
	water.Push([]types.Update{{Op: types.Add, Id: "/user/1", Service: &types.Service{
		Id:       "1",
		Name:     "user",
		Version:  "1.0.0",
		Url:      "http://localhost:8080",
		Platform: "1",
		Meta: map[string]interface{}{
			"appid": 1,
			"key":   "123",
		},
	}}})

	time.Sleep(time.Second * 1)
	t.Log(query("/user/1"))
	time.Sleep(time.Second * 1)
}
