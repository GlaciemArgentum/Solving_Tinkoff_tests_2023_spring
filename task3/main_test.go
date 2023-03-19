package main

import (
	"fmt"
	"testing"
)

type Test struct {
	n   int
	s   []byte
	res int
}

var tests = []Test{
	{12, []byte("aabbccddbadd"), 5},
	{16, []byte("aaaabbbbccccdddd"), 10},
	{7, []byte("dbbccca"), 7},
	{7, []byte("abcabac"), -1},
	{0, []byte(""), -1},
	{1, []byte("a"), -1},
	{2, []byte("ab"), -1},
	{3, []byte("abc"), -1},
	{4, []byte("abcd"), 4},
	{4, []byte("abcb"), -1},
	{12, []byte("abcabcabcabc"), -1},
	{12, []byte("abcabcabcabd"), 4},
	{12, []byte("abbbbbcccccd"), 12},
	{18, []byte("aaaaabcaaaadbbbbbb"), 7},
	{12, []byte("aaaaaaaaaaaa"), -1},
}

func TestFunc(t *testing.T) {
	for i, test := range tests {
		if res := minLenString(test.n, test.s); res != test.res {
			t.Error(fmt.Sprintf("Test %d: expected %d, got %d", i+1, test.res, res))
		}
	}
}
