package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	s := os.Args[1]
	fmt.Print(contaCifre(s))

}

func contaCifre(s string) [10]int {
	var array [10]int
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit([]rune(string(s[i]))[0]) {
			switch s[i] {
			case '0':
				array[0]++
			case '1':
				array[1]++
			case '2':
				array[2]++
			case '3':
				array[3]++
			case '4':
				array[4]++
			case '5':
				array[5]++
			case '6':
				array[6]++
			case '7':
				array[7]++
			case '8':
				array[8]++
			case '9':
				array[9]++
			}
		}
	}
	return array
}
