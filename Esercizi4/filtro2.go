/*Scrivere un programma ​filtro.go che, data una sequenza di numeri interi positivi inseriti
da linea di comando, stampi a schermo le potenze di 3, nell'ordine in cui appaiono.

Ad esempio:
$ go run filtro.go 1 9 2 27 3 25 6
1 9 27 3*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := os.Args[1:]

	for _, r := range s {
		n, _ := strconv.Atoi(r)
		if Potenza(n) {
			fmt.Print(n, " ")
		}
	}
}

func Potenza(n int) bool {
	l := n
	for i := 0; i < l; i++ {
		if n%3 == 0 {
			n /= 3
		}
		if n == 1 {
			return true
		}
	}
	return false
}
