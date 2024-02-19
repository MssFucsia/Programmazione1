package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if isIsogram(scanner.Text()) {
			fmt.Print(scanner.Text(), ": SI\n")
		} else {
			fmt.Print(scanner.Text(), ": NO\n")
		}
	}
}

func isIsogram(s string) bool {
	lettere := make(map[rune]int)
	for _, char := range s {
		if unicode.IsSpace(char) == false && char != '-' {
			lettere[char]++
		}
	}

	for k := range lettere {
		if lettere[k] == 1 {
			continue
		}
		return false
	}
	return true
}
