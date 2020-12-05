package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const (
	BirthYear      = "byr"
	IssueYear      = "iyr"
	ExpirationYear = "eyr"
	Height         = "hgt"
	HairColor      = "hcl"
	EyeColor       = "ecl"
	PassportID     = "pid"
	CountryID      = "cid"
)

var (
	BirthYearRule      = []int{1920, 2002}
	IssueYearRule      = []int{2010, 2020}
	ExpirationYearRule = []int{2020, 2030}
	HeightRule         = map[string][]int{"cm": {150, 193}, "in": {59, 76}}
	HairColorRule, _   = regexp.Compile("[#][0-9a-f]{6}")
	EyeColorRule       = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	PassportIDRule, _  = regexp.Compile("[0-9]{9}")
)

func importPassports(name string) []string {
	text, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
	}

	lines := strings.Split(string(text), "\n\n")
	var passports []string

	for _, line := range lines {
		passports = append(passports, strings.ReplaceAll(line, "\n", " "))
	}

	return passports
}

func checkSubstrings(str string, subStrs []string) bool {
	for _, sub := range subStrs {
		if !strings.Contains(str, sub) {
			return false
		}
	}

	return true
}

func arrContainsElem(arr []string, elem string) bool {
	for _, item := range arr {
		if item == elem {
			return true
		}
	}

	return false
}

func checkFieldRules(passport string) bool {
	passportArr := strings.Split(passport, " ")

	for _, field := range passportArr {
		split := strings.Split(field, ":")
		key := split[0]
		value := split[1]

		switch key {
		case BirthYear:
			year, _ := strconv.Atoi(value)
			if year < BirthYearRule[0] || year > BirthYearRule[1] {
				return false
			}
		case IssueYear:
			year, _ := strconv.Atoi(value)
			if year < IssueYearRule[0] || year > IssueYearRule[1] {
				return false
			}
		case ExpirationYear:
			year, _ := strconv.Atoi(value)
			if year < ExpirationYearRule[0] || year > ExpirationYearRule[1] {
				return false
			}
		case Height:
			reg, _ := regexp.Compile("[^0-9]+")
			num, _ := strconv.Atoi(reg.ReplaceAllString(value, ""))
			cmStr := "cm"
			inStr := "in"

			if strings.Contains(value, cmStr) {
				if num < HeightRule[cmStr][0] || num > HeightRule[cmStr][1] {
					return false
				}
			} else if strings.Contains(value, inStr) {
				if num < HeightRule[inStr][0] || num > HeightRule[inStr][1] {
					return false
				}
			} else {
				return false
			}
		case HairColor:
			if !HairColorRule.MatchString(value) {
				return false
			}
		case EyeColor:
			if !arrContainsElem(EyeColorRule, value) {
				return false
			}
		case PassportID:
			if !PassportIDRule.MatchString(value) || len(value) > 9 {
				return false
			}
		}
	}

	return true
}

func validatePassports(passports []string, requiredFields []string) int {
	numValid := 0
	var filtered []string

	// filter out passports that don't have the required fields
	for _, passport := range passports {
		if checkSubstrings(passport, requiredFields) {
			filtered = append(filtered, passport)
		}
	}

	// check field rules
	for _, passport := range filtered {
		if checkFieldRules(passport) {
			numValid++
		}
	}

	return numValid
}

func main() {
	requiredFields := []string{BirthYear, IssueYear, ExpirationYear, Height, HairColor, EyeColor, PassportID}

	passports := importPassports("./passports.txt")

	numValid := validatePassports(passports, requiredFields)

	fmt.Printf("Num valid: %v\n", numValid)
}
