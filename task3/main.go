package main

import "fmt"

func isMapFull(data map[byte]int) bool {
	for _, v := range []byte("abcd") {
		if i, _ := data[v]; i == -1 {
			return false
		}
	}
	return true
}

func myRead() (int, []byte) {
	var (
		n   int
		str string
	)
	fmt.Scanf("%d\n%s", &n, &str)
	return n, []byte(str)
}

func myWrite(l int) {
	fmt.Println(l)
}

func minLenString(n int, data []byte) int {
	if n < 4 {
		return -1
	}

	alphabet := make(map[byte]int, 4)
	for _, v := range data {
		alphabet[v]++
	}
	if !isMapFull(alphabet) {
		return -1
	}

	for lenS := 4; lenS <= n; lenS++ {
		strPiece := data[:lenS]
		alphabet = map[byte]int{byte('a'): -1, byte('b'): -1, byte('c'): -1, byte('d'): -1}
		for _, v := range strPiece {
			alphabet[v]++
		}
		if isMapFull(alphabet) {
			return lenS
		}

		for i := lenS; i < n; i++ {
			alphabet[data[i-lenS]]--
			alphabet[data[i]]++
			if isMapFull(alphabet) {
				return lenS
			}
		}
	}
	return -1
}

func main() {
	myWrite(minLenString(myRead()))
}
