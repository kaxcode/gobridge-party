package main

import (
	"testing"
	"reflect"
)

// Arrays Challenge #1
func TestPrintingCamelids(t *testing.T) {
	// Write a test for the printingCamelids
}

// Test for Challenge #2
func TestDoubleNumbers(t *testing.T) {
	actual := doubleNumbers([]int{1, 2, 3})
	expected := []int{2, 4, 6}

	if !reflect.DeepEqual(actual, expected) {
		 t.Errorf("Slice  was incorrect, got: %d, want: %d.", actual, expected)
	}
}

// Test for Challenge #3
func TestStudentClasses(t *testing.T) {
	d := studentSchedules()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}
}

// Test for Challenge #4
func TestTwoLargestNumbers(t *testing.T) {
	actual := TwoLargestNumbers([]int{45, 100, 1459, 563, 10})

	expected := [2]int{ 563, 1459}

	if reflect.DeepEqual(actual, expected) {
		t.Errorf("Slice  was incorrect, got: %d, want: %d.", actual, expected)
	}
}