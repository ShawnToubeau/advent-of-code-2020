package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// checks whether an array contains a element
func hasElem(arr []string, elemToFind string) bool {
	for _, elem := range arr {
		if elem == elemToFind {
			return true
		}
	}
	return false
}

func importAnswers(name string) []string {
	text, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
	}

	answerGroups := strings.Split(string(text), "\n\n")

	return answerGroups
}

func countUniqueAnswers(group string) int {
	answers := strings.Split(strings.ReplaceAll(group, "\n", ""), "")
	var unique []string

	for _, answer := range answers {
		if !hasElem(unique, answer) {
			unique = append(unique, answer)
		}
	}

	return len(unique)
}

func countUnanimousAnswers(group string) int {
	// responses from individuals within a group
	responses := strings.Split(group, "\n")
	numPeople := len(responses)
	resMap := make(map[string]int)
	numUnanimous := 0

	// if the group only has one person, return the sum of their responses
	if numPeople == 1 {
		return len(responses[0])
	}

	// loop over every person's response and record their response in a map
	for _, res := range responses {
		answers := strings.Split(res, "")

		for _, answer := range answers {
			resMap[answer] = resMap[answer] + 1
		}
	}

	// loop over the response map and count the number of responses that were shared by everyone
	for i, _ := range resMap {
		if resMap[i] == numPeople {
			numUnanimous++
		}
	}

	return numUnanimous
}

func main() {
	answerGroups := importAnswers("./answers.txt")
	sum := 0

	for _, group := range answerGroups {
		sum += countUnanimousAnswers(group)
	}

	fmt.Printf("Sum %v\n", sum)
}
