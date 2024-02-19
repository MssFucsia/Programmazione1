package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	temperature := make(map[int]int)

	for scanner.Scan() {
		t, _ := strconv.Atoi(scanner.Text())
		temperature[t]++
	}

	var keys []int
	for k := range temperature {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	fmt.Print("minima: ", keys[0], ", ", temperature[keys[0]], " volte\n")
	fmt.Print("massima: ", keys[len(keys)-1], ", ", temperature[keys[len(keys)-1]], " volte\n")

	for k, v := range temperature {
		fmt.Print(k, ":", v, " ")
	}
	fmt.Println()

	//ordina le chiavi in base al valore
	sort.SliceStable(keys, func(i, j int) bool {
		return temperature[keys[i]] > temperature[keys[j]]
	})

	m := make(map[int]int)

	for i := range keys {
		m[i] = temperature[i]
	}
	fmt.Println(m)
	var frequenze_massime []int

	max_frequenza := temperature[keys[0]]

	frequenze_massime = append(frequenze_massime, max_frequenza)

	for _, v := range temperature {
		if v == max_frequenza {
			frequenze_massime = append(frequenze_massime, v)
		}
	}

	fmt.Println(frequenze_massime, max_frequenza)

	/*pairs := make([][2]int, len(temperature))
	i := 0
	for k, v := range temperature {
		pairs[i] = [2]int{k, v}
		i++
	}

	// Ordina la slice in base al valore (secondo elemento della coppia)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})

	var frequenze_massime []int

	max_frequenza := pairs[0][1]

	for i := 0; i < len(pairs); i++ {
		if pairs[i][1] == max_frequenza {
			frequenze_massime = append(frequenze_massime, pairs[i][0])
		}
	}

	fmt.Println("temperature", frequenze_massime, "con frequenza", max_frequenza)*/

}
