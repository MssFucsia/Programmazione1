package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := os.Args[1:]
	posizione, ok := sequenzaValida(s)
	if ok {
		fmt.Println("Sequenza valida")
	} else {
		fmt.Println("Valore in posizione", posizione, "non valido")
	}
}

func sequenzaValida(s []string) (int, bool) {
	var somma int
	for i, char := range s {
		n, err := strconv.Atoi(char)
		if err == nil {
			if i%2 == 0 {
				if n > somma {
					somma += n
					continue
				} else {
					return i, false
				}
			}
			somma += n
		}
	}
	return -1, true
}
