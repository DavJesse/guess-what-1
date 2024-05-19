package main

import (
	"math"
	"testing"
)

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

func TestVariance(t *testing.T) {
	intSlc := []int{1, 2, 3, 4, 5, 6, 7}
	mean := average(intSlc)
	got := math.Round(variance(intSlc, mean))
	expected := float64(5)

	if got != expected {
		t.Errorf("Got: %f", got)
		t.Errorf("Expected: %f", expected)
		t.Error("TestVariance Failed")
	}
}

// func testStandardDeviation(*testing) {
// }
