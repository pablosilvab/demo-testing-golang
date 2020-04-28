package main

import "testing"

type TestTable struct {
	input1   int
	input2   int
	expected int
}

func TestSum(t *testing.T) {
	if Sum(5, 5) != 10 {
		t.Error("Se esperaba que 5 + 5 = 10")
	}
}

func TestTableSum(t *testing.T) {
	tests := []TestTable{
		{1, 1, 2},
		{1, 2, 3},
		{1, 6, 7},
		{1, -11, -10},
	}
	for _, test := range tests {
		if output := Sum(test.input1, test.input2); output != test.expected {
			t.Error("Test Fail {} input1: {} input2: {} - Expected: {} , Received: {}", test.input1, test.input2, test.expected, output)
		}
	}

}
