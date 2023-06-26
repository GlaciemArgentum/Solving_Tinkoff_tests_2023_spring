package main

import (
	"fmt"
	"os"
	"sort"
)

type Score struct {
	l int
	r int
}

var null = Score{0, 0}

type Scores []Score

func (a Scores) Len() int           { return len(a) }
func (a Scores) Less(i, j int) bool { return a[i].l < a[j].l || (a[i].l == a[j].l && a[i].r < a[j].r) }
func (a Scores) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

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
	fmt.Println(answer)
}

func getMaxValue(data []Score) int {
	max := data[0].r
	for _, v := range data[1:] {
		if right := v.r; right > max {
			max = right
		}
	}
	return max
}

func getCountOfValue(v int, data *[]Score) int {
	counter := 0
	for i := len(*data) - 1; i >= 0; i-- {
		if (*data)[i].l <= v && v <= (*data)[i].r {
			counter++
			(*data)[i] = Score{0, 0}
		}
	}
	return counter
}

func getMoreThenValue(n, s, i, sum, counterRight int, data *[]Score) (bool, int) {
	for v := len(*data) - 1; counterRight < (n+1)/2 && v >= 0; v-- {
		if (*data)[v] == null {
			continue
		}
		if (*data)[v].r < i {
			return false, 0
		}
		sum += (*data)[v].l
		if sum > s {
			return false, 0
		}
		counterRight++
		(*data)[v] = Score{0, 0}
	}

	if counterRight < (n+1)/2 {
		return false, 0
	}

	return true, sum
}

func getLessThenValue(s, sum int, data *[]Score) bool {
	for v := len(*data) - 1; v >= 0; v-- {
		if (*data)[v] == null {
			continue
		}
		sum += (*data)[v].l
		if sum > s {
			return false
		}
	}
	return true
}

func bestMiddle(n, s int, scores []Score) int {
	sort.Sort(Scores(scores))
	maxPossible := (2*s + 1 - n) / (n + 1)
	if maxValue := getMaxValue(scores); maxValue < maxPossible {
		maxPossible = maxValue
	}

	var answer int
	for i := maxPossible; i >= 0; i-- {
		data := make([]Score, n, n)
		copy(data, scores)

		counterRight := getCountOfValue(i, &data)
		if counterRight == 0 {
			continue
		}

		flag, sum := getMoreThenValue(n, s, i, counterRight*i, counterRight, &data)
		if !flag {
			continue
		}
		if getLessThenValue(s, sum, &data) {
			answer = i
			break
		}

	}

	return answer
}

func main() {
	myWrite(bestMiddle(myRead("std")))
}
