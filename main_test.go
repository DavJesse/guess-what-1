package main

import "testing"

func TestAverage(t *testing.T) {
	intSlc := []int{1, 2, 3}
	got := average(intSlc)
	expected := float64(2)

	if got != expected {
		t.Errorf("Got: %f", got)
		t.Errorf("Expected: %f", expected)
		t.Errorf("TestAverage Failed!")
	}
}

func TestMedian(t *testing.T) {
	intSlc := []int{1, 2, 3}
	got := median(intSlc)
	expected := float64(2)

	if got != expected {
		t.Errorf("Got: %f", got)
		t.Errorf("Expected: %f", expected)
		t.Error("TestMedian Failed")
	}
}

// func testVariance(*testing) {
// }

// func testStandardDeviation(*testing) {
// }
