package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	l, _ := strconv.Atoi(os.Args[1])
	n, _ := strconv.Atoi(os.Args[2])
	c := []byte(os.Args[3])[0]
	var k int
	if l > 0 && n > 0 {
		for i := 0; i < n; i++ {
			if i == 0 {
				fmt.Println(DrawSegment(c, k, l))
				k++
				for j := 0; j < 2; j++ {
					fmt.Println(DrawPoint(c, k))
				}

			} else {
				fmt.Println(DrawSegment(c, k+1, l))
				k++
				for j := 0; j < 2; j++ {
					fmt.Println(DrawPoint(c, k+1))
				}

			}

		}
	}
}

func DrawPoint(c byte, k int) string {
	return strings.Repeat(" ", k+1) + string(c)
}

func DrawSegment(c byte, k, l int) string {
	return strings.Repeat(" ", k) + strings.Repeat(string(c), l)
}
