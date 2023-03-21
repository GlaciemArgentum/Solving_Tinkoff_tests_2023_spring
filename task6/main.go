package main

import (
	"fmt"
	"os"
)

type Score struct {
	l int
	r int
}

func myRead(name string) (int, int, []Score) {
	var file *os.File
	if name == "" {
		file, _ = os.Open("input.txt")
	} else if name == "std" {
		file = os.Stdin
	} else {
		file, _ = os.Open(name + ".txt")
	}
	defer file.Close()

	var n, s int
	fmt.Fscanf(file, "%d %d\n", &n, &s)

	scores := make([]Score, 0, n)
	var score Score
	for i := 0; i < n; i++ {
		fmt.Fscanf(file, "%d %d\n", &score.l, &score.r)
		scores = append(scores, score)
	}

	return n, s, scores
}

func myWrite(answer int) {
	//file, _ := os.Create("output.txt")
	//_, _ = file.WriteString(fmt.Sprintf("%d\n", answer))
	fmt.Println(answer)
}

func bestMiddle(n, s int, scores []Score) int {

}

func main() {
	myWrite(bestMiddle(myRead("std")))
}
