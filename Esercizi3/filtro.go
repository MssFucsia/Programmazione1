package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	euro, _ := strconv.Atoi(os.Args[1])
	fmt.Println("Taglio: 100 Quantità:", euro/100)
	fmt.Println("Taglio: 50 Quantità:", (euro%100)/50)
	fmt.Println("Taglio: 20 Quantità:", ((euro%100)%50)/20)
	fmt.Println("Taglio: 10 Quantità:", (((euro%100)%50)%20)/10)
	fmt.Println("Taglio: 5 Quantità:", ((((euro%100)%50)%20)%10)/5)
	fmt.Println("Taglio: 2 Quantità:", (((((euro%100)%50)%20)%10)%5)/2)
	fmt.Println("Taglio: 1 Quantità:", (((((euro%100)%50)%20)%10)%5)%2)

}
