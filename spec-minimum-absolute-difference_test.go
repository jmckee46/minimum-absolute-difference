package main

import "testing"

func TestMinimumAbsoluteDifference1(t *testing.T) {
	arr := []int32{3, -7, 0}

	minDiff := minimumAbsoluteDifference(arr)

	if minDiff != 3 {
		t.Errorf("got %d instead of 3", minDiff)
	}
}

func TestMinimumAbsoluteDifference2(t *testing.T) {
	arr := []int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75}

	minDiff := minimumAbsoluteDifference(arr)

	if minDiff != 1 {
		t.Errorf("got %d instead of 1", minDiff)
	}
}

func TestMinimumAbsoluteDifference3(t *testing.T) {
	arr := []int32{1, -3, 71, 68, 17}

	minDiff := minimumAbsoluteDifference(arr)

	if minDiff != 3 {
		t.Errorf("got %d instead of 3", minDiff)
	}
}

func TestMinimumAbsoluteDifference4(t *testing.T) {
	arr := []int32{-2, 2, 4}

	minDiff := minimumAbsoluteDifference(arr)

	if minDiff != 2 {
		t.Errorf("got %d instead of 2", minDiff)
	}
}
