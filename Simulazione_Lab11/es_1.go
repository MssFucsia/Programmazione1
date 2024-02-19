package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := os.Args[1:]

	for i, parola := range s {
		fmt.Print(TrasformaParola(parola, i), " ")
	}

	fmt.Println()
}

func TrasformaParola(parola string, posizione int) (parolaTrasformata string) {
	if posizione%2 == 0 {
		for i := range parola {
			if i%2 == 0 {
				parolaTrasformata += strings.ToUpper(string(parola[i]))
			} else {
				parolaTrasformata += strings.ToLower(string(parola[i]))

			}
		}
	} else {
		for i := range parola {
			if i%2 == 0 {
				parolaTrasformata += strings.ToLower(string(parola[i]))
			} else {
				parolaTrasformata += strings.ToUpper(string(parola[i]))

			}
		}
	}
	return parolaTrasformata
}
