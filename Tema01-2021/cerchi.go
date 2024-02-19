package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cerchio struct {
	nome         string
	x, y, raggio float64
}

func newCircle(descr string) Cerchio {
	parts := strings.Fields(descr)
	x, _ := strconv.ParseFloat(parts[1], 64)
	y, _ := strconv.ParseFloat(parts[2], 64)
	raggio, _ := strconv.ParseFloat(parts[3], 64)
	return Cerchio{parts[0], x, y, raggio}

}

func distanzaPunti(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

func tocca(cerc1, cerc2 Cerchio) bool {
	a := math.Abs(distanzaPunti(cerc1.x, cerc1.y, cerc2.x, cerc2.y)-math.Abs((cerc1.raggio-cerc2.raggio))) <= 1e-6
	b := math.Abs(distanzaPunti(cerc1.x, cerc1.y, cerc2.x, cerc2.y)-(cerc1.raggio+cerc2.raggio)) <= 1e-6
	if a || b {
		return true

	}

	return false
}

func maggiore(cerc1, cerc2 Cerchio) bool {
	return cerc1.raggio > cerc2.raggio
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	prec := Cerchio{"", 0, 0, 0}

	for scanner.Scan() {
		t := scanner.Text()
		cerchio := newCircle(t)

		fmt.Print(cerchio)

		if tocca(prec, cerchio) {
			fmt.Print(", tangente, ")
		} else {
			fmt.Print(", non tangente, ")
		}

		if maggiore(cerchio, prec) {
			fmt.Println("maggiore")
		} else {
			fmt.Println("minore o uguale")
		}

		prec = cerchio
	}

}
