package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Punto struct {
	etichetta         string
	ascissa, ordinata float64
}

func main() {
	tragitto := NuovoTragitto()
	lunghezza := Lunghezza(tragitto)
	fmt.Printf("Lunghezza percorso: %.3f \n", lunghezza)
	fmt.Println("Punto oltre metà:", puntoOltreMetà(lunghezza, tragitto))
}

func NuovoTragitto() (tragitto []Punto) {
	//file, _ := os.Open(os.Args[1])

	//defer file.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ";")

		etichetta := parts[0]
		ascissa, _ := strconv.ParseFloat(parts[1], 64)
		ordinata, _ := strconv.ParseFloat(parts[2], 64)

		tragitto = append(tragitto, Punto{etichetta, ascissa, ordinata})
	}
	return
}

func Distanza(p1, p2 Punto) float64 {
	return math.Sqrt(math.Pow((p1.ascissa-p2.ascissa), 2) + math.Pow((p1.ordinata-p2.ordinata), 2))
}

func String(p Punto) string {
	return fmt.Sprintf("%s = (%.1f,%.1f)", p.etichetta, p.ascissa, p.ordinata)
}

func Lunghezza(tragitto []Punto) float64 {
	prec := tragitto[0]
	var l float64
	for i := 1; i < len(tragitto); i++ {
		l += Distanza(prec, tragitto[i])
		prec = tragitto[i]
	}
	return l
}

func puntoOltreMetà(l float64, tragitto []Punto) string {
	var d float64
	prec := tragitto[0]
	for i := 1; i < len(tragitto); i++ {
		d += Distanza(prec, tragitto[i])
		prec = tragitto[i]
		if d > l/2 {
			return String(tragitto[i])
		}
	}
	return "0"
}
