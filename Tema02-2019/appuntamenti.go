package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Appuntamento struct {
	giorno, oraInizio, oraFine int
}

func NewAppuntamento(gg, h1, h2 int) (Appuntamento, bool) {
	if gg >= 1 && gg <= 366 && h1 <= h2 && h1 >= 0 && h1 <= 24 && h2 >= 0 && h2 <= 24 {
		return Appuntamento{gg, h1, h2}, true
	}
	return Appuntamento{0, 0, 0}, false
}

func CheckSovrapposizioneAppuntamenti(app1, app2 Appuntamento) bool {
	return (app1.giorno == app2.giorno) && app2.oraInizio <= app1.oraFine
}

func AddAppuntamento(app Appuntamento, agenda *[]Appuntamento) bool {
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var appuntamenti []Appuntamento
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		gg, _ := strconv.Atoi(parts[0])
		h1, _ := strconv.Atoi(parts[1])
		h2, _ := strconv.Atoi(parts[2])

		appuntamento, ok := NewAppuntamento(gg, h1, h2)
		if ok {
			appuntamenti = append(appuntamenti, appuntamento)
		}
	}

	var agenda []Appuntamento
	agenda = append(agenda, appuntamenti[0])
	prec := appuntamenti[0]
	for i := 1; i < len(appuntamenti); i++ {
		fmt.Println(prec)
		if CheckSovrapposizioneAppuntamenti(prec, appuntamenti[i]) {
			prec = appuntamenti[i]
			continue
		}
		agenda = append(agenda, appuntamenti[i])
		prec = appuntamenti[i]
	}
	fmt.Println(agenda)

}
