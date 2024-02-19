package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dimensione, _ := strconv.Atoi(os.Args[1])
	x, _ := strconv.Atoi(os.Args[2])
	y, _ := strconv.Atoi(os.Args[3])

	//rigaOrizzontale(dimensione)
	//rigaVerticale(dimensione)
	for i := 0; i < dimensione; i++ {
		for j := 0; j <= dimensione; j++ {
			if i == 0 {
				fmt.Print(j)
			}
			/*if i == y-1 && j == x+1 {
				fmt.Print(strings.Repeat("p", x-1))
				fmt.Print("+")
			}*/
		}
		if i == 0 {
			fmt.Println()
		}
		if i == y-1 {
			fmt.Print(i+1, strings.Repeat(" ", x-1), "*\n")
		} else {
			fmt.Println(i + 1)
		}
	}
}

func rigaOrizzontale(dimensione int) {
	for i := 0; i <= dimensione; i++ {
		fmt.Print(i)
	}
}

func rigaVerticale(dimensione int) {
	for i := 1; i <= dimensione; i++ {
		fmt.Println(i)
	}
}
