package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	s := os.Args[1:]
	var slice []rune
	for _, char := range s {
		slice = append(slice, []rune(char)[0])
	}
	for i, char := range slice {
		if unicode.IsDigit(char) {
			n, _ := strconv.Atoi(string(char))
			fmt.Print(strings.Repeat(string(slice[i-1]), n))
		}
	}
}
