package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Studente struct {
	nome, cognome       string
	matricola, presenze int
}

func incrementaPresenze(p_studente *Studente) {
	p_studente.presenze++
}

func main() {
	var slice []Studente

	fileName := os.Args[1]
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		matricola, _ := strconv.Atoi(parts[2])
		presenze, _ := strconv.Atoi(parts[3])

		slice = append(slice, Studente{parts[0], parts[1], matricola, presenze})
	}
	char := os.Args[2]

	for _, s := range slice {
		if char == string(s.nome[0]) {
			incrementaPresenze(&s)
		}
		fmt.Println(s.nome, s.cognome, s.matricola, s.presenze)
	}

}
