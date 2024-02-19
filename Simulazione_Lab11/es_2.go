package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	s := os.Args[1]
	var parole []string
	m := make(map[string]int)
	var ss string
	for i := 0; i < len(s); i++ {
		for j := i; j <= len(s); j++ {
			ss = s[i:j]
			if len(ss) >= 3 && ss[0] == s[j-1] {
				parole = append(parole, ss)
			}
		}
	}

	for _, p := range parole {
		m[p]++
	}

	for k, v := range m {
		fmt.Println(k, "-> Occorrenze:", v)
	}

	// Estrai le chiavi dalla mappa
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	ordineAlfabetico(keys, m)
	ordineLunghezza(keys, m)

}

func ordineAlfabetico(keys []string, m map[string]int) {
	fmt.Println("stampo in ordine alfabetico")

	// Ordina l'array delle chiavi
	sort.Strings(keys)

	// Itera sull'array ordinato delle chiavi
	for _, key := range keys {
		fmt.Println(key, "-> Occorrenze:", m[key])
	}
}

func ordineLunghezza(keys []string, m map[string]int) {
	fmt.Println("stampo in ordine di lunghezza")

	sort.SliceStable(keys, func(x, y int) bool {
		return len(keys[x]) < len(keys[y])
	})

	fmt.Println("stampo in ordine alfabetico")

	for _, key := range keys {
		fmt.Println(key, "-> Occorrenze:", m[key])
	}
}
