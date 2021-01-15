package calculator

type CalculatorType int

const (
	InfixCalculator CalculatorType = iota
	OtherCalculator
)

func (t CalculatorType) String() string {
  return [...]string{"InfixCalculator", "OtherCalculator"}[t]
}

type Calculator struct {
	calculatorType CalculatorType
}

func (cal Calculator) calculate(input string) (answer float64) {

	switch cal.calculatorType {
	case InfixCalculator:
		var c InfixCalculation
		answer = c.execute(input)
	case OtherCalculator:
		// Any others calculator
	}

	return
}
