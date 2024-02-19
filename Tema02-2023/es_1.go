package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Print("Inserire input: ")
	fmt.Scan(&s)
	sottostringhe_max := SottostringaMaggiore(SottostringheAscii(SottostringheTotali(s)))
	if len(sottostringhe_max) != 0 {
		for _, r := range sottostringhe_max {
			fmt.Println(r)
		}
	}

}

func SottostringheTotali(s string) (ss []string) {
	sr := []rune(s)

	for i := 0; i < len(sr); i++ {
		for j := i; j <= len(sr); j++ {
			ss = append(ss, string(sr[i:j]))
		}
	}
	return ss
}

func OnlyAscii(s string) bool {
	if len(s) == len([]rune(s)) {
		return true
	}
	return false
}

func SottostringheAscii(ss []string) (ss_ascii []string) {
	for _, r := range ss {
		if OnlyAscii(r) {
			ss_ascii = append(ss_ascii, r)
		}
	}
	return
}

func SottostringaMaggiore(ss_ascii []string) (slice []string) {
	ss_max := ss_ascii[0]
	for _, r := range ss_ascii {
		if len(r) > len(ss_max) {
			ss_max = r
		}
	}

	for _, r := range ss_ascii {
		if len(r) == len(ss_max) {
			slice = append(slice, r)
		}
	}
	return
}
