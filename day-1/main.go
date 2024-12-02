package main

// Input file: day-1/input.txt

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// Read input file
func readInputFile() []string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data), "\n")
}

// Calculate the absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Part 1
func part1() {
	input := readInputFile()
	var left []int
	var right []int
	var totalDistance int

	// Loop through the input array split each line in to two seperatated by a empty space and create tow arrays
	for _, v := range input {
		// Split the line in to two
		split := strings.Split(v, "   ")

		// Convert the string to int and append to the array
		leftInt, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, leftInt)

		rightInt, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		right = append(right, rightInt)
	}

	// Sort the arrays
	for i := 0; i < len(left); i++ {
		for j := 0; j < len(left)-1; j++ {
			if left[j] > left[j+1] {
				left[j], left[j+1] = left[j+1], left[j]
			}
		}
	}

	for i := 0; i < len(right); i++ {
		for j := 0; j < len(right)-1; j++ {
			if right[j] > right[j+1] {
				right[j], right[j+1] = right[j+1], right[j]
			}
		}
	}

	// Loop through the arrays and compare the values and add the absolute distance to the total distance
	for i := 0; i < len(left); i++ {
		totalDistance += abs(left[i] - right[i])
	}

	// Print the total distance
	println(totalDistance)
}

func main() {
	part1()
}
