package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func myRead(name string) (int, []int) {
	var buf *bufio.Scanner
	if name == "" {
		file, _ := os.Open("input.txt")
		defer file.Close()
		buf = bufio.NewScanner(file)
	} else if name == "std" {
		buf = bufio.NewScanner(os.Stdin)
	} else {
		file, _ := os.Open(name + ".txt")
		defer file.Close()
		buf = bufio.NewScanner(file)
	}

	buf.Scan()
	n, _ := strconv.Atoi(buf.Text())
	buf.Scan()
	rawData := buf.Text()
	sliceData := strings.Split(rawData, " ")
	data := make([]int, len(sliceData), len(sliceData))
	for i, v := range sliceData {
		data[i], _ = strconv.Atoi(v)
	}

	return n, data
}

func myWrite(answer int) {
	fmt.Println(answer)
}

func isMapFull(myMap map[int]int) bool {
	counterMap := make(map[int]int, 3)
	for _, v := range myMap {
		if v == 0 {
			continue
		}
		counterMap[v]++

		if len(counterMap) > 2 {
			return false
		}
	}

	var (
		countes       = make([]int, 0, 2)
		countOfCounts = make([]int, 0, 2)
	)
	for key, v := range counterMap {
		countes = append(countes, key)
		countOfCounts = append(countOfCounts, v)
	}
	if len(counterMap) == 1 {
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

func findPrefix(n int, data []int) int {
	if n <= 3 {
		return n - 1
	}

	library := make(map[int]int, n)
	maxLen := 3
	for i := 0; i < 3; i++ {
		library[data[i]]++
	}

	for i := 3; i < n; i++ {
		library[data[i]]++
		if isMapFull(library) {
			maxLen = i
		}
	}
	return maxLen + 1
}

func main() {
	myWrite(findPrefix(myRead("std")))
}
