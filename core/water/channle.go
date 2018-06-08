package water

import (
	"github.com/vlorc/hprose-gateway/core/types"
	"io"
)

func NewChannelWater(length int) types.NamedWatcher {
	return ChannelWater(make(chan []types.Update, length))
}

func (c ChannelWater) Push(event []types.Update) error {
	c <- event
	return nil
}

func (c ChannelWater) Pop() ([]types.Update, error) {
	event, ok := <-c
	if !ok {
		return event, io.EOF
	}
	return event, nil
}

func (c ChannelWater) Close() error {
	close(c)
	return nil
}
