package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func importPasses(name string) []string {
	text, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
	}

	passes := strings.Split(string(text), "\n")

	return passes
}

// first number in the return is the min, second is the max
func getHalf(min float64, max float64, half string) []float64 {
	median1 := (((max - min) / 2) - 0.5) + min
	median2 := (((max - min) / 2) + 0.5) + min

	if half == "B" || half == "R" {
		return []float64{median2, max}
	} else if half == "F" || half == "L" {
		return []float64{min, median1}
	}

	return []float64{0, 0}
}

func main() {
	const numRows = 128
	const numCols = 8
	biggestId := float64(0)
	passes := importPasses("./passes.txt")
	seatingMap := [numRows][numCols]int{}

	// loop over each pass
	for _, pass := range passes {
		// reset variables
		rowMin := float64(0)
		rowMax := float64(numRows - 1)
		colMin := float64(0)
		colMax := float64(numCols - 1)

		// split each pass into a character array
		passArr := strings.Split(pass, "")
		for i, passElem := range passArr {
			// handle row part of the pass
			if i < 7 {
				arr := getHalf(rowMin, rowMax, passElem)
				rowMin = arr[0]
				rowMax = arr[1]
			} else { // handle col part of the pass
				arr := getHalf(colMin, colMax, passElem)
				colMin = arr[0]
				colMax = arr[1]
			}
		}
		seatId := (rowMax * 8) + colMax

		if seatId > biggestId {
			biggestId = seatId
		}
		// record position of the current ticket; max and mins will have the same value
		seatingMap[int(rowMax)][int(colMax)] = 1
	}
	fmt.Printf("Biggest: %v\n", biggestId)

	// loop over the seating map and print all un-taken seat IDs
	for i := 0; i < numRows-1; i++ {
		for j := 0; j < numCols-1; j++ {
			if seatingMap[i][j] == 0 && seatingMap[i][j+1] == 1 && seatingMap[i][j-1] == 1 {
				fmt.Printf("My Seat ID: %v\n", (i*8)+j)
			}
		}
	}

}
