package main

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	mathRand "math/rand"
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

func myInit() int64 {
	var b [8]byte
	_, err := cryptoRand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	return int64(binary.LittleEndian.Uint64(b[:]))
}

func Generate(k int) []byte {
	mathRand.Seed(myInit())

	str := make([]byte, 0, k)
	for i := 0; i < k; i++ {
		str = append(str, uint8(mathRand.Intn(3)+97))
	}
	return str
}

func TestFunc(t *testing.T) {
	n := 100
	k := 500
	for i := 0; i < n; i++ {
		str := Generate(k)
		if minLenString(k, str) != slowMinLenString(k, str) {
			t.Error("Error")
		}
	}
}

func slowMinLenString(n int, data []byte) int {
	min := 200_001
	for l := 0; l < n-1; l++ {
		for r := l + 1; r < n; r++ {
			if min > (r-l+1) && isHereEverySymbol(data[l:r+1]) {
				min = r - l + 1
			}
		}
	}
	if min == 200_001 {
		return -1
	} else {
		return min
	}
}

func isHereEverySymbol(data []byte) bool {
	var a, b, c, d int
	for _, v := range data {
		if v == byte('a') {
			a++
		}
		if v == byte('b') {
			b++
		}
		if v == byte('c') {
			c++
		}
		if v == byte('d') {
			d++
		}
	}
	if a*b*c*d == 0 {
		return false
	} else {
		return true
	}
}
