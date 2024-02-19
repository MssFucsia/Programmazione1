package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var line string
	riga := 1
	dizionario := make(map[string][]int)

	for {
		fmt.Println("+ (Legge una riga e memorizza le parole)")
		fmt.Println("? (Legge una parola e indica le righe che la contengono)")
		fmt.Println("p (Stampa le parole)")

		if scanner.Scan() {
			line = scanner.Text()
		} else {
			break
		}

		parts := strings.Fields(line)
		simbolo := parts[0]
		parole := parts[1:]

		if simbolo == "+" {
			creaDizionario(parole, riga, dizionario)
			riga++
		} else if simbolo == "?" {
			cercaParola(dizionario, parole[0])
		} else {
			fmt.Println(dizionario)
		}

	}

}

func creaDizionario(parole []string, riga int, dizionario map[string][]int) {
	for _, word := range parole {
		dizionario[word] = append(dizionario[word], riga)
	}
}

func cercaParola(dizionario map[string][]int, s string) {
	fmt.Println("parola:", s)
	fmt.Println("righe", dizionario[s])

}
