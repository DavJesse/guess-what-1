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

func TestSplitString(t *testing.T) {
	str := "Unit-Testing"
	sep := "-"
	got := splitString(str, sep)
	expected := []string{"Unit", "Testing"}

	for i := 0; i < len(expected); i++ {
		if got[i] != expected[i] || len(got) != len(expected) {
			t.Errorf("Got: %s", got)
			t.Errorf("Expected: %s", expected)
			t.Errorf("TestSplitString Failed!")
		}
	}
}

func TestItoa(t *testing.T) {
	num := 123
	got := Itoa(num)
	expected := "123"

	if got != expected {
		t.Errorf("Got: %s", got)
		t.Errorf("Expected: %s", expected)
		t.Errorf("TestItoa Failed!")
	}
}

// got := hasSuffix(str, subStr)
// 	expected := true

// 	if got != expected {
// 		t.Errorf("Got: %", got)
// 		t.Errorf("Expected: %", expected)
// 		t.Errorf("Test Failed!")
// 	}
