package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func myRead(name string) []int {
	var buf *bufio.Scanner
	if name == "" {
		file, _ := os.Open("input.txt")
		buf = bufio.NewScanner(file)
	} else if name == "std" {
		buf = bufio.NewScanner(os.Stdin)
	} else {
		file, _ := os.Open(name + ".txt")
		buf = bufio.NewScanner(file)
	}

	buf.Scan()
	rawData := buf.Text()
	sliceData := strings.Split(rawData, " ")
	data := make([]int, len(sliceData), len(sliceData))
	for i, v := range sliceData {
		data[i], _ = strconv.Atoi(v)
	}

	return data
}

func myWrite(answer bool) {
	//file, _ := os.Create("output.txt")
	if answer {
		//_, _ = file.WriteString("YES")
		fmt.Println("YES")
	} else {
		//_, _ = file.WriteString("NO")
		fmt.Println("NO")
	}

}

func isQueueInFormat(queue []int) bool {
	var isIncrease bool
	pre := queue[0]
	newStart := 0
	for i, v := range queue[1:] {
		if pre != v {
			if v > pre {
				isIncrease = true
			} else {
				isIncrease = false
			}
			newStart = i
			break
		}
	}
	for i, v := range queue[newStart+1:] {
		if v != queue[i-1] && ((v > queue[i-1]) != isIncrease) {
			return false
		}
	}
	return true
}

func main() {
	myWrite(isQueueInFormat(myRead("std")))
}
