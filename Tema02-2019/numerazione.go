package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	s := os.Args[1]

	if reverse, ok := valido(s); ok {
		var tot float64
		for i := 0; i < len(reverse); i++ {
			switch reverse[i] {
			case 'z':
				tot += 0
			case 'u':
				tot += math.Pow(3, float64(i))
			case 'd':
				tot += 2 * math.Pow(3, float64(i))
			}
		}
		fmt.Println(tot)
	} else {
		fmt.Println("input non valido")
	}

}

func valido(s string) (reverse string, b bool) {
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != "u" && string(s[i]) != "z" && string(s[i]) != "d" {
			return "", false
		}
		reverse += string(s[i])
	}
	return reverse, true
}
