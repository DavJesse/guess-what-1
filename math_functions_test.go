package main

import (
	"math"
	"testing"
)

// Test average function
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

// Test median function
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

// Test variance function
func TestVariance(t *testing.T) {
	intSlc := []int{1, 2, 3, 4, 5, 6, 7}
	mean := average(intSlc)
	got := math.Round(variance(intSlc, mean))
	expected := float64(4)

	if got != expected {
		t.Errorf("Got: %f", got)
		t.Errorf("Expected: %f", expected)
		t.Error("TestVariance Failed")
	}
}

// Test standardDeviation function
func TestStandardDeviation(t *testing.T) {
	variance := float64(9)
	got := standardDeviation(variance)
	expected := float64(3)

	if got != expected {
		t.Errorf("Got: %f", got)
		t.Errorf("Expected: %f", expected)
		t.Error("TestStandardDeviation Failed")
	}
}

// Test square function
func TestSquare(t *testing.T) {
	num := float64(2)
	got := square(num)
	expected := float64(4)

	if got != expected {
		t.Errorf("Got: %f", got)
		t.Errorf("Expected: %f", expected)
		t.Error("TestSquare Failed")
	}
}

// Test sortSlice function
func TestSortSlice(t *testing.T) {
	testSlc := []int{4, 2, 3, 1}
	expected := []int{1, 2, 3, 4}
	got := sortSlice(testSlc)

	//Range over slices, comparing every element at ever index
	for i := 0; i < len(expected); i++ {
		if got[i] != expected[i] || len(got) != len(expected) {
			t.Errorf("Got: %d", got)
			t.Errorf("Expected: %d", expected)
			t.Errorf("TestSortSlice Failed!")
		}
	}
}
