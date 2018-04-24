package main

import (
	"testing"
)

func TestStringRpt(t *testing.T) {
	string := StringRpt("Go", 3)
	if string != "GoGoGo" {
		 t.Errorf("String was incorrect, got: %s, want: %s.", string, "GoGoGo")
	}
}

func TestStringHasPrefix(t *testing.T) {
	string := StringHasPrefix("PartyTime", "Party")
	if string != true {
		 t.Errorf("Incorrect, got: %t, want: %t.", string, true)
	}
}
