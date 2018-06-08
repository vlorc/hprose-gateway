package etcd

import (
	"encoding/json"
	"github.com/vlorc/hprose-gateway/core/types"
	"time"
)

func NewPrintWatcher(output func(string, ...interface{})) types.NamedWatcher {
	return printWatcher(output)
}

func (p printWatcher) Push(event []types.Update) error {
	for i := range event {
		buf, _ := json.MarshalIndent(&event[i], "", "    ")
		println("[", time.Now().Format("2006-01-02 15:04:05"), "][", event[i].Op.String(), "]: ", string(buf))
	}
	return nil
}

func (p printWatcher) Pop() ([]types.Update, error) {
	return nil, nil
}

func (p printWatcher) Close() error {
	return nil
}
