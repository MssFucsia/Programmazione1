package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {

	var s string
	fmt.Scan(&s)

	s_rune := []rune(s)

	var sequenza []int

	for _, char := range s_rune {
		if unicode.IsDigit(char) {
			n, _ := strconv.Atoi(string(char))
			sequenza = append(sequenza, n)
		}
	}

	if sequenzaValida(sequenza) {
		fmt.Println("Sequenza nascosta ordinata.")
	} else {
		fmt.Println("Sequenza nascosta non ordinata")
	}
}

func sequenzaValida(sequenza []int) bool {
	precedente := sequenza[0]
	for i := 1; i < len(sequenza); i++ {
		if sequenza[i] < precedente {
			precedente = sequenza[i]
			continue
		}
		return false
	}
	return true
}
