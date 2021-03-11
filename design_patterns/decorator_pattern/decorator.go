package decorator_pattern

import "fmt"

/*
装饰器模式:

用法理解: 接口继承类中可以任意组合使用同类接口继承类

*/

//装饰器行为接口，这个还不适合暴露给外部使用，
type DecoratorHandler interface {
	Show()
}

type ComponenentAct interface {
	Add(entry DecoratorEntry)
	Remove(entry DecoratorEntry)
}

//这里我仿造了HTTP路由HandleFunc 的注入方式定义结构体
type DecoratorEntry struct {
	handle DecoratorHandler
	Name   string
}

func (d *DecoratorEntry) Show() {
	fmt.Println(d.Name, "entry show")
}

var DecoratorEntryHandler DecoratorComponent

//我要约束它有Show 的方法，有Add , Remove 方法
type DecoratorComponent struct {
	Name  string
	md    []*DecoratorEntry
	Entry *DecoratorEntry
	act   ComponenentAct
}

func (d *DecoratorComponent) Show() {
	fmt.Println(d.Name, "decorator show")
	for _, entry := range d.md {
		entry.Show()
	}
}

func (d *DecoratorComponent) Add(name string, entry DecoratorHandler) {
	d.md = append(d.md, &DecoratorEntry{
		Name:   name,
		handle: entry,
	})
}

func (d *DecoratorComponent) Remove(entry *DecoratorHandler) {

}
