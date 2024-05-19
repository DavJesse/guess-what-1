package main

import "testing"

func TestHasSuffix(t *testing.T) {
	str := "String"
	subStr := "ing"
	got := hasSuffix(str, subStr)
	expected := true

	if got != expected {
		t.Errorf("Got: %t", got)
		t.Errorf("Expected: %t", expected)
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
