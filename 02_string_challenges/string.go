package main

import (
	"strings"
)

func stringRpt(value string, repetitions int) string {
  return strings.Repeat(value, repetitions)
}

// String Challenge #1
func stringHasPrefix(value string, prefix string) bool {
	return strings.HasPrefix(value, prefix)
}

// String Challenge #2
func stringSplit(a, b string) []string {
	return strings.SplitAfter(a,b)
}