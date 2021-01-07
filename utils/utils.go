package utils

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Returns the contents of a file as a string array. Separated by new line characters.
func ReadFile(fileName string) []string {
	text, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
	}

	data := strings.Split(string(text), "\n")

	return data
}

// Converts an array of strings to an array of integers
func ConvertToNumArr(numbers []string) []int {
	newArr := make([]int, len(numbers))

	for i, strNum := range numbers {
		num, _ := strconv.Atoi(strNum)
		newArr[i] = num
	}

	return newArr
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// ArrContainsElem check whether a string array contains a specific string
func ArrContainsElem(arr []string, elem string) bool {
	for i := range arr {
		if arr[i] == elem {
			return true
		}
	}

	return false
}

// ArrNumOfOccurrences returns the number of occurrences an element is present in an array
func ArrNumOfOccurrences(arr []string, elem string) int {
	count := 0

	for _, s := range arr {
		if s == elem {
			count++
		}
	}

	return count
}
