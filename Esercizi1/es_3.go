package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	anno, mese, giorno int
}

type Operazione struct {
	data    Data
	tipo    string
	importo int
}

type Mappa map[Data]int

func main() {
	// Apri il file
	file, _ := os.Open(os.Args[1])

	// Chiude file correttamente alla fine della lettura di tutte le righe
	defer file.Close()

	//scanner := bufio.NewScanner(os.Stdin)
	scanner := bufio.NewScanner(file)

	var operazioni []Operazione

	var date []Data

	for scanner.Scan() {

		// divido gli elementi di una linea dopo ogni ";"
		parts := strings.Split(scanner.Text(), ";")

		data := strings.Split(parts[0], "_")
		tipo := parts[1]
		importo, _ := strconv.Atoi(parts[2])

		//divido data dopo ogni "_"
		anno, _ := strconv.Atoi(data[0])
		mese, _ := strconv.Atoi(data[1])
		giorno, _ := strconv.Atoi(data[2])

		data2 := Data{anno, mese, giorno}
		date = append(date, data2)
		//array contenente ogni linea
		operazioni = append(operazioni, Operazione{data2, tipo, importo})
	}

	m := trova(operazioni)

	for k, v := range m {
		fmt.Println(formatData(k), v)
	}
}

func NuovaMappa() Mappa {
	return make(Mappa)
}

func trova(operazioni []Operazione) Mappa {
	m := NuovaMappa()
	for _, v := range operazioni {
		if v.tipo == "V" {
			m[v.data] += v.importo
		} else {
			m[v.data] -= v.importo
		}
	}
	return m
}

func formatData(d Data) string {
	return strconv.Itoa(d.anno) + "/" + strconv.Itoa(d.mese) + "/" + strconv.Itoa(d.giorno) + " ->"
}
