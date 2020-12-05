package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func importPattern(name string, xMove int) []string {
	text, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
	}

	lines := strings.Split(string(text), "\n")
	lineChars := len(lines[0])
	numLines := len(lines)
	numRepeated := numLines / lineChars
	if numLines%lineChars != 0 {
		numRepeated++
	}

	var trees []string
	for _, line := range lines {
		// we need to also account for the number of moves we'll be making in the X direction
		trees = append(trees, strings.Repeat(line, numRepeated*xMove))
	}

	return trees
}

func traverseSlope(slope []string, xMove int, yMove int) int {
	numTrees := 0
	currX := 0
	currY := 0

	for currY < len(slope)-1 {
		currPos := slope[currY+yMove][currX+xMove]

		if currPos == '#' {
			numTrees++
		}
		currX += xMove
		currY += yMove
	}

	return numTrees
}

func findBiggestXMove(paths [][]int) int {
	biggest := 0

	for _, path := range paths {
		if path[0] > biggest {
			biggest = path[0]
		}
	}

	return biggest
}

func findArrProduct(arr []int) int {
	product := 0

	for _, num := range arr {
		if product == 0 {
			product = num
		} else {
			product = product * num
		}
	}

	return product
}

func main() {
	paths := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	slope := importPattern("./trees.txt", findBiggestXMove(paths))
	var numTreeArr []int

	for _, path := range paths {
		numTrees := traverseSlope(slope, path[0], path[1])

		numTreeArr = append(numTreeArr, numTrees)
	}

	fmt.Printf("Num Trees: %v\n", findArrProduct(numTreeArr))
}
