package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dimensions struct {
	Length int
	Width  int
	Height int
}

func separateDimensions(data string) []Dimensions {
	lines := strings.Split(data, "\n")

	var dimensions []Dimensions

	for _, line := range lines {
		values := strings.Fields(line)
		individualValues := strings.Split(values[0], "x")
		
		l, err := strconv.Atoi(individualValues[0])

		if err != nil {
			fmt.Println(err)
			continue
		}

		w, err := strconv.Atoi(individualValues[1])

		if err != nil {
			fmt.Println(err)
			continue
		}

		h, err := strconv.Atoi(individualValues[2])

		if err != nil {
			fmt.Println(err)
			continue
		}

		dimensions = append(dimensions, Dimensions{l, w, h})
	}
	return dimensions
}

func RunTwo() {
	input, err := os.ReadFile("inputs/2-IWasToldThereWouldBeNoMath.txt")

	if err != nil {
		panic(err)
	}

	allDimensions := separateDimensions(string(input))

	for _, dim := range allDimensions {
		fmt.Printf("l: %d, w: %d, h: %d\n", dim.Length, dim.Width, dim.Height)
	}
}
