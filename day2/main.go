package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

var input string

// Read input from a file and store it in the global `input` variable
func readInputFile(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	input = string(data)
	return nil
}

func processInput(input string) [][]int {
	var report [][]int
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		// Split the line into individual string elements (separated by spaces)
		strValues := strings.Split(line, " ")

		// Create a temporary slice to hold the integers for this line
		var row []int

		// Convert each string value to an integer and append to the row
		for _, str := range strValues {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
			}
			row = append(row, num)
		}

		// Append the row to the 2D slice
		report = append(report, row)
	}

	return report
}

func isAscending(level []int) bool {
	for i := 1; i < len(level); i++ {
		if level[i] < level[i-1] {
			return false
		}
	}
	return true
}

// isDescending checks if the slice is sorted in descending order
func isDescending(level []int) bool {
	for i := 1; i < len(level); i++ {
		if level[i] > level[i-1] {
			return false // If any element is greater than the previous one, it's not descending
		}
	}
	return true // If no such pair is found, the slice is descending
}

func isSafe(level []int) bool {
	asc, desc := isAscending(level), isDescending(level)
	if !asc && !desc {
		return false
	} else {
		for j := 1; j < len(level); j++ {
			a := level[j]
			b := level[j-1]
			diff := int(math.Abs(float64(a - b)))
			if diff < 1 || diff > 3 {
				return false
			}
		}
	}
	return true
}

func part1(input string) int {
	report := processInput(input)

	safe := 0
	for i := 0; i < len(report); i++ {
		isValid := isSafe(report[i])
		if isValid {
			safe++
		}
	}
	return safe
}

func canTolerate(level []int) bool {
	for i := 0; i < len(level); i++ {
		temp := make([]int, len(level))
		copy(temp, level)
		temp = append(temp[:i], temp[i+1:]...)
		if isSafe(temp) {
			return true
		}
	}
	return false
}

func part2(input string) int {
	report := processInput(input)
	safe := 0
	for i := 0; i < len(report); i++ {
		isValid := isSafe(report[i])
		if isValid {
			safe++
		} else {
			tolerate := canTolerate(report[i])
			if tolerate {
				safe++
			}
		}
	}
	a := canTolerate(report[1])
	fmt.Println(a)
	return safe
}

func main() {
	fileName := "input.txt"
	err := readInputFile(fileName)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Add flag for selecting part 1 or part 2
	var part int
	flag.IntVar(&part, "part", 1, "Choose part 1 or 2")
	flag.Parse()

	fmt.Println("Running part", part)

	var result int
	switch part {
	case 1:
		result = part2(input)
	default:
		fmt.Println("Invalid part selected")
		return
	}

	// Copy the result to the clipboard
	err = clipboard.WriteAll(fmt.Sprintf("%v", result))
	if err != nil {
		fmt.Println("Error copying to clipboard:", err)
	} else {
		fmt.Println("Result copied to clipboard")
	}

	// Print the output
	fmt.Println("Output:", result)
}
