package main

// Input file: day-1/input.txt

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

// Read input file
func readInputFile() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

// Part 1
func part1() {
	data := readInputFile()
	r := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	macthes := r.FindAllString(data, -1)

	// fmt.Println(macthes)

	total := 0
	for _, v := range macthes {
		// v = mul(5,5)
		r := regexp.MustCompile(`[0-9]{1,3}`)
		numbers := r.FindAllString(v, -1)

		// Convert strings to integers
		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])

		if err1 != nil || err2 != nil {
			println("Error converting numbers:", err1, err2)
			continue
		}

		total += num1 * num2
	}
	println("Part 1:", total)
}

// Part 2
func part2() {
	// Print the total distance
	println("Part 2:")
}

func main() {
	part1()
	part2()
}
