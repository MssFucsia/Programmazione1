package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	s := os.Args[1:]

	var p_errata int

	for _, parola := range s {
		if parolaEsatta(parola) {
			fmt.Print(parola, " ")
		} else {
			p_errata++
		}
	}
	fmt.Println(p_errata)
}

func parolaEsatta(p string) bool {
	for _, char := range p {
		if unicode.IsDigit(char) {
			return false
		}
	}
	return true
}
