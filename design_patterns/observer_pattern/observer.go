package observer_pattern

import (
	"fmt"
	"sync"
)

/*
观察者模式：一对多依赖关系对象中，当一个对象发生变化时，依赖它的对象自动触发变化，又叫发布订阅模式
理解：发布订阅模式以前在多层架构中用过。个人认为，进程内的发布订阅和进程外的消息队列差不多，可以用来解耦调用方和被调用方的依赖关系。
      特别在多层系统层与层之间存在复杂调用关系时，用起来贼爽

本质用法：用一个监听类，负责缓存订阅者和发布者的关系以及当有消息发布时，触发通知订阅者。 我觉得rabbitmq 等消息队列也是用这个原理
*/

type IAction interface {
	Do()
}
type Obeserver struct {
	handlers map[string][]IAction
	Locker   sync.Locker
}

func NewObeserver() *Obeserver {
	return &Obeserver{
		handlers: make(map[string][]IAction), //主要是为了一对多
		Locker:   &sync.Mutex{},
	}
}

func (o *Obeserver) Add(htype string, h IAction) {
	o.Locker.Lock()
	defer o.Locker.Unlock()
	o.handlers[htype] = append(o.handlers[htype], h)
}

func (o *Obeserver) Run() {
	for _, hs := range o.handlers {
		for _, h := range hs {
			h.Do()
		}
	}
}

type Mom struct{}

func (m *Mom) Do() {
	fmt.Println("母亲抚慰")
}

type Children struct{}

func (c *Children) Do() {
	fmt.Println("婴儿哭")
}

type Cat struct{}

func (c *Cat) Do() {
	fmt.Println("猫追")
}

type Mouse struct{}

func (m *Mouse) Do() {
	fmt.Println("老鼠跑")
}
