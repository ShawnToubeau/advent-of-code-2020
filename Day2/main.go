package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseExpenseReport(name string) []string {
	text, _ := ioutil.ReadFile(name)

	lines := strings.Split(string(text), ",")

	passwords := make([]string, len(lines))

	for index, l := range lines {
		passwords[index] = l
	}

	return passwords
}

func checkPasswordOldPolicy(line string) bool {
	// Separate the rule count, rule letter, and password
	pieces := strings.Split(line, " ")
	// ex. "2-7" -> 2, 7
	ruleCountLower, _ := strconv.ParseInt(strings.Split(strings.TrimSpace(pieces[0]), "-")[0], 10, 32)
	ruleCountUpper, _ := strconv.ParseInt(strings.Split(strings.TrimSpace(pieces[0]), "-")[1], 10, 32)
	ruleLetter := strings.Replace(pieces[1], ":", "", -1)
	password := pieces[2]

	//fmt.Printf("LOWER: %v\n", ruleCountLower)
	//fmt.Printf("COUNT: %v\n", strings.Count(password, ruleLetter))
	//fmt.Printf("UPPER: %v\n", ruleCountUpper)
	//fmt.Printf("RES: %v\n", int(ruleCountLower) <= strings.Count(password, ruleLetter) && strings.Count(password, ruleLetter) <= int(ruleCountUpper))
	//fmt.Println("")
	if int(ruleCountLower) <= strings.Count(password, ruleLetter) &&
		strings.Count(password, ruleLetter) <= int(ruleCountUpper) {
		return true
	}

	return false
}

func checkPasswordNewPolicy(line string) bool {
	// Separate the rule count, rule letter, and password
	pieces := strings.Split(line, " ")
	// ex. "2-7" -> 2, 7
	firstPos, _ := strconv.ParseInt(strings.Split(strings.TrimSpace(pieces[0]), "-")[0], 10, 32)
	secondPos, _ := strconv.ParseInt(strings.Split(strings.TrimSpace(pieces[0]), "-")[1], 10, 32)
	ruleLetter := strings.Replace(pieces[1], ":", "", -1)
	password := pieces[2]

	fmt.Printf("PASSWORD: %v\n", password)
	fmt.Printf("FIRST POS: %v\n", firstPos)
	fmt.Printf("SECOND POS: %v\n", secondPos)
	fmt.Printf("LETTER: %v\n", ruleLetter)


	//fmt.Printf("UPPER: %v\n", ruleCountUpper)
	//fmt.Printf("RES: %v\n", int(ruleCountLower) <= strings.Count(password, ruleLetter) && strings.Count(password, ruleLetter) <= int(ruleCountUpper))
	//fmt.Println("")
	fmt.Printf("FIRST: %v = %v == %v\n", string(password[firstPos - 1]), ruleLetter, string(password[firstPos - 1]) == ruleLetter)
	fmt.Printf("FIRST: %v = %v == %v\n", string(password[secondPos - 1]), ruleLetter, string(password[secondPos - 1]) == ruleLetter)
	if string(password[firstPos - 1]) == ruleLetter &&
			string(password[secondPos - 1]) != ruleLetter ||
			string(password[firstPos - 1]) != ruleLetter &&
			string(password[secondPos - 1]) == ruleLetter {
		return true
	}

	return false
}

func main() {
	numValid1 := 0
	numValid2 := 0
	passwords := parseExpenseReport("./passwords.txt");
	fmt.Printf("PASSWORDS: %v\n", passwords)

	for _, line := range passwords {
		if checkPasswordOldPolicy(line) {
			numValid1++
		}
	}

	fmt.Printf("Part 1 Num Valid: %v\n", numValid1)

	for _, line := range passwords {
		if checkPasswordNewPolicy(line) {
			numValid2++
		}
	}

	fmt.Printf("Part 2 Num Valid: %v\n", numValid2)
}
