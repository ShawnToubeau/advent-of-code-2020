package main

import (
	"fmt"
	"github.com/shawntoubeau/advent-of-code-2020/utils"
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

func checkNextNumberValidity(group []int, nextNumber int) bool {
	num1 := 0
	num2 := 0

	//fmt.Printf("Checking validity\n")
	//fmt.Printf("Group %v - next num: %v\n", group, nextNumber)

	for i := 0; i < len(group); i++ {
		num1 = group[i]
		for j := i; j < len(group); j++ {
			num2 = group[j]

			if i != j && num1+num2 == nextNumber {
				return true
			}
		}
	}

	return false
}

func main() {
	numbers := convertToNumArr(utils.ReadFile("./numbers.txt"))
	offset := 0

	//fmt.Printf("%v\n", numbers)

	for offset < len(numbers)-preamble {
		group := getNumberGroup(numbers, offset)
		nextNum := numbers[preamble+offset]

		if !checkNextNumberValidity(group, nextNum) {
			fmt.Printf("Next number is not valid %v\n", nextNum)
		}

		offset++
	}
}
