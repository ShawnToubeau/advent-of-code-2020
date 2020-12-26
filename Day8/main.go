package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func loadInstructions(file string) []string {
	text, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
	}

	instructions := strings.Split(string(text), "\n")

	return instructions
}

func readInstruction(instr string) (string, string) {
	split := strings.Split(instr, " ")
	return split[0], split[1]
}

// run a instruction set until it either reaches the end or tries to run a instruction that was previously ran.
// returns the last accumulator value and whether the instruction set was able to make it to the end or not
func part1(instructions []string) (int, bool) {
	// current instruction step
	currStep := 0
	// accumulator
	acc := 0
	// visited array
	visited := make([]int, len(instructions))

	// loop while our current step hasn't reached the end yet
	for currStep < len(instructions) - 1 {
		currLine := instructions[currStep]
		// get the current instructions command and value
		instr, value := readInstruction(currLine)

		// check if the current instruction has been visited before
		if visited[currStep] == 1 {
			return acc, false
		} else {
			visited[currStep] = 1
		}

		// process the command and its value
		switch instr {
		case "acc":
			instrVal, _ := strconv.Atoi(value)
			acc = acc + instrVal
			currStep++
		case "jmp":
			instrVal, _ := strconv.Atoi(value)
			currStep = currStep + instrVal
		case "nop":
			currStep++
		}
	}

	return acc, true
}

// loop over every instruction and if the current instruction (instruction[i]) is a "nop" or "jmp" command we'll swap it with the
// opposite command. Then we'll run through the modified instruction set until we find the one that reaches the end
func part2(instructions []string) int {
	// loop for every instruction
	for i := range instructions {
		// make a copy of our instructions
		instrCopy := make([]string, len(instructions))
		copy(instrCopy, instructions)
		// get the command and value of the current instruction
		command, instrVal := readInstruction(instrCopy[i])
		// check if we should swap the current instruction
		switch command {
		case "acc":
			continue
		case "jmp":
			newInstr := fmt.Sprintf("%v %v", "nop", instrVal )
			instrCopy[i] = newInstr
		case "nop":
			newInstr := fmt.Sprintf("%v %v", "jmp", instrVal )
			instrCopy[i] = newInstr
		}
		// run the instruction set until it reaches an end (end of instruction set or runs a duplicate instruction)
		accVal, reachedEnd := part1(instrCopy)
		// if the run was able to reach the end we've found our answer
		if reachedEnd {
			return accVal
		}
	}

	return 0
}

func main() {
	instructions := loadInstructions("./instructions.txt")
	acc := 0

	acc,_ = part1(instructions)

	fmt.Printf("Par 1 accumulator value: %v\n", acc)

	acc = part2(instructions)

	fmt.Printf("Par 2 accumulator value: %v\n", acc)
}
