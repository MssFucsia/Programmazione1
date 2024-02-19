package main

import (
	"fmt"
	"os"
)

func main() {
	n := os.Args[1]
	fmt.Println(Romano2decimale(n))
}

func Romano2decimale(n string) int {
	var s []int
	for _, char := range n {
		switch char {
		case 'I':
			s = append(s, 1)
		case 'V':
			s = append(s, 5)
		case 'X':
			s = append(s, 10)
		case 'L':
			s = append(s, 50)
		case 'C':
			s = append(s, 100)
		case 'D':
			s = append(s, 500)
		case 'M':
			s = append(s, 1000)
		}
	}
	var tot int
	for i := 0; i <= len(s); i++ {
		if i == len(s)-1 {
			tot += s[i]
			break
		}
		if s[i] < s[i+1] {
			tot -= s[i]
		} else {
			tot += s[i]
		}
	}
	return tot
}
