package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	/*file, _ := os.Open(os.Args[1])

	defer file.Close()*/

	scanner := bufio.NewScanner(os.Stdin)
	var lista []string
	for scanner.Scan() {
		lista = strings.Fields(scanner.Text())
	}

	fmt.Println("tot", cancella(lista))

}

func cancella(lista []string) []string {
	var nuovo []string
	for i := 0; i < len(lista); i++ {
		if unicode.IsDigit([]rune(lista[i])[0]) == false {
			nuovo = append(nuovo, lista[i])
		} else {
			n, _ := strconv.Atoi(lista[i])
			i += n
		}
	}
	return nuovo
}
