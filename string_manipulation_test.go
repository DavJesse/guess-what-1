package main

import "testing"

func TestHasSuffix(t *testing.T) {
	str := "String"
	subStr := "ing"
	got := hasSuffix(str, subStr)
	expected := true

	// Compare 'got' and 'expected'
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

	// Compare 'got' and 'expected' at various indices
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

	// Compare 'got' and 'expected'
	if got != expected {
		t.Errorf("Got: %s", got)
		t.Errorf("Expected: %s", expected)
		t.Errorf("TestItoa Failed!")
	}
}
