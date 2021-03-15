package visitor_pattern

/*
访问者模式:将作用于数据结构的中各元素独立分离出来封装为独立的类
理解:装饰器模式体现的是，同类继承间相互注入使用，它们都继承自一个基类。访问者模式注重分离成继承自不同的基类
例子：装饰器中，假如要装扮一个人，衣服子类中，可以假如上衣，鞋子等等子类，这些类都继承自服饰类，
	 而访问者，假如某在线教学网站，学生继承自各类角色类(vip,普通游客，vvip)，下载分为各个方式下载（百度云盘，这能在线观看等），
	 学生角色和下载继承自不同的基类，但他们都相互依赖
接口具体的实现还是在visitor 的实现类中执行
*/

//接口定义，相互依赖
type IVisitor interface {
	Visit(d IDownload)
}

type IDownload interface {
	Download(v IVisitor)
}

type IWatch interface {
	Play(v IVisitor)
}

//学习网站类（对象结构类），违反了开闭原则，如果下次有其他功能，这里得改(比如在线看视频，这里网站有接口了，但visitor 还没实现)
type LearnWeb struct {
	Vistor    IVisitor
	Downloads []IDownload
	Watchs    []IWatch
}

//下面是实现类
type VipStudent struct{}

func (s *VipStudent) Visit(d IDownload) {

}

type VVipStudent struct{}

func (s *VVipStudent) Visit(d IDownload) {

}

type NormalStudent struct{}

func (s *NormalStudent) Visit(d IDownload) {

}
