

package calculator

import (
	"regexp"
	"strings"
	"strconv"
)

type InfixCalculation struct {
	numbers		[]float64
	operators []string
}

func (c *InfixCalculation) eval() (result float64) {
	var x, y float64
	y = c.popNumber()
	x = c.popNumber()

	switch c.popOperator() {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		result = x / y
	}

	if (len(c.operators) > 0) && (c.lastOperator() == "(") {
		c.popOperator()
	}

	return result
}

func (c *InfixCalculation) pushToNumbers(i float64) {
	c.numbers = append(c.numbers, i)
}

func (c *InfixCalculation) pushToOperators(i string) {
	c.operators = append(c.operators, i)
}

func (c *InfixCalculation) popOperator() (operator string) {
	operator = c.operators[len(c.operators) - 1]
	c.operators = c.operators[:len(c.operators) - 1]
	return
}

func (c *InfixCalculation) popNumber() (number float64) {
	number = c.numbers[len(c.numbers) - 1]
	c.numbers = c.numbers[:len(c.numbers) - 1]
	return
}

func (c InfixCalculation) operatorPriority(i string) (value uint8) {
	switch i {
	case "+", "-":
		value = 1
	case "*", "/":
		value = 2
	case "(", ")":
		value = 3
	}
	return
}

func (c InfixCalculation) lastOperator() (value string) {
	value = c.operators[len(c.operators)-1]
	return
}

func (c InfixCalculation) canPushToOperators(input string) (result bool) {
	if len(c.operators) == 0 {
		return false
	}

	if c.lastOperator() == "(" {
		return true
	}

	if input == ")" {
		return false
	}

	result = c.operatorPriority(input) > c.operatorPriority(c.lastOperator())
	return
}

func (c InfixCalculation) execute(input string) (result float64) {
	var inputs = strings.Split(input, " ")

	for i := 0; i < len(inputs); i++ {

		isNumber, _ := regexp.MatchString("[0-9]", inputs[i])
		isOperator, _ := regexp.MatchString("[\\*\\/()+-]", inputs[i])

		if isNumber {
			a, _ := strconv.ParseFloat(inputs[i], 64)
			c.pushToNumbers(a)
		} 

		if isOperator {
			if len(c.operators) > 0 && !c.canPushToOperators(inputs[i]) {
				c.pushToNumbers(c.eval())
			}

			if inputs[i] != ")" {
				c.pushToOperators(inputs[i])
			}
		}
	}

	for {
		result = c.eval()
		c.pushToNumbers(result)
		if (len(c.numbers) == 1){
			break
		}
	}

	return
}
