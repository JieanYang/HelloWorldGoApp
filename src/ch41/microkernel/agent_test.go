package microkernel

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

type DemoCollector struct {
	evtReceiver EventReceiver
	agtCtx      context.Context
	stopChan    chan struct{}
	name        string
	content     string
}

func NewCollect(name string, content string) *DemoCollector {
	return &DemoCollector{
		stopChan: make(chan struct{}),
		name:     name,
		content:  content,
	}
}

func (c *DemoCollector) Init(evtReceiver EventReceiver) error {
	fmt.Println("initialize collector", c.name)
	c.evtReceiver = evtReceiver
	return nil
}

func (c *DemoCollector) Start(agtCtx context.Context) error {
	fmt.Println("start collector", c.name)
	for {
		select {
		case <-agtCtx.Done():
			c.stopChan <- struct{}{}
			break
		default:
			time.Sleep(time.Millisecond * 50)
			c.evtReceiver.OnEvent(Event{c.name, c.content})
			// 开启 Collector 相应的功能：网站服务器
		}
	}
}

func (c *DemoCollector) Stop() error {
	fmt.Println("stop collector", c.name)
	select {
	case <-c.stopChan:
		return nil
	case <-time.After(time.Second * 1):
		return errors.New("failed to stop for timeout")
	}
}

func (c *DemoCollector) Destory() error {
	fmt.Println(c.name, "released resources.")
	// 类似资源释放，删除临时数据或文件夹
	return nil
}

func TestAgent(t *testing.T) {
	agt := NewAgent(100)
	c1 := NewCollect("c1", "1")
	c2 := NewCollect("c2", "2")
	agt.RegisterCollector("c1", c1)
	agt.RegisterCollector("c2", c2)
	if err := agt.Start(); err != nil {
		fmt.Printf("start error %v\n", err)
	}
	fmt.Println(agt.Start())
	time.Sleep(time.Second * 1)
	agt.Stop()
	agt.Destory()
	fmt.Println(agt.collectors)
}
