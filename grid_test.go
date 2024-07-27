package main

import (
	"testing"
)

type TestData struct {
	Input    Vec2
	Expected []Vec2
}

func TestGetNeighbors(t *testing.T) {
	testData := []TestData{
		{
			Input: Vec2{X: 0, Y: 0},
			Expected: []Vec2{
				{X: 1, Y: 0},
				{X: 0, Y: 1},
			},
		},
		{
			Input: Vec2{X: 2, Y: 1},
			Expected: []Vec2{
				{X: 1, Y: 1},
				{X: 3, Y: 1},
				{X: 2, Y: 0},
				{X: 2, Y: 2},
			},
		},
	}

	for _, tv := range testData {
		neighbors := GetNeighbors(tv.Input)

		if !isEqual(neighbors, tv.Expected) {
			t.Fatalf("Neighbors is not the same as Expected.\nNeighbors: %v\nExpected: %v\n", neighbors, tv.Expected)
		}
	}
}

func isEqual(a, b []Vec2) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
