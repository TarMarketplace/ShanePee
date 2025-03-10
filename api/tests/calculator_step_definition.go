package tests

import (
	"fmt"

	"github.com/cucumber/godog"
)

type Calculator struct {
	result int
}

var calc Calculator

func aCalculator() error {
	calc = Calculator{}
	return nil
}

func iAddAnd(arg1, arg2 int) error {
	calc.result = arg1 + arg2
	return nil
}

func theResultShouldBe(expected int) error {
	if calc.result != expected {
		return fmt.Errorf("expected %d, but got %d", expected, calc.result)
	}
	return nil
}

func RegisterCalculatorSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^a Calculator$`, aCalculator)
	ctx.Step(`^I add (\d+) and (\d+)$`, iAddAnd)
	ctx.Step(`^the result should be (\d+)$`, theResultShouldBe)
}
