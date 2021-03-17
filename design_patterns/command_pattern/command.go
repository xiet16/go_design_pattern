package command_pattern

import (
	"errors"
	"fmt"
	"sync"
)

/*
命令模式：通过中间件访问具体服务
理解： 这个模式主要是客户端和服务器之间的解耦，个人认为和外观模式相似，外观模式在是客户端的组合，命令模式在添加的中间层组合
       外观模式是客户端直接调用对外提供的接口，命令模式是客户端需要设置命令接口，但是客户端不知道具体服务提供的对外接口
	   一个命令模式对应一个具体的服务实现
*/

type RealServer interface {
	Do()
}

type AServer struct{}

func (a *AServer) Do() {
	fmt.Println("A server do")
}

type BServer struct{}

func (a *BServer) Do() {
	fmt.Println("B server do")
}

func NewACommandMiddleware() *CommandMiddleware {
	return &CommandMiddleware{
		s: &AServer{},
	}
}

type ICommandMiddleware interface {
	Excute()
}
type CommandMiddleware struct {
	s RealServer
}

func (m *CommandMiddleware) Excute() {
	if m.s != nil {
		fmt.Println("command middleware do")
		m.s.Do()
	}
}

type User struct {
	m ICommandMiddleware //这里可以设置为map, 懒得写了
	sync.Mutex
}

func NewUser() *User {
	return &User{}
}

func (u *User) SetCommand(c ICommandMiddleware) {
	u.Mutex.Lock()
	defer u.Mutex.Unlock()
	u.m = c
}

func (u *User) Invoke() error {
	fmt.Println("User invoke")
	if u.m == nil {
		return errors.New("please set command before invoke")
	}

	u.m.Excute()
	return nil
}
