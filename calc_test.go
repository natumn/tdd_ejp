package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	calcStrCase1 := "15*5"
	calcStrCase2 := "6+1"
	answerCase1 := calc(calcStrCase1)
	answerCase2 := calc(calcStrCase2)

	if answerCase1 != 75 {
		t.Fatal("miss calculator.")
	}
	t.Log(calcStrCase1, "=", answerCase1)
	if answerCase2 != 7 {
		t.Fatal("miss calculator.")
	}
	t.Log(calcStrCase2, "=", answerCase2)
}
