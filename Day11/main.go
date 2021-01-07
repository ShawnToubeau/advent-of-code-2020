package main

import (
	"fmt"
	"github.com/shawntoubeau/advent-of-code-2020/utils"
	"strings"
)

const (
	floor = "."
	empty = "L"
	taken = "#"
)

type Seating struct {
	seats [][]string
}

func (s *Seating) getAdjacentSeats(row int, col int) []string {
	var adjSeats []string

	// top left corner
	if row == 0 && col == 0 {
		adjSeats = append(adjSeats,
			s.seats[row+1][col],
			s.seats[row+1][col+1],
			s.seats[row][col+1],
		)
	}
	// top right corner
	if row == 0 && col == len(s.seats[0])-1 {
		adjSeats = append(adjSeats,
			s.seats[row+1][col],
			s.seats[row+1][col-1],
			s.seats[row][col-1],
		)
	}
	// bottom left corner
	if row == len(s.seats)-1 && col == 0 {
		adjSeats = append(adjSeats,
			s.seats[row][col+1],
			s.seats[row-1][col+1],
			s.seats[row-1][col],
		)
	}
	// bottom right corner
	if row == len(s.seats)-1 && col == len(s.seats[0])-1 {
		adjSeats = append(adjSeats,
			s.seats[row-1][col],
			s.seats[row-1][col-1],
			s.seats[row][col-1],
		)
	}
	// top edge
	if row == 0 && col > 0 && col < len(s.seats[0])-1 {
		//fmt.Printf("Test: %v\n", s.seats[row][col+1])
		adjSeats = append(adjSeats,
			s.seats[row][col-1],
			s.seats[row+1][col-1],
			s.seats[row+1][col],
			s.seats[row+1][col+1],
			s.seats[row][col+1],
		)
	}
	// right edge
	if row > 0 && row < len(s.seats)-1 && col == len(s.seats[0])-1 {
		adjSeats = append(adjSeats,
			s.seats[row-1][col],
			s.seats[row-1][col-1],
			s.seats[row][col-1],
			s.seats[row+1][col-1],
			s.seats[row+1][col],
		)
	}
	// bottom edge
	if row == len(s.seats)-1 && col > 0 && col < len(s.seats[0])-1 {
		adjSeats = append(adjSeats,
			s.seats[row][col-1],
			s.seats[row-1][col-1],
			s.seats[row-1][col],
			s.seats[row-1][col+1],
			s.seats[row][col+1],
		)
	}
	// left edge
	if row > 0 && row < len(s.seats)-1 && col == 0 {
		adjSeats = append(adjSeats,
			s.seats[row-1][col],
			s.seats[row-1][col+1],
			s.seats[row][col+1],
			s.seats[row+1][col+1],
			s.seats[row+1][col],
		)
	}
	// everything else within bounds
	if len(adjSeats) == 0 {
		adjSeats = append(adjSeats,
			s.seats[row-1][col-1],
			s.seats[row-1][col],
			s.seats[row-1][col+1],
			s.seats[row][col+1],
			s.seats[row+1][col+1],
			s.seats[row+1][col],
			s.seats[row+1][col-1],
			s.seats[row][col-1],
		)
	}

	return adjSeats
}

func (s *Seating) checkRules(row int, col int) string {
	currSeat := s.seats[row][col]

	switch currSeat {
	case floor:
		return currSeat
	case empty:
		adjSeats := s.getAdjacentSeats(row, col)

		// if all adjacent seats are empty
		if !utils.ArrContainsElem(adjSeats, taken) {
			return taken
		}
		// else, no change
		return currSeat
	case taken:
		adjSeats := s.getAdjacentSeats(row, col)

		// if the number of adjacent seats that are occupied is four or more
		if utils.ArrNumOfOccurrences(adjSeats, taken) >= 4 {
			return empty
		}
		// else, no change
		return currSeat
	default:
		return currSeat
	}
}

func (s *Seating) getNumOccupied() int {
	count := 0

	for _, row := range s.seats {
		for _, col := range row {
			if col == taken {
				count++
			}
		}
	}

	return count
}

func (s *Seating) printSeats() {
	for _, row := range s.seats {
		fmt.Printf("%v\n", row)
	}
}

func (s *Seating) simulate() int {
	hasChanges := true
	iterations := 0

	for hasChanges {
		iterations++

		if iterations == 2 {
			s.printSeats()
		}

		numChanges := 0

		// create a copy of the current seating chart to store all the changes which will get applied at the end
		duplicate := make([][]string, len(s.seats))
		for i := range s.seats {
			duplicate[i] = make([]string, len(s.seats[i]))
			copy(duplicate[i], s.seats[i])
		}

		for i, row := range s.seats {
			for j, col := range row {
				updatedSeat := s.checkRules(i, j)
				duplicate[i][j] = updatedSeat

				if updatedSeat != col {
					numChanges++
				}
			}
		}

		s.seats = duplicate

		if numChanges == 0 {
			hasChanges = false
		}
	}

	return s.getNumOccupied()
}

func loadSeating(file string) Seating {
	seatRows := utils.ReadFile(file)
	var seating Seating

	for _, row := range seatRows {
		rowArr := strings.Split(row, "")
		seating.seats = append(seating.seats, rowArr)
	}

	return seating
}

func main() {
	seats := loadSeating("./seats.txt")

	numOccupied := seats.simulate()

	fmt.Printf("%v\n", numOccupied)
}
