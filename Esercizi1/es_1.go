package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numero := os.Args[1]

	for j := 1; j <= 3; j++ {
		for i := 0; i <= len(numero)-j; i++ {
			n, _ := strconv.Atoi(numero[:i] + numero[i+j:])
			if èPrimo(n) {
				fmt.Println(n)
			}
		}
	}

}

func èPrimo(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
