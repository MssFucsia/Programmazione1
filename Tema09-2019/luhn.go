package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := os.Args[1:]
	if len(s) == 0 {
		fmt.Println("Nessun input")
	} else {
		for _, r := range s {
			if luhn(r) {
				fmt.Println(r, "valido")
			} else {
				fmt.Println(r, "non valido")
			}

		}
	}
}

func luhn(cifra string) bool {
	var reverse string
	for i := len(cifra) - 1; i >= 0; i-- {
		reverse += string(cifra[i])
	}
	somma, _ := strconv.Atoi(string(reverse[0]))
	for i := 1; i < len(reverse); i++ {
		n, err := strconv.Atoi(string(reverse[i]))
		if err != nil {
			return false
		}
		if i%2 != 0 {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}

		somma += n

	}

	if somma%10 == 0 {
		return true
	}
	return false
}
