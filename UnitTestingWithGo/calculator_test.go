package main

import "testing"

func TestCalculator(t *testing.T) {
	value := 2
	result := Calculate(value)
	if result != 4 {
		t.Errorf("Expected %d+2 = 4 but got %v", value, result)
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input, expected int
	}{
		{2,4},
		{-1, 1},
		{0, 2},
		{9999, 10001},
		{3, 5},
	}
	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Errorf("Test failed. Expected %d but found %d", test.expected, output)
		}
	}
}
