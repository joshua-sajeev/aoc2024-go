package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
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

func part1(input string) int {
	// var results [][]int
	input = strings.TrimSpace(input)

	lines := strings.Split(input, "\n")

	lineSum := 0
	for _, line := range lines {
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if len(match) == 3 {
				a, _ := strconv.Atoi(match[1])
				b, _ := strconv.Atoi(match[2])
				lineSum += a * b
			}
		}
	}

	return lineSum
}

type Expr struct {
	startIndex int
	endIndex   int
	a          int
	b          int
	doIndex    [][]int
	dontIndex  [][]int
	valid      bool
}

var lineStruct [50]Expr

// Add the mul() and instruction indices to the struct, and check if mul() is valid
func addToStruct(line string) {
	mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := mulRe.FindAllStringSubmatch(line, -1)
	matchIndex := mulRe.FindAllStringSubmatchIndex(line, -1)
	lenMatch := len(matches)

	// Find all do() and don't() positions
	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`do(n't)?\(\)`)
	doIndexmatches := reDo.FindAllStringIndex(line, -1)
	dontIndexmatches := reDont.FindAllStringIndex(line, -1)

	// For each mul match, store relevant data in the struct
	for i := 0; i < lenMatch; i++ {
		// Parse the numbers from mul()
		intA, _ := strconv.Atoi(matches[i][1])
		intB, _ := strconv.Atoi(matches[i][2])

		// Store mul data in lineStruct
		lineStruct[i].a = intA
		lineStruct[i].b = intB
		lineStruct[i].startIndex = matchIndex[i][0]
		lineStruct[i].endIndex = matchIndex[i][1]

		// Initialize the do/don't indices for each mul
		lineStruct[i].doIndex = doIndexmatches
		lineStruct[i].dontIndex = dontIndexmatches

		// Check if mul is enabled or disabled based on do() and don't() instructions
		lineStruct[i].valid = isValidMul(lineStruct[i], doIndexmatches, dontIndexmatches)
	}
}

// Check if a mul operation is valid (enabled) based on the latest do() or don't() position
func isValidMul(expr Expr, doIndexmatches, dontIndexmatches [][]int) bool {
	mulIndex := expr.startIndex

	// If there are no do() or don't() instructions, mul is enabled
	if len(doIndexmatches) == 0 && len(dontIndexmatches) == 0 {
		return true
	}

	// Check if the latest do()/don't() is before the mul() index
	for _, do := range doIndexmatches {
		if mulIndex < do[0] {
			return true // mul is enabled if do() comes before it
		}
	}
	for _, dont := range dontIndexmatches {
		if mulIndex < dont[0] {
			return false // mul is disabled if don't() comes before it
		}
	}
	return true
}

// Compute the total sum of valid mul() operations
func part2(input string) int {
	// Took help from https://www.reddit.com/r/adventofcode/comments/1h5frsp/2024_day_3_solutions/
	// Realised i have skill issues
	strs := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`).FindAllString(string(input), -1)
	fmt.Println(strs)
	result := 0
	mulEnabed := true
	for _, s := range strs {
		if mulEnabed = s != "don't()" && (mulEnabed || s == "do()"); mulEnabed {
			var a, b int
			fmt.Sscanf(s, "mul(%d,%d)", &a, &b)
			result += a * b
		}
	}
	fmt.Println(result)
	return result
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
