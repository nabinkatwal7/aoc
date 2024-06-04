package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Dimensions struct {
	Length int
	Width  int
	Height int
}

func calculateSmallestArea (dimensions Dimensions) (int){
	numbers := []int{dimensions.Length, dimensions.Width, dimensions.Height}

	sort.Ints(numbers)
	return numbers[0] * numbers[1]
}

func calculateSmallestPerimeter (dimensions Dimensions) (int){
	numbers := []int{dimensions.Length, dimensions.Width, dimensions.Height}

	sort.Ints(numbers)
	return 2 * numbers[0] + 2 * numbers[1]
}

func calculateCubicFeet (dimensions Dimensions) (int){
	return dimensions.Length * dimensions.Width * dimensions.Height
}

func RunTwo() {
	input, err := os.ReadFile("inputs/2-IWasToldThereWouldBeNoMath.txt")

	var dimensions []Dimensions

	var surfaceArea = 0
	var smallestArea = 0
	var smallestPerimeter = 0
	var cubicFeet = 0

	if err != nil {
		panic(err)
	}

	split := strings.Split(string(input), "\n")

	for i := 0; i < len(split); i++ {
		spaceSplit := strings.Split(split[i], " ")
		multiplySplit := strings.Split(spaceSplit[0], "x")

		if len(multiplySplit) != 3 {
			fmt.Printf("Error on line %d\n", i)
			continue
		}

		l, err := strconv.Atoi(multiplySplit[0])

		if err != nil {
			panic(err)
		}

		w, err := strconv.Atoi(multiplySplit[1])

		if err != nil {
			panic(err)
		}

		h, err := strconv.Atoi(multiplySplit[2])

		if err != nil {
			panic(err)
		}

		dimensions = append(dimensions, Dimensions{l, w, h})
	}

	for i := 0; i < len(dimensions); i++ {
		surfaceArea += 2 * dimensions[i].Length * dimensions[i].Width + 2 * dimensions[i].Width * dimensions[i].Height + 2 * dimensions[i].Height * dimensions[i].Length

		smallestArea += calculateSmallestArea(dimensions[i])
		smallestPerimeter += calculateSmallestPerimeter(dimensions[i])
		cubicFeet += calculateCubicFeet(dimensions[i])
	}

	fmt.Println(surfaceArea)
	fmt.Println(smallestArea)
	fmt.Println("Total area is", surfaceArea + smallestArea)

	fmt.Println("Ribbon required is ", smallestPerimeter)
	fmt.Println("Cubic feet for bolt is ", cubicFeet)

	fmt.Println("Total cubic feet is ", smallestPerimeter + cubicFeet)
}
