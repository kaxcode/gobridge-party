package main

import (
	"fmt"
)

func printingCamelids() {
	camelids := []string{"Camel", "Llama", "Guanaco", "Alpaca", "Vicuña"}
	for i, breed := range camelids {
		fmt.Println(i, breed)
	}

	// In Workshop will show how to iterrate and append to map and to slice.
}

// // Arrays Challenge #1
// func doubleNumbers(x []int) []int {
// 	// get a slice of integers, return a new slice with each value doubled.
// }

// // Arrays Challenge #2
// func studentSchedules() []string {
// 	schedules := []string{}

// 	students := []string{"Kenia", "Troy", "Kristen", "Paul Chin Jr."}
// 	classes := []string{"Chemistry I", "Calculus", "Software Systems", "Mathematical Foundations of Computing"}

// 	// iterate students and classes and append the results to the schedule

// 	// return the schedules slice
// }

// // Arrays Challenge #3
// func twoLargestNumbers(numbers []int) [2]int {
// 	// Sort the slics, grab the 2 largest numbers, but the 2 largest numbers in a new slice. Return new slice.
// 	// CLUE: Use the sort library https://golang.org/pkg/sort/
// }