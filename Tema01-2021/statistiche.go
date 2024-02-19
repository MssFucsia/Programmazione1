package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var serie []float64
	for scanner.Scan() {
		n, _ := strconv.ParseFloat(scanner.Text(), 64)
		serie = append(serie, n)
	}
	fmt.Println("moda:", moda(serie))
	fmt.Println("mediana:", mediana(serie))
}

func moda(serie []float64) float64 {
	m := make(map[float64]int)
	for _, n := range serie {
		m[n]++
	}

	keys := make([]float64, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	var minimi []float64
	min := m[keys[0]]
	for _, v := range keys {
		if m[v] == min {
			minimi = append(minimi, v)
		}
	}
	sort.SliceStable(minimi, func(i, j int) bool {
		return minimi[i] < minimi[j]
	})

	return minimi[0]
}

func mediana(serie []float64) float64 {
	sort.SliceStable(serie, func(i, j int) bool {
		return serie[i] < serie[j]
	})

	if len(serie)%2 != 0 {
		return serie[(len(serie)-1)/2]
	}

	return (serie[len(serie)/2] + serie[(len(serie)/2)-1]) / 2
}
