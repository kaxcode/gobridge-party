package main

import (
	"fmt"
	"sort"
)

func printingCamelids() {
	camelids := []string{"Camel", "Llama", "Guanaco", "Alpaca", "Vicuña"}
	for i, breed := range camelids {
		fmt.Println(i, breed)
	}

	// In Workshop will show how to iterrate and append to map and to slice.
}

// Arrays Challenge #1
func doubleNumbers(x []int) []int {
	numbers := []int{}
	for _, num := range x {
		numbers = append(numbers, num*2)
	}

	return numbers
}

// // Arrays Challenge #2
func studentSchedules() []string {
	schedules := []string{}

	students := []string{"Kenia", "Troy", "Kristen", "Paul Chin Jr."}
	classes := []string{"Chemistry I", "Calculus", "Software Systems", "Mathematical Foundations of Computing"}

	for _, student := range students {
		for _, class := range classes {
			schedules = append(schedules, student + " has been added to " + class )
		}
	}

	return schedules
}

// Arrays Challenge #3
func twoLargestNumbers(numbers []int) [2]int {
  a, b := 0, 0
  for _, v := range numbers {
    if v > b {
      a, b = b, v
    } else if v > a {
      a = v
    }
  }
  return [2]int{a, b}
}