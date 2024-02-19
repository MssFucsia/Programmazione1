package main

import (
	"fmt"

	"strings"
)

func main() {
	var s string
	var uno int
	for {
		fmt.Print("Inserisci un numero: ")
		fmt.Scan(&s)

		if s == "0" {
			break
		}

		for _, r := range s {

			n := strings.Count(s, string(r))

			if n == 1 {
				uno++
			}
		}

		if uno == len(s) {
			fmt.Println("Il numero non ha cifre ripetute")
		} else {
			fmt.Println("Il numero ha cifre ripetute")
		}
	}
}
