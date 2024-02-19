package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n string
	var x int
	var vittorie, sconfitte, iterazioni float64
	rand.Seed(time.Now().UnixNano())

	for {
		numeroCasuale := rand.Intn(2)
		fmt.Print(numeroCasuale)

		fmt.Println("Dove sto guardando?")
		fmt.Scan(&n)
		if n == "d" {
			x = 0
		} else if n == "s" {
			x = 1
		} else {
			break
		}
		// Inizializza il generatore di numeri casuali

		// Genera un numero casuale tra 0 e 1
		if x == numeroCasuale {
			fmt.Println("Giusto, hai vinto!")
			vittorie++
		} else {
			fmt.Println("Sbagliato, hai perso!")
			sconfitte++
		}
		iterazioni++
	}

	fmt.Printf("Vinto: %2.f%% Perso: %2.f%% \n", (vittorie/iterazioni)*100, (sconfitte/iterazioni)*100)

}
