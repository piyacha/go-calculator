package calculator

import (
	"testing"
)

func TestCalculator(t *testing.T) {

	test_case := map[string]float64{
		"1.1 + 1": 2.1,
		"2 - 2": 0,
		"1 - 100": -99,
		"10 * 10": 100,
		"144 / 12": 12,
		"1 + 2 * 3 - 1": 6,
		"1 + 2 * ( 3 - 1 )": 5,
		"( 1 + 2 ) * 3 - 1": 8,
		"122 * ( ( 5 + 78 ) * 5 ) / 2 - 99": 25216,
	}

	for expression, answer := range test_case {
		c := Calculator { calculatorType: InfixCalculator }
		result := c.calculate(expression)

		if result != answer {
			t.Errorf("Expression %q = %f but got %f", expression, answer, result)
		}
	}
}
