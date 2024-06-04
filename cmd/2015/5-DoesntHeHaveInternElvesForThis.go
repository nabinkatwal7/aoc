package main

import (
	"fmt"
	"os"
)

func RunFive(){
	input, err := os.ReadFile("inputs/day5.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(input))
}