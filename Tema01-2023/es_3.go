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
	etichetta string
	x, y      float64
}

type Triangolo struct {
	punto1, punto2, punto3 Punto
}

func LeggiPunti() (punti []Punto) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ";")

		etichetta := parts[0]
		x, _ := strconv.ParseFloat(parts[1], 64)
		y, _ := strconv.ParseFloat(parts[2], 64)

		punti = append(punti, Punto{etichetta, x, y})
	}
	return
}

func Area(t Triangolo) float64 {
	l1 := Distanza(t.punto1, t.punto2)
	l2 := Distanza(t.punto1, t.punto3)
	l3 := Distanza(t.punto3, t.punto2)
	p := (l1 + l2 + l3) / 2
	return math.Sqrt((p * (p - l1) * (p - l2) * (p - l3)))
}

func GeneraTriangoli(punti []Punto) (triangoli []Triangolo) {
	for i := 0; i < len(punti); i++ {
		for j := i; j < len(punti); j++ {
			for k := j; k < len(punti); k++ {
				if i != k && k != j && i != j {
					triangoli = append(triangoli, Triangolo{punti[i], punti[j], punti[k]})
				}
			}
		}
	}
	return
}

func Distanza(p1, p2 Punto) float64 {
	return math.Sqrt(math.Pow((p1.x-p2.x), 2) + math.Pow((p1.y-p2.y), 2))
}

func StringPunto(p Punto) string {
	return fmt.Sprintf("%s = (%.1f,%.1f) ", p.etichetta, p.x, p.y)
}

func StringTriangolo(t Triangolo) string {
	return fmt.Sprint("Triangolo rettangolo con vertici ", StringPunto(t.punto1), StringPunto(t.punto2), StringPunto(t.punto3), "e area ")
}

func Rettangolo(t Triangolo) bool {
	if t.punto1.y == t.punto2.y || t.punto1.y == t.punto3.y || t.punto2.y == t.punto3.y {
		return true
	}
	return false
}

func Quadrante(t Triangolo) bool {
	if (t.punto1.x < 0 && t.punto2.x < 0 && t.punto3.x < 0) && ((t.punto1.y > 0 && t.punto2.y > 0 && t.punto3.y > 0) || (t.punto1.y < 0 && t.punto2.y < 0 && t.punto3.y < 0)) {
		return true
	} else if (t.punto1.x > 0 && t.punto2.x > 0 && t.punto3.x > 0) && ((t.punto1.y > 0 && t.punto2.y > 0 && t.punto3.y > 0) || (t.punto1.y < 0 && t.punto2.y < 0 && t.punto3.y < 0)) {
		return true
	}
	return false
}

func FiltraTriangoli(triangoli []Triangolo) (triangoli_filtrati []Triangolo) {
	for _, t := range triangoli {
		if Rettangolo(t) && Quadrante(t) {
			triangoli_filtrati = append(triangoli_filtrati, t)
		}
	}
	return
}

func main() {
	triangoli := FiltraTriangoli(GeneraTriangoli(LeggiPunti()))

	if len(triangoli) != 0 {
		area_max := Area(triangoli[0])
		triangolo_max := triangoli[0]
		for i := 1; i < len(triangoli); i++ {
			if Area(triangoli[i]) > area_max {
				area_max = Area(triangoli[i])
				triangolo_max = triangoli[i]
			}
		}

		fmt.Print(StringTriangolo(triangolo_max))
		fmt.Printf("%.1f \n", area_max)
	}

}
