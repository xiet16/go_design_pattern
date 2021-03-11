package strategy_pattern

func StrategyTest() {
	add := &StrategySub{}
	context := &StrategyContext{}
	context.AddStrategyHandler(add)

	context.ExcuteStrategy()
}
