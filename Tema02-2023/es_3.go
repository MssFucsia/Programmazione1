package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Prodotto struct {
	nome   string
	codice string
}

type Magazzino map[Prodotto]int

func NuovoMagazzino() Magazzino {
	return make(map[Prodotto]int)
}

func StringProdotto(p Prodotto) string {

	return fmt.Sprintf("Codice: %s, Prodotto: %s", p.codice, p.nome)
}

func CreaProdotto(nome string, codice string) Prodotto {
	return Prodotto{nome, codice}
}

func NumeroProdotti(m Magazzino) int {
	return len(m)
}

func Quantità(m Magazzino, p Prodotto) int {
	return m[p]
}

func ModificaProdotto(m Magazzino, p Prodotto, variazione int) (Magazzino, bool) {
	_, found := m[p]

	if found {
		m[p] += variazione
		if m[p] > 0 {
			return m, true
		} else if m[p] == 0 {
			delete(m, p)
			return m, true
		} else {
			m[p] -= variazione
			return m, false
		}
	} else if found == false && variazione > 0 {
		m[p] += variazione
	}
	return m, true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	m := NuovoMagazzino()
	riga := 1
	var prodotti []Prodotto

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ";")
		quantità, _ := strconv.Atoi(parts[2])
		//m[CreaProdotto(parts[1], parts[0])] += q
		nome := parts[1]
		codice := parts[0]

		p := CreaProdotto(nome, codice)
		prodotti = append(prodotti, p)

		if m, ok := ModificaProdotto(m, p, quantità); ok == false {
			fmt.Print("Errore alla riga ", riga, ", Codice: ", codice, ", Prodotto: ", nome, ", Quantità: ", m[p], ", Richiesta: ", quantità, "\n")
			return
		}
		riga++
	}
	fmt.Println("Il magazzino contiene", NumeroProdotti(m), "prodotti:")

	for p := range m {
		prodotti = append(prodotti, p)
	}

	sort.Slice(prodotti, func(i, j int) bool {
		return prodotti[i].codice < prodotti[j].codice
	})

	for _, p := range prodotti {
		fmt.Println(StringProdotto(p), "Quantità", m[p])
	}

}
