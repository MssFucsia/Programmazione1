package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := os.Args[1:]
	var pari int
	for _, char := range s {
		numero, _ := strconv.Atoi(char)
		if numero%2 == 0 {
			pari++
		}
	}
	fmt.Print("Pari: ", pari, ",", " dispari: ", len(s)-pari, "\n")
}
