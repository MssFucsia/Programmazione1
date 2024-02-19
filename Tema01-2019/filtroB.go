package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := os.Args[1:]

	if len(s) == 0 {
		fmt.Println("Nessun valore in ingresso")
	} else {
		i, ok := valida(s)
		if ok {
			fmt.Println("Sequenza valida")
		} else {
			fmt.Println("Elemento in posizione", i, "non valido")
		}
	}
}

func valida(s []string) (int, bool) {
	prec, _ := strconv.Atoi(s[0])
	for i := 1; i < len(s); i++ {
		n, err := strconv.Atoi(s[i])
		if err == nil && ((prec >= 0 && n <= 0) || (prec <= 0 && n >= 0)) {
			prec = n
			continue
		} else {
			return i + 1, false
		}
	}
	return -1, true
}
