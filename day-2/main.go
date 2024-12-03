package main

// Input file: day-1/input.txt

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// Read input file and resturn array of arrays of int [][]int
func readInputFile() [][]int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	// Loop through the lines array split each seperatated by a empty space
	var reports [][]int
	for _, v := range lines {
		// Split the line in to two
		split := strings.Split(v, " ")

		// Convert the string to int and append to the array
		var report []int
		for _, v := range split {
			i, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			report = append(report, i)
		}
		reports = append(reports, report)
	}

	return reports
}

// Part 1
func part1() {
	reports := readInputFile()
	safeReports := 0

	for _, v := range reports {
		isAscending := v[0] < v[1]

		for i := 0; i < len(v)-1; i++ {
			delta := v[i+1] - v[i]
			if !isAscending {
				delta = -delta
			}
			if delta > 3 || delta <= 0 {
				goto nextReport
			}
		}
		safeReports++

	nextReport:
	}

	println("Part 1:", safeReports)
}

func checkSafeReport(report []int) bool {
	isAscending := report[0] < report[1]

	for i := 0; i < len(report)-1; i++ {
		delta := report[i+1] - report[i]
		if !isAscending {
			delta = -delta
		}
		if delta > 3 || delta <= 0 {
			return false
		}
	}
	return true
}

// Part 2
func part2() {
	reports := readInputFile()
	safeReports := 0

	// Check if the report is safe
	if checkSafeReport(reports[0]) {
		safeReports++
	} else {

		for _, report := range reports {
			isSafe := false
			// Remove one number from the at a time and see if the report is still safe
			for i := 0; i < len(report); i++ {
				newReport := make([]int, len(report)-1)
				copy(newReport, report[:i])
				copy(newReport[i:], report[i+1:])
				if checkSafeReport(newReport) {
					isSafe = true
					break
				}
			}

			if isSafe {
				safeReports++
			}
		}
	}
	println("Part 2:", safeReports)
}

func main() {
	part1()
	part2()
}
