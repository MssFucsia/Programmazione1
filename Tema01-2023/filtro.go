package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[1]
	l := len(s)
	soloAscii(s, l)
	unicode(s)
}

func soloAscii(s string, l int) {
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i == j || j == l-i-1 {
				fmt.Print(string(s[i]))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()

	}
}

func unicode(s string) {
	r := []rune(s)
	l := len(r)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i == j || j == l-i-1 {
				fmt.Print(string(r[i]))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()

	}
}
