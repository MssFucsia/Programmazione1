package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	if s == "L" {
		elle()
	} else if s == "T" {
		ti()
	} else if s == "Z" {
		zeta()
	} else {
		fmt.Println("input non valido")
	}
}

func elle() {
	for i := 0; i < 4; i++ {
		fmt.Println("*")
	}

	fmt.Print(strings.Repeat("*", 5))
	fmt.Println()
}

func ti() {
	fmt.Print(strings.Repeat("*", 5))
	fmt.Println()

	for i := 0; i < 4; i++ {
		fmt.Println(strings.Repeat(" ", 1), "*")
	}
}

func zeta() {
	for i := 0; i < 5; i++ {
		if i == 0 || i == 4 {
			fmt.Print(strings.Repeat("*", 5))
			fmt.Println()
		} else {
			fmt.Println(strings.Repeat(" ", 3-i), "*")

		}
	}

}
