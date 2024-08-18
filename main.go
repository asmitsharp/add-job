package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num1Str := os.Getenv("NUMBER1")
	num2Str := os.Getenv("NUMBER2")

	num1, err := strconv.Atoi(num1Str)
	if err != nil {
		fmt.Println("Error converting the NUMBER1:", err)
		os.Exit(1)
	}

	num2, err := strconv.Atoi(num2Str)
	if err != nil {
		fmt.Println("Error converting the NUMBER2:", err)
		os.Exit(1)
	}

	result := num1 + num2

	fmt.Printf("The sum of %d and %d is : %d\n", num1, num2, result)
}
