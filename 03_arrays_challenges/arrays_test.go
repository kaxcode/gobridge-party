package main

import (
	"testing"
	"reflect"
)

// Arrays Challenge #1
func TestPrintingCamelids(t *testing.T) {
	// Write a test for the printingCamelids
}

func TestdoubleNumbers(t *testing.T) {
	actual := []int{1, 2, 3}
	expected := []int{2, 4, 6}

	if reflect.DeepEqual(actual, expected) {
		 t.Errorf("Array was incorrect, got: %p, want: %p.", actual, expected)
	}
}

func TestStudentClasses(t *testing.T) {
	d := studentSchedules()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}
}