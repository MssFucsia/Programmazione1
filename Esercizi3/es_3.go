package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[1]
	m := make(map[rune]int)
	for _, char := range s {
		m[char]++
	}

	delete(m, ' ')

	if eterogramma(m) {
		fmt.Println("È eterogramma")
	} else {
		fmt.Println("Non è eterogramma")
	}

}

func eterogramma(m map[rune]int) bool {
	for _, v := range m {
		if v != 1 {
			return false
		}
	}
	return true

}
