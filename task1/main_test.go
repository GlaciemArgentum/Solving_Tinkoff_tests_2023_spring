package main

import (
	"fmt"
	"testing"
)

type Test struct {
	in  []int
	res bool
}

var tests = []Test{
	{[]int{1, 2, 3, 4}, true},
	{[]int{1, 1, 3, 4}, true},
	{[]int{1, 2, 4, 4}, true},
	{[]int{1, 2, 2, 4}, true},
	{[]int{4, 3, 2, 1}, true},
	{[]int{4, 4, 2, 1}, true},
	{[]int{4, 3, 1, 1}, true},
	{[]int{4, 3, 3, 1}, true},
	{[]int{4, 4, 4, 4}, true},
	{[]int{4, 4, 4, 1}, true},
	{[]int{1, 1, 1, 4}, true},
	{[]int{1, 2, 3, 2}, false},
	{[]int{4, 3, 2, 5}, false},
}

func TestFunc(t *testing.T) {
	for i, test := range tests {
		if res := isQueueInFormat(test.in); res != test.res {
			t.Error(fmt.Sprintf("Test %d: expected %v, got %v", i+1, test.res, res))
		}
	}
}
