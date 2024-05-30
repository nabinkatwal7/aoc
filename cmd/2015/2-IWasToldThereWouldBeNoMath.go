package main

import (
	"fmt"
	"os"
)

func RunTwo() {
	input, err := os.ReadFile("inputs/2-IWasToldThereWouldBeNoMath.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(input))
}
