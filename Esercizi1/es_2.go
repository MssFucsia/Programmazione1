package main

import (
	"fmt"
	"strings"
)

func main() {

	var n, L, H int

	fmt.Scan(&n, &L, &H)

	var c, p, q, t, primo int
	for i := 0; i < 2*n+1; i++ {
		if i%2 == 0 {
			if p%2 == 0 {
				for j := 0; j < L; j++ {
					fmt.Print(c)
					c++
					if c == 10 {
						c = 0
					}
				}
			} else {
				t = c + L - 1
				if t >= 10 {
					ripetizioni := t - 9
					var contatore int
					for s := 0; s < ripetizioni; s++ {
						if s == 0 {
							primo = contatore
						}
						fmt.Print(contatore)
						contatore++
						t = 9
					}
					for j := t; j >= c; j-- {
						fmt.Print(j)
					}
				} else {
					for j := t; j >= c; j-- {
						if j == t {
							primo = t
						}
						fmt.Print(j)
					}
				}
				c = primo + 1
			}
			p++
			fmt.Print("\n")
		} else {
			if q%2 == 0 {
				for j := 0; j < H-2; j++ {
					fmt.Println(strings.Repeat(" ", L-2), c)
					c++
					if c == 10 {
						c = 0
					}
				}
			} else {
				for j := 0; j < H-2; j++ {
					fmt.Println(c)
					c++
					if c == 10 {
						c = 0
					}
				}
			}
			q++
		}
	}
	fmt.Println()
}
