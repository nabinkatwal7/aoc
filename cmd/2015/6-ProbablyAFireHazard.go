package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x, y int
}

func (c coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func extractXY(s string) (int, int) {
	parts := strings.Split(s, ", ")
	if len(parts) != 2 {
		panic("Invalid input")
	}
	a, err := strconv.Atoi(parts[0])

	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(parts[1])

	if err != nil {
		panic(err)
	}

	return a, b
}

func RunSix() {
	b, err := os.ReadFile("inputs/day6.txt")

	if err != nil {
		panic(err)
	}

	instructions := strings.Split(string(b), "\n")
	lightGrid := make(map[string]int)

	fmt.Println("Initializing grid")

	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			lightGrid[fmt.Sprintf("%d,%d", x, y)] = 0
		}
	}

	fmt.Println("Running instructions")
	for _, instruction := range instructions {
		fmt.Print(".")
		runInstruction(lightGrid, instruction)
	}

	fmt.Println("Counting lights")
	lightOnCount := 0

	for _, v := range lightGrid {
		lightOnCount += v
	}

	fmt.Printf("Lights on: %d\n", lightOnCount)

}

func runInstruction(grid map[string]int, instruction string) {
	// for better tokenizing
	instruction = strings.Replace(instruction, "toggle", "toggle x", 1)

	parts := strings.Split(instruction, " ")
	if len(parts) != 5 {
		log.Printf("unexpected number of parts: %s", instruction)
		return
	}

	action := parts[0]
	state := parts[1]
	start := parts[2]
	end := parts[4]

	x1, y1 := extractXY(start)
	x2, y2 := extractXY(end)

	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			if action == "toggle" {
				grid[coordinate{x, y}.String()] = grid[coordinate{x, y}.String()] + 2
			}
			if state == "on" {
				grid[coordinate{x, y}.String()] = grid[coordinate{x, y}.String()] + 1
			}
			if state == "off" {
				grid[coordinate{x, y}.String()] = grid[coordinate{x, y}.String()] - 1
				if grid[coordinate{x, y}.String()] < 0 {
					grid[coordinate{x, y}.String()] = 0
				}
			}
		}
	}
}

