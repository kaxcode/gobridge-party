package main

import (
	"testing"
)

func TestStringRpt(t *testing.T) {
	string := stringRpt("Go", 3)
	if string != "GoGoGo" {
		 t.Errorf("String was incorrect, got: %s, want: %s.", string, "GoGoGo")
	}
}

// func TestStringHasPrefix(t *testing.T) {
// 	string := stringHasPrefix("PartyTime", "Party")
// 	if string != true {
// 		 t.Errorf("Incorrect, got: %t, want: %t.", string, true)
// 	}
// }

func TestStringSplit(t *testing.T) {
	result := stringSplit("Go-Party", "-")
	actual := []string{"Go-", "Party"}
	if result[0]!= actual[0] || result[1] != actual[1] {
		t.Error(
			"For input: ", []string{},
			"expected:", 0,
			"got:", result,
		)
	}
}
