package main

import "testing"

func TestHasSuffix(t *testing.T) {
	str := "String"
	subStr := "Str"
	got := hasSuffix(str, subStr)
	expected := true

	if got != expected {
		t.Errorf("Got: %b", got)
		t.Errorf("Expected: %b", expected)
		t.Errorf("TestHasSuffix Failed!")
	}
}

// got := hasSuffix(str, subStr)
// 	expected := true

// 	if got != expected {
// 		t.Errorf("Got: %", got)
// 		t.Errorf("Expected: %", expected)
// 		t.Errorf("Test Failed!")
// 	}
