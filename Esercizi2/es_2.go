package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	N, _ := strconv.Atoi(os.Args[1])
	k, _ := strconv.Atoi(os.Args[2])

	pari := FiltraNumeri(GeneraNumeri(N, k))
	for _, numero := range pari {
		fmt.Println(numero)
	}
}

func GeneraNumeri(N, k int) []int {
	s := strconv.Itoa(N)
	var numeri []int
	for i := 0; i < len(s); i++ {
		for j := i; j <= len(s); j++ {
			if len(s[i:j]) == k {
				n, _ := strconv.Atoi(s[i:j])
				numeri = append(numeri, n)
			}
		}
	}
	return numeri
}

func FiltraNumeri(sl []int) []int {
	var pari []int
	for _, numero := range sl {
		if numero%2 == 0 {
			pari = append(pari, numero)
		}
	}
	return pari
}
