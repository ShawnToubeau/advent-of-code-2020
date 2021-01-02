package main

import (
	"fmt"
	"github.com/shawntoubeau/advent-of-code-2020/utils"
	"sort"
)

const (
	deviceRating     = 3
	adaptorMaxRating = 3
)

// memoized map for tracking previously calculated recursive calls
var dpMap = make(map[int]int)

// Recursively finds the number of combination paths to the end of our adaptors list at the given adaptor index
func findValidAdaptor(idx int, adaptors []int) int {
	// if our current index is at the end of our list
	if idx == len(adaptors)-1 {
		return 1
	}
	// check if we've previously calculated the answer for the current index
	if dpMap[idx] != 0 {
		return dpMap[idx]
	}
	// # of combinations to the end from the current index
	combos := 0
	// branch out to the next 3 adaptors
	for j := 1; j <= 3; j++ {
		// check if we're 1) in bounds 2) the difference in adaptors is within our range (3 or less)
		if (idx+j) <= len(adaptors)-1 && adaptors[idx+j]-adaptors[idx] <= adaptorMaxRating {
			// recursively call our function if the next index (idx + j) is valid and sum the return to our combos
			combos += findValidAdaptor(idx+j, adaptors)
		}
	}
	// record the current combination total
	dpMap[idx] = combos
	return combos
}

func main() {
	adaptors := utils.ConvertToNumArr(utils.ReadFile("./adaptors.txt"))
	currJoltage := 0
	joltDiffMap := make(map[int]int)
	// account for the starting joltage (0)
	adaptors = append(adaptors, 0)
	// sort our adaptors in ascending order
	sort.Ints(adaptors)
	// account for our device's joltage rating (device rating plus the highest rating of our adaptors)
	adaptors = append(adaptors, adaptors[len(adaptors)-1]+deviceRating)

	for _, adaptor := range adaptors {
		if adaptor-adaptorMaxRating <= currJoltage {
			joltDiffMap[utils.Abs(adaptor-currJoltage)]++
			currJoltage = adaptor
		} else {
			fmt.Printf("Adaptor not compatible\n")
		}
	}

	fmt.Printf("Part 1 Answer: %v\n", joltDiffMap[1]*joltDiffMap[3])

	part2Ans := findValidAdaptor(0, adaptors)

	fmt.Printf("Part 2 Answer: %v\n", part2Ans)
}
