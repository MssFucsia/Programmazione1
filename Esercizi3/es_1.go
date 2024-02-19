package main

import (
	"fmt"
	"unicode"
)

func main() {
	var parola string
	for {
		fmt.Println("inserisci una parola")
		fmt.Scan(&parola)
		if parola == "0" {
			break
		}
		messaggio, valore, ok := parolaCorretta(parola)
		if ok {
			fmt.Println("corretto", messaggio, valore)
		} else if valore == -1 {
			fmt.Println("errore")
		} else {
			fmt.Println("sbagliato")
		}
	}
}

func parolaCorretta(parola string) (lettera string, valore int, ok bool) {
	vocali := make(map[rune]int)

	for _, lettera := range parola {
		unicode.ToLower(lettera)
		if lettera < 'a' || lettera > 'z' {
			return "errore", -1, false
		}

		if lettera == 'a' || lettera == 'e' || lettera == 'i' || lettera == 'o' || lettera == 'u' {
			vocali[lettera]++
		}
	}

	for k, v := range vocali {
		lettera = string(k)
		valore = v
	}

	if len(vocali) == 1 {
		return lettera, valore, true
	}

	return "sbagliato", -2, false
}
