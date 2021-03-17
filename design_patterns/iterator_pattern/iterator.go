package iterator_pattern

/*
迭代器模式: 通过提供一个对象来访问聚合对象中的一系列数据，而不是暴露聚合对象内部的表示
理解:迭代器分离了数据对象和对数据对象的行为
例子:典型的例子就是list, list 可以是一种存储某类对象集合的数据结构，这类结构的可能会很多的业务行为，如果在list这个类中封装行为，
    一个是不方便扩展，另外一个是暴露了它的内部表示

结构:
聚合类:数据和迭代器的封装，并且提供了对数据的增删改查接口
抽象聚合类:面向接口编程
迭代器具体实现类:提供了对数据的行为接口
抽象迭代器：面向对象编程
*/

type ICollectionAggregate interface {
	Add(l interface{})
	Remove(l interface{})
	Foreach()
}

type ListAggregate struct {
}

type Iterator interface{}

type ListIterator struct {
}
