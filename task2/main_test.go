package main

import (
	"fmt"
	"testing"
)

type Test struct {
	n   int
	m   int
	k   int
	res int
}

var tests = []Test{
	{3, 2, 2, 3},
	{7, 3, 2, 5},
	{1, 1, 1, 1},
	{2, 1, 1, 2},
	{8, 8, 2, 2},
	{10000, 5001, 1, 2},
	{10000, 10000, 10000, 10000},
	{10000, 4999, 1, 3},
}

func TestFunc(t *testing.T) {
	for i, test := range tests {
		if res := timeReview(test.n, test.m, test.k); res != test.res {
			t.Error(fmt.Sprintf("Test %d: expected %d, got %d", i+1, test.res, res))
		}
	}
}
