package strategy_pattern

import (
	"errors"
)

/*
策略模式
理解: 策略模式是一种把实现、创建、调用 分开的一种编程模式，由使用者创建、注入以及调用，由服务提供方提供一系列实现方案供用户选择
*/

//接口
type StrategyHandler StrategyOperation

type StrategyOperation interface {
	Count(a, b int) (interface{}, error)
}

//具体实现类
type StrategyAdd struct{}

func (s *StrategyAdd) Count(a, b int) (interface{}, error) {
	return a + b, nil
}

type StrategySub struct{}

func (s *StrategySub) Count(a, b int) (interface{}, error) {
	return a - b, nil
}

type StrategyDivided struct{}

func (s *StrategyDivided) Count(a, b int) (interface{}, error) {
	if b == 0 {
		return 0, errors.New("被除数不能为0")
	}
	return a / b, nil
}

//
type StrategyContext struct {
	Handler StrategyHandler
}

func (c *StrategyContext) AddStrategyHandler(h StrategyHandler) {
	c.Handler = h
}

func (c *StrategyContext) ExcuteStrategy() (interface{}, error) {
	if c.Handler == nil {
		return nil, errors.New("handler is null")
	}

	result, err := c.Handler.Count(0, 1)
	return result, err
}
