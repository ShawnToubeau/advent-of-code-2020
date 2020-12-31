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

func findValidAdaptor(idx int, adaptors []int, dpMap map[int]int) int {
	fmt.Printf("IDX %v\n", idx)
	if idx == len(adaptors)-1 {
		return 1
	}
	if dpMap[idx] != 0 {
		return dpMap[idx]
	}
	combos := 0
	for j := idx + 1; j < len(adaptors); j++ {
		fmt.Printf("J %v\n", j)
		if adaptors[j]-adaptors[idx] <= adaptorMaxRating {
			combos = combos + findValidAdaptor(j, adaptors, dpMap)
		}
	}
	dpMap[idx] = combos
	return combos
}

func main() {
	adaptors := utils.ConvertToNumArr(utils.ReadFile("./test.txt"))
	currJoltage := 0
	joltDiffMap := make(map[int]int)
	// sort our adaptors in ascending order
	sort.Ints(adaptors)
	// account for our device's joltage rating (max rating plus the highest rating of our adaptors)
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

	dpMap := make(map[int]int)
	part2Ans := findValidAdaptor(0, adaptors, dpMap)

	fmt.Printf("Part 2 Answer: %v\n", part2Ans)
}
