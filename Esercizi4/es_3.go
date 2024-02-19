package main

import (
	"fmt"
)

func main() {
	var parola string

	for {
		fmt.Print("Inserisci una parola: ")
		fmt.Scan(&parola)

		if parola == "0" {
			break
		}

		if len(Doppie(parola)) == 0 {
			fmt.Println("Non ci sono doppie")
		} else {
			fmt.Println(Doppie(parola))
		}
	}
}

func Doppie(parola string) (doppie []string) {
	p := parola[0]

	for i := 1; i < len(parola); i++ {
		if p == parola[i] {
			doppie = append(doppie, string(parola[i]))
		}
		p = parola[i]
	}
	return
}
