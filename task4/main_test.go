package main

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"fmt"
	mathRand "math/rand"
	"testing"
)

func myInit() int64 {
	var b [8]byte
	_, err := cryptoRand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	return int64(binary.LittleEndian.Uint64(b[:]))
}

func Generate(n int) (int, []int) {
	mathRand.Seed(myInit())

	k := mathRand.Intn(n) + 2
	str := make([]int, 0, k)
	for i := 0; i < k; i++ {
		str = append(str, mathRand.Intn(20))
	}
	return k, str
}

func TestFunc(t *testing.T) {
	n := 10
	k := 50
	for i := 0; i < n; i++ {
		count, str := Generate(k)
		f := findPrefix(count, str)
		s := slowFindPrefix(count, str)
		if f != s {
			t.Error(fmt.Sprintf("Test %d: expected %d, got %d for %v", i+1, s, f, str))
		}
	}
}

func slowFindPrefix(n int, data []int) int {
	if n <= 3 {
		return n - 1
	}
	max := 3
	for r := 3; r < n; r++ {
		if isBoring(data[:r+1]) {
			max = r
		}
	}
	return max
}

func isBoring(data []int) bool {
	dataMap := make(map[int]int, len(data))
	for _, v := range data {
		dataMap[v]++
		if len(dataMap) > 2 {
			return false
		}
	}
	var (
		countes       = make([]int, 0, 2)
		countOfCounts = make([]int, 0, 2)
	)
	for key, v := range dataMap {
		countes = append(countes, key)
		countOfCounts = append(countOfCounts, v)
	}
	if len(dataMap) == 1 {
		if countes[0] == 1 {
			return true
		} else {
			return false
		}
	}
	if !(countOfCounts[1] == 1 && (countes[1]-1 == countes[0] || countes[1]-1 == 0)) && !(countOfCounts[0] == 1 && (countes[0]-1 == countes[1] || countes[0]-1 == 0)) {
		return false
	}

	return true
}
