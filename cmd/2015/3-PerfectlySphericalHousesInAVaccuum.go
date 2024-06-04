package main

import (
	"fmt"
	"os"
)

func countUniqueCoordinates(coords []string) int {
  uniqueCoords := make(map[interface{}]bool)
  for _, coord := range coords {
    uniqueCoords[coord] = true
  }
  return len(uniqueCoords)
}

func RunThree () {
	input, err := os.ReadFile("inputs/day3.txt")

	var coordinates = []int{0, 0}
	var robotCoordinates = []int{0, 0}
	var coordinatesList = []string{"0,0"}

	north := "^"
	south := "v"
	west := "<"
	east := ">"

	if err != nil {
		panic(err)
	}

	for i := 0; i < len(input); i++ {
		direction := string(input[i])
		if direction == north {
			if i % 2==0{
				coordinates[1]++
			}else{
				robotCoordinates[1]++
			}
		}

		if direction == south {
			if i % 2==0{
				coordinates[1]--
			}else{
				robotCoordinates[1]--
			}
		}

		if direction == west {
			if i % 2==0{
				coordinates[0]--
			}else{
				robotCoordinates[0]--
			}
		}

		if direction == east {
			if i % 2==0{
				coordinates[0]++
			}else{
				robotCoordinates[0]++
			}
		}

		if i % 2==0{
			coordinatesList = append(coordinatesList, fmt.Sprintf("%d,%d", coordinates[0], coordinates[1]))
		}else{
			coordinatesList = append(coordinatesList, fmt.Sprintf("%d,%d", robotCoordinates[0], robotCoordinates[1]))
		}
	}

	fmt.Println(countUniqueCoordinates(coordinatesList))
}