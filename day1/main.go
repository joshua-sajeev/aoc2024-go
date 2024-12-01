package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	left  []int
	right []int
)

func processLine(line string) {
	parts := strings.Fields(line)
	if len(parts) == 2 {
		leftNum, err1 := strconv.Atoi(parts[0])
		rightNum, err2 := strconv.Atoi(parts[1])
		if err1 == nil && err2 == nil {
			left = append(left, leftNum)
			right = append(right, rightNum)
		} else {
			fmt.Println("Error converting line:", line)
		}
	} else {
		fmt.Println("Invalid line format:", line)
	}
}

func part1() int {
	sum := 0
	for i := range left {
		diff := left[i] - right[i]
		absDiff := math.Abs(float64(diff))
		sum += int(absDiff)
	}
	return sum
}

func part2() int {
	sum := 0
	for i := 0; i < len(left); i++ {
		x := left[i]
		y := 0
		for j := 0; j < len(right); j++ {
			if x == right[j] {
				y++
			}
		}
		sum += x * y
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		processLine(line)
	}
	sort.Ints(left)
	sort.Ints(right)
	part1 := part1()
	part2 := part2()
	fmt.Println(part1)
	fmt.Println(part2)
}
