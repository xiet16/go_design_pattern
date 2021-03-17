package main

import (
	"fmt"

	"github.com/xiet16/go_design_patterns/design_patterns/command_pattern"
	"github.com/xiet16/go_design_patterns/design_patterns/observer_pattern"
)

func main() {
	fmt.Println("welcome to design pattern")
	// strategy_pattern.StrategyTest()
	// decorator_pattern.DecoratorTest()
	//proxy_pattern.ProxyTest()
	//adapter_pattern.AdapterTest()
	command_pattern.CommandTest()
	observer_pattern.ObserverTest()
}
