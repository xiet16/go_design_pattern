package decorator_pattern

/*
装饰器模式:

用法理解: 接口继承类中可以任意组合使用同类接口继承类

*/

//装饰器行为接口，这个还不适合暴露给外部使用，
type DecoratorHandler interface {
	Show()
}

//这里我仿造了HTTP路由HandleFunc 的注入方式定义结构体
type DecoratorEntry struct {
	handle DecoratorHandler
	Name   string
}

//我要约束它有Show 的方法，有Add , Remove 方法
type DecoratorComponent struct {
	md    []DecoratorEntry
	entry DecoratorEntry
}

func (d *DecoratorComponent) Show() {
	for _, entry := range d.md {
		entry.handle.Show()
	}
}

func (d *DecoratorComponent) Add(entry DecoratorEntry) {

}

func (d *DecoratorComponent) Remove(entry DecoratorEntry) {

}
