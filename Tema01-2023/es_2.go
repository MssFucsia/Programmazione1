package main

import (
	"fmt"
	"os"
	"sort"
	"unicode"
)

func main() {
	s := os.Args[1:]

	sottosequenze := GeneraSottosequenze(s)

	var keys []string
	for k := range sottosequenze {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	fmt.Println(sottosequenze)

	for _, k := range keys {
		fmt.Println(k, sottosequenze[k])
	}
}

func GeneraSottosequenze(sequenza []string) (sottosequenze map[string]int) {
	var stringa string
	for _, char := range sequenza {
		stringa += string(unicode.ToLower([]rune(char)[0]))
	}
	sottosequenze = make(map[string]int)
	for i := 0; i < len(stringa); i++ {
		for j := i; j <= len(stringa); j++ {
			ss := stringa[i:j]
			if len(ss) > 1 && ss[0] == ss[len(ss)-1] {
				sottosequenze[ss]++
			}

		}
	}
	return
}
