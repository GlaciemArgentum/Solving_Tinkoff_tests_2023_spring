package main

import (
	"fmt"
	"math"
)

func myRead() (int, int, int) {
	var n, m, k int
	fmt.Scanf("%d %d %d", &n, &m, &k)
	return n, m, k
}

func timeReview(n, m, k int) int {
	var num = float64(n*k) / float64(m)
	return int(math.Ceil(num))
}

func myWrite(res int) {
	fmt.Println(res)
}

func main() {
	myWrite(timeReview(myRead()))
}
