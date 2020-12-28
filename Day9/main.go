package main

import (
	"fmt"
	"github.com/shawntoubeau/advent-of-code-2020/utils"
	"sort"
	"strconv"
)

const preamble = 25

// Converts an array of strings to an array of integers
func convertToNumArr(numbers []string) []int {
	newArr := make([]int, len(numbers))

	for i, strNum := range numbers {
		num, _ := strconv.Atoi(strNum)
		newArr[i] = num
	}

	return newArr
}

func getNumberGroup(numbers []int, offset int) []int {
	return numbers[offset : preamble+offset]
}

// Given our preamble size and number input, check to see if the number immediately following the group is valid
// i.e the sum of any two numbers within the group equal the next number
func checkNextNumberValidity(group []int, nextNumber int) bool {
	num1 := 0
	num2 := 0

	for i := 0; i < len(group); i++ {
		num1 = group[i]
		for j := i + 1; j < len(group); j++ {
			num2 = group[j]

			if num1+num2 == nextNumber {
				return true
			}
		}
	}

	return false
}

// Finds a contiguous set of numbers within our input whose sum equals the invalid number we found in part 1.
// Uses a double for loop to incrementally crawl our numbers and record the sum as it goes. If the sum equals the invalid number,
// we found our range
func findContiguousRange(numbers []int, inValidNum int) []int {
	var numRange []int

	for i := 0; i < len(numbers); i++ {
		tempSum := 0
		// add the start of the range
		tempSum = tempSum + numbers[i]

		for j := i + 1; j < len(numbers); j++ {
			// add the following numbers that follow our range start
			tempSum = tempSum + numbers[j]

			// if our sum go over, break out of the inner loop, increment the outer, & reset our sum
			if tempSum > inValidNum {
				break
			} else if tempSum == inValidNum {
				fmt.Printf("Range found i: %v, j: %v\n", i, j)

				// append all the numbers within the range to a slice
				for i < j {
					numRange = append(numRange, numbers[i])
					i++
				}
				fmt.Printf("Num range: %v\n", numRange)

				return numRange
			}
		}
	}

	return numRange
}

func main() {
	numbers := convertToNumArr(utils.ReadFile("./numbers.txt"))
	offset := 0
	invalidNum := 0

	// traverse our number list using an offset
	for offset < len(numbers)-preamble {
		// get the current group of numbers using the offset
		group := getNumberGroup(numbers, offset)
		// next number immediately after our number group
		nextNum := numbers[preamble+offset]

		// if the next number isn't valid, we've found the number we've been looking for
		if !checkNextNumberValidity(group, nextNum) {
			fmt.Printf("Next number is not valid %v\n", nextNum)
			invalidNum = nextNum
			break
		}

		offset++
	}
	// finds contiguous range of numbers whose sum equal the invalid number we found
	numRange := findContiguousRange(numbers, invalidNum)
	// sort the number range
	sort.Ints(numRange)
	// answer is the sum of the smallest number and biggest number in the contiguous range
	fmt.Printf("XMAS Weakness: %v\n", numRange[0]+numRange[len(numRange)-1])
}
