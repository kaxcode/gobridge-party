package main

import "testing"

func TestSum(t *testing.T) {
    total := Sum(5, 5)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}

func TestMultiply(t *testing.T) {
	total := Multiply(5, 5)
	if total != 25 {
		 t.Errorf("Multiply was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSubtraction(t *testing.T) {
	total := Subtraction(37, 5)
	if total != 32 {
		 t.Errorf("Subtraction was incorrect, got: %d, want: %d.", total, 32)
	}
}

func TestDivision(t *testing.T) {
	total := Division(30, 6)
	if total != 5 {
		 t.Errorf("Divison was incorrect, got: %d, want: %d.", total, 5)
	}
}