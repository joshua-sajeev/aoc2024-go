package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var directions = [][2]int{
	{-1, 0},  // Up
	{1, 0},   // Down
	{0, -1},  // Left
	{0, 1},   // Right
	{-1, -1}, // Up-left (diagonal)
	{-1, 1},  // Up-right (diagonal)
	{1, -1},  // Down-left (diagonal)
	{1, 1},   // Down-right (diagonal)
}

func checkPatternInDirection(grid [][]string, r, c int, dr, dc int) bool {
	word := "XMAS"
	for i := 0; i < len(word); i++ {
		newR := r + i*dr
		newC := c + i*dc

		if newR < 0 || newR >= len(grid) || newC < 0 || newC >= len(grid[0]) {
			return false
		}

		if grid[newR][newC] != string(word[i]) {
			return false
		}
	}

	return true
}

func countXMASOccurrences(grid [][]string) int {
	count := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			for _, dir := range directions {
				dr, dc := dir[0], dir[1]
				if checkPatternInDirection(grid, r, c, dr, dc) {
					count++
				}
			}
		}
	}
	return count
}

func stringTo2DArray(input string) [][]string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var result [][]string
	for _, line := range lines {
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		result = append(result, row)
	}
	return result
}

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sb.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return sb.String(), nil
}

func main() {
	filename := "input.txt"
	input, err := readFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	grid := stringTo2DArray(input)

	result := countXMASOccurrences(grid)

	fmt.Printf("The word 'XMAS' appears %d times.\n", result)
}
