package main

import "fmt"

type gasFuel struct {
	mpg   uint8
	price uint8
}

func main() {
	MygasFuel := gasFuel{mpg: 100, price: 20}
	printStruct(MygasFuel)
}

func printStruct(anygasFuel gasFuel) {
	fmt.Printf("MPG: %d, Price: %d\n", anygasFuel.mpg, anygasFuel.price)
}
