package main

import (
	"errors"
	"fmt"
)

func main() {
	printValue := "Hello world"
	printMe(printValue) //Hello World

	var numerator int = 11
	var denominator int = 0
	var result, remainder, err = intDivision(numerator, denominator)

	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("the result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by Zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
