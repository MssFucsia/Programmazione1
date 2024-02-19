package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var t string
	for scanner.Scan() {
		t += scanner.Text()
	}
	nospazi := strings.ReplaceAll(t, " ", "")
	vocali := make(map[rune]int)
	ContaVocali(nospazi, vocali)
}

func ContaVocali(s string, vocali map[rune]int) {
	for _, v := range s {
		if Vocale(v, s) {
			vocali[v]++
		}
	}

	for k, v := range vocali {
		fmt.Println(string(k), ":", v)
	}

	// ORDINE ALFABETICO DELLE CHIAVI
	fmt.Println()
	fmt.Println("Stampa in ordine alfabetico")
	var keys []rune

	for k := range vocali {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, v := range keys {
		fmt.Println(string(v), ":", vocali[v])
	}
}

func Vocale(r rune, s string) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}

}
