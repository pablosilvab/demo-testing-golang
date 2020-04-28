package main

import "testing"

func TestSum(t *testing.T) {
	if Sum(5, 5) != 10 {
		t.Error("Se esperaba que 5 + 5 = 10")
	}
}
