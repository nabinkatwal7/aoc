package main

import (
	"fmt"
	"os"
)

func RunSix(){
	input, err := os.ReadFile("inputs/day6.txt")

	if err != nil {
		panic(err)
	}

	var lights = make([][]int, 1000)

	fmt.Println(string(input))

	fmt.Println(lights)

	// for i := 0; i < 1000; i++ {
	// 	for j := 0; j < 1000; j++ {
	// 		fmt.Println(lights[i][j])
	// 	}
	// }
}