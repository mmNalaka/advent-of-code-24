package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Read input file
func readInput() ([]string, []string) {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Split by empty line
	parts := strings.Split(strings.TrimSpace(string(file)), "\n\n")
	rules := strings.Split(strings.TrimSpace(parts[0]), "\n")
	updates := strings.Split(strings.TrimSpace(parts[1]), "\n")

	return rules, updates
}

// Parse rules into a map of before->after relationships
func parseRules(rules []string) map[int][]int {
	dependencies := make(map[int][]int)

	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])

		// Add after to the list of pages that must come after before
		dependencies[before] = append(dependencies[before], after)
	}

	fmt.Println(dependencies)
	return dependencies
}

// Check if an update is valid according to the rules
func isValidUpdate(update []int, rules map[int][]int) bool {
	// For each pair of pages in the update
	for i := 0; i < len(update); i++ {
		for j := i + 1; j < len(update); j++ {
			before := update[i]
			after := update[j]

			// If there's a rule saying after should come before before, the update is invalid
			if deps, ok := rules[after]; ok {
				for _, dep := range deps {
					if dep == before {
						return false
					}
				}
			}
		}
	}
	return true
}

// Modify the order of the pages in an update to make it valid
func updateOrder(update []int, rules map[int][]int) []int {
	for i := 0; i < len(update); i++ {
		for j := i + 1; j < len(update); j++ {
			before := update[i]
			after := update[j]

			// If there's a rule saying after should come before before, swap them
			if deps, ok := rules[after]; ok {
				for _, dep := range deps {
					if dep == before {
						update[i] = after
						update[j] = before
					}
				}
			}
		}
	}

	return update
}

// Get middle page number from an update
func getMiddlePage(update []int) int {
	return update[len(update)/2]
}

// Part 1: Find sum of middle pages from valid updates
func part1(rules []string, updates []string) int {
	// Parse rules into dependency map
	ruleMap := parseRules(rules)

	sum := 0

	// Process each update
	for _, update := range updates {
		// Convert update string to integers
		var pages []int
		for _, p := range strings.Split(update, ",") {
			page, _ := strconv.Atoi(p)
			pages = append(pages, page)
		}

		// If update is valid, add its middle page to sum
		if isValidUpdate(pages, ruleMap) {
			sum += getMiddlePage(pages)
		}
	}

	return sum
}

func part2(rules []string, updates []string) int {
	// Parse rules into dependency map
	ruleMap := parseRules(rules)

	sum := 0
	invalidSum := 0

	// Process each update
	for _, update := range updates {
		// Convert update string to integers
		var pages []int
		for _, p := range strings.Split(update, ",") {
			page, _ := strconv.Atoi(p)
			pages = append(pages, page)
		}

		// If update is valid, add its middle page to sum
		if isValidUpdate(pages, ruleMap) {
			sum += getMiddlePage(pages)
		} else {
			// Update order of update to make it valid
			pages = updateOrder(pages, ruleMap)
			// If update is still not valid, it must be invalid
			if isValidUpdate(pages, ruleMap) {
				invalidSum += getMiddlePage(pages)
			}
		}
	}

	return invalidSum
}

func main() {
	rules, updates := readInput()

	result := part1(rules, updates)
	fmt.Printf("Part 1: %d\n", result)

	result = part2(rules, updates)
	fmt.Printf("Part 2: %d\n", result)
}
