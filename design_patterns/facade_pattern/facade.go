package facade_pattern

import "fmt"

/*
外观模式：
理解：外观模式是一种对外提供下层各个类型接口的整合
区别代理模式：
代理模式是对实现类的扩展，旨在实现类的功能都要实现，可以在实现前后进行自己的业务处理。  外观模式整合需要的各种类型的接口，
*/

//假设有一个生活类app ，可以调用支付宝接口、京东支付、看爱奇艺视频、腾讯视频等等
type FacadeApp struct {
}

func (f *FacadeApp) FacadeStart() {
	ali := &AliPay{}
	ali.Pay()

	jd := &JingdongPay{}
	jd.Pay()

	aiqy := &AQiYi{}
	aiqy.Play()

	//比如暂时不支持腾讯视频
}

type AliPay struct{}

func (p *AliPay) Pay() {
	fmt.Println("支付宝支付")
}

type JingdongPay struct{}

func (p *JingdongPay) Pay() {
	fmt.Println("京东支付")
}

type AQiYi struct{}

func (s *AQiYi) Play() {
	fmt.Println("爱奇艺app播放视频")
}

type Tencent struct{}

func (*Tencent) Play() {

	fmt.Println("腾讯app播放视频")
}
