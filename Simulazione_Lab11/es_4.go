package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := os.Args[1]
	d, _ := strconv.Atoi(os.Args[2])

	sottostringhe := riduci(n)

	// slice di strighe di lunghezza min len(n)-d
	for i := 0; i < d-1; i++ {
		for _, r := range sottostringhe {
			sottostringhe = append(sottostringhe, riduci(r)...)
		}

	}
	fmt.Println(min(sottostringhe))

}

// toglie un carattere
func riduci(n string) (sottostringhe []string) {
	for i := 0; i < len(n); i++ {
		sottostringhe = append(sottostringhe, n[:i]+n[i+1:])
	}
	return
}

func min(sottostringhe []string) (min string) {
	min = sottostringhe[0]
	for _, s := range sottostringhe {
		if s < min {
			min = s
		}
	}
	return min
}
