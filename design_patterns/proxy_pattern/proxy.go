package proxy_pattern

import "fmt"

/*
代理模式:给某一对象提供代理对象
理解：代理模式的代理实现核心实际还是对原对象的调用，只是在处理原来逻辑前后，有自己的业务逻辑，是一种扩展模式。
意义：对原有逻辑或功能的扩展
*/

type TransportHandler interface {
	GoOut()
}

type TransportByCar struct {
	Name string
}

func (t *TransportByCar) GoOut() {
	fmt.Println("乘汽车出行")
}

type TransportByVehicle struct {
	Name string
}

func (t *TransportByVehicle) GoOut() {
	fmt.Println("骑自行车出行")
}

type TransportByRailWay struct {
	Name string
}

func (t *TransportByRailWay) GoOut() {
	fmt.Println("乘火车出行")
}

type TransportProxy struct {
	Name string
	TransportHandler
}

func (t *TransportProxy) GoOut() {
	fmt.Println("租车代理行代理租车出行")
	fmt.Println("租车代理行出行前先约车")
	t.TransportHandler.GoOut()
	fmt.Println("客户到达目的地,代理完成")
}
