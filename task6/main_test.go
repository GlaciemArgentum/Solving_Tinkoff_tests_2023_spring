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

func Generate(k, n, s int) (int, int, []Score) {
	mathRand.Seed(myInit())

	slice := make([]Score, 0, n)
	for i := 0; i < n; i++ {
		l := mathRand.Intn(k) + 1
		r := mathRand.Intn(k) + 1 + l
		slice = append(slice, Score{l, r})
	}
	return n, s, slice
}

func TestFunc(t *testing.T) {
	n := 100
	for i := 0; i < n; i++ {
		k, s, ss := Generate(10, 15, 30)
		fmt.Println(k, s, ss)
		bestMiddle(k, s, ss)
	}
}
