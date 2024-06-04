package main

import (
	"fmt"
	"os"
)

func countUniqueCoordinates(coords []string) int {
  // Create a map to store coordinates as keys
  uniqueCoords := make(map[interface{}]bool)
  
  // Iterate through the coord_list
  for _, coord := range coords {
    uniqueCoords[coord] = true
  }
  
  // Return the length of the map (number of unique keys)
  return len(uniqueCoords)
}

func RunThree () {
	input, err := os.ReadFile("inputs/day3.txt")

	var coordinates = []int{0, 0}
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
			coordinates[1]++
		}

		if direction == south {
			coordinates[1]--
		}

		if direction == west {
			coordinates[0]--
		}

		if direction == east {
			coordinates[0]++
		}

		coordinatesList = append(coordinatesList, fmt.Sprintf("%d,%d", coordinates[0], coordinates[1]))
	}

	fmt.Println(countUniqueCoordinates(coordinatesList))
}