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

func countNormSegment(n int, data []int) int {
	var (
		answer, start int
		flag          bool
	)

	if n == 1 && data[0] == 0 {
		return 1
	}

	for left := 0; left < n-1; left++ {
		sum := data[left]
		for right := left + 1; right < n; right++ {
			sum += data[right]
			if sum == 0 {
				answer += (left + 1) * (n - right)
				start = left + 1
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}

	if answer == 0 {
		return 0
	}

	for left := start; left < n-1; left++ {
		sum := data[left]
		for right := left + 1; right < n; right++ {
			sum += data[right]
			if sum == 0 {
				answer += n - right
				break
			}
		}
	}

	return answer
}

func main() {
	myWrite(countNormSegment(myRead("std")))
}
