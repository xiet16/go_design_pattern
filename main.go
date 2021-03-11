package main

import (
	"fmt"

	"github.com/xiet16/go_design_patterns/design_patters/decorator_pattern"
	"github.com/xiet16/go_design_patterns/design_patters/strategy_pattern"
)

func main() {
	fmt.Println("welcome to design pattern")
	strategy_pattern.StrategyTest()
	decorator_pattern.DecoratorTest()
}
