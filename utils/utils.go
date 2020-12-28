package utils

import (
	"fmt"
	"io/ioutil"
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
