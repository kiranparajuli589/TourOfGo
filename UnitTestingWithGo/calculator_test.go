package main

import (
	"io/ioutil"
	"testing"
)

type AddResult struct {
	x, y, expected int
}

var addResults = []AddResult{
	{1,1,2},
	{1,-1,0},
	{0,0,0},
	{41,41,82},
}

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

func TestAdd(t *testing.T) {
	for _, test := range addResults {
		result := Add(test.x, test.y)
		if result != test.expected {
			t.Fatalf("Test failed. Expected %d but found %d", test.expected, result)
		}
	}
}

func TestReadFile(t *testing.T) {
	expected := "hello world\n"
	data, err := ioutil.ReadFile("testData/test.data")
	if err != nil {
		t.Fatal("Could not open file!")
	}
	if string(data) != expected {
		t.Fatalf("Test failed: Expected %q but found %q", expected, string(data))
	}
}
