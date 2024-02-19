package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := os.Args[1:]
	Sottosequenze(s)
}

func Sottosequenze(s []string) {
	//var numeri []int

	for i := 0; i < len(s); i++ {
		for j := i; j <= len(s); j++ {
			ss := s[i:j]
			if len(ss) != 0 && Somma(ss) {
				fmt.Println(ss)
			}
		}
	}
}

func Somma(ss []string) bool {
	var somma int
	for _, char := range ss {
		n, _ := strconv.Atoi(char)
		somma += n
	}
	if somma == 0 {
		return true
	}
	return false
}
