package main

import (
	"fmt"
	"os"
)

func main() {
	s := []rune(os.Args[1])
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i:]), string(s[0:i]), "\n")
	}

}
