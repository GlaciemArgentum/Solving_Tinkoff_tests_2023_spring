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
	rawData := buf.Text()
	sliceData := strings.Split(rawData, " ")
	data := make([]int, len(sliceData), len(sliceData))
	for i, v := range sliceData {
		data[i], _ = strconv.Atoi(v)
	}

	return data
}

func myWrite(answer bool) {
	if answer {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

func isQueueInFormat(queue []int) bool {
	if (queue[0] <= queue[1] && queue[1] <= queue[2] && queue[2] <= queue[3]) || (queue[0] >= queue[1] && queue[1] >= queue[2] && queue[2] >= queue[3]) {
		return true
	} else {
		return false
	}
}

func main() {
	myWrite(isQueueInFormat(myRead("std")))
}
