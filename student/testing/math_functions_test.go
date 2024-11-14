package testing

import (
	"math"
	"testing"

	"student/format"
	maths "student/maths"
)

func TestAverage(t *testing.T) {
	intSlc := []int{1, 2, 3}
	got := maths.Average(intSlc)
	expected := float64(2)

	// Compare 'got' and 'expected'
	if got != expected {
		t.Errorf("Got: %f", got)
		t.Errorf("Expected: %f", expected)
		t.Errorf("TestAverage Failed!")
	}
}

func TestMedian(t *testing.T) {
	intSlc := []int{1, 2, 3}
	got := maths.Median(intSlc)
	expected := float64(2)

	// Compare 'got' and 'expected'
	if got != expected {
		t.Errorf("Got: %f", got)
		t.Errorf("Expected: %f", expected)
		t.Error("TestMedian Failed")
	}
}

func TestVariance(t *testing.T) {
	intSlc := []int{1, 2, 3, 4, 5, 6, 7}
	mean := maths.Average(intSlc)
	got := math.Round(maths.Variance(intSlc, mean))
	expected := float64(4)

	// Compare 'got' and 'expected'
	if got != expected {
		t.Errorf("Got: %f", got)
		t.Errorf("Expected: %f", expected)
		t.Error("TestVariance Failed")
	}
}

func TestStandardDeviation(t *testing.T) {
	variance := float64(9)
	got := maths.StandardDeviation(variance)
	expected := float64(3)

	// Compare 'got' and 'expected'
	if got != expected {
		t.Errorf("Got: %f", got)
		t.Errorf("Expected: %f", expected)
		t.Error("TestStandardDeviation Failed")
	}
}

func TestSquare(t *testing.T) {
	num := float64(2)
	got := maths.Square(num)
	expected := float64(4)

	// Compare 'got' and 'expected'
	if got != expected {
		t.Errorf("Got: %f", got)
		t.Errorf("Expected: %f", expected)
		t.Error("TestSquare Failed")
	}
}

func TestSortSlice(t *testing.T) {
	testSlc := []int{4, 2, 3, 1}
	expected := []int{1, 2, 3, 4}
	got := format.SortSlice(testSlc)

	// Range over slices, comparing every element at ever index
	for i := 0; i < len(expected); i++ {
		if got[i] != expected[i] || len(got) != len(expected) {
			t.Errorf("Got: %d", got)
			t.Errorf("Expected: %d", expected)
			t.Errorf("TestSortSlice Failed!")
		}
	}
}

func TestRemoveOutlier_SmallData(t *testing.T) {
	subject := []int{1, 2, 3, 4, 5}
	got := maths.RemoveOutlier(subject)
	expected := []int{1, 2, 3, 4, 5}

	var g, e int

	for g < len(got) && e < len(expected) {
		if got[g] != expected[e] {
			t.Errorf("Got: %d", got)
			t.Errorf("Expected: %d", expected)
			t.Error("TestRemoveOutlier Failed!")
		}
		g++
		e++
	}
}

func TestRemoveOutlier_BigDataNoOutlier(t *testing.T) {
	subject := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	got := maths.RemoveOutlier(subject)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	var g, e int

	if len(got) != len(expected) {
		t.Error("TestRemoveOutlier Failed!")
		t.Errorf("Expected slice legth of: %d, got: %d\n", len(expected), len(got))
		return
	}

	for g < len(got) && e < len(expected) {
		if got[g] != expected[e] {
			t.Errorf("Got: %d", got)
			t.Errorf("Expected: %d", expected)
			t.Error("TestRemoveOutlier Failed!")
		}
		g++
		e++
	}
}

func TestRemoveOutlier_BigDataWithOutlier(t *testing.T) {
	subject := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12000, 13, 14, 15, 16}
	got := maths.RemoveOutlier(subject)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15, 16}

	var g, e int

	if len(got) != len(expected) {
		t.Error("TestRemoveOutlier Failed!")
		t.Errorf("Expected slice legth of: %d, got: %d\n", len(expected), len(got))
		return
	}

	for g < len(got) && e < len(expected) {
		if got[g] != expected[e] {
			t.Errorf("Got: %d", got)
			t.Errorf("Expected: %d", expected)
			t.Error("TestRemoveOutlier Failed!")
		}
		g++
		e++
	}
}