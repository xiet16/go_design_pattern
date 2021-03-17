package state_pattern

/*
状态模式： 状态模式是一种将状态和对应的行为封装在同一个类的模式。主要为了解决多条件判断的扩展
关键点: 既然是为了解决多个if,switch 的问题，那它们之间是怎么调用的呢？ 答案就是各个状态得知道其他所有的状态情况，也就是各个状态间耦合了，这个模式违反了开闭原则
*/

type IScoreState interface {
	CheckScore()
}

//低中高分段检测
type LowScoreState struct{}

func (s *LowScoreState) CheckScore() {

}

type MiddleScoreState struct{}

func (s *MiddleScoreState) CheckScore() {

}

type HighScoreState struct{}

func (s *HighScoreState) CheckScore() {

}
