package solutions_day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(filename string) []string {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func moveGuard(grid []string, i int, j int, dir Direction) (int, int, Direction, bool) {
	var nextI, nextJ = i, j
	switch dir {
	case Up:
		nextI = i - 1
	case Right:
		nextJ = j + 1
	case Down:
		nextI = i + 1
	case Left:
		nextJ = j - 1
	}

	if nextI < 0 || nextJ < 0 || nextI >= len(grid) || nextJ >= len(grid[0]) {
		return -1, -1, -1, true
	}

	if grid[nextI][nextJ] == '#' {
		dir++
		if dir >= 4 {
			dir = 0 // Up
		}
		nextI, nextJ = i, j
	}
	return nextI, nextJ, dir, false

}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

type Position struct {
	i   int
	j   int
	dir Direction
}

func contains(memo []Position, x Position) bool {
	if memo == nil {
		return false
	}
	for _, pos := range memo {
		if pos.dir == x.dir && pos.i == x.i && pos.j == x.j {
			return true
		}
	}
	return false
}

func test(grid []string, guardI int, guardJ int) bool {
	var memo []Position = nil
	// Move the guard until failure, memorize visited locations (with respect to direction)
	guardDirection := Up
	for err := false; !err; guardI, guardJ, guardDirection, err = moveGuard(grid, guardI, guardJ, guardDirection) {
		pos := Position{guardI, guardJ, guardDirection}
		if contains(memo, pos) {
			return false
		} else {
			memo = append(memo, pos)
		}
	}

	return true
}

func Solve() {
	grid := readInput("solutions/6/input.txt")

	// Part 1
	// Find guard
	guardI := 0
	guardJ := 0
	var initialGuardI, initialGuardJ int
	for i, row := range grid {
		for j, el := range row {
			if el == '^' {
				guardI = i
				guardJ = j
				initialGuardI, initialGuardJ = i, j
				break
			}
		}
	}
	guardDirection := Up

	// Move the guard until failure
	grid[guardI] = replaceAtIndex(grid[guardI], 'X', guardJ)
	for err := false; !err; guardI, guardJ, guardDirection, err = moveGuard(grid, guardI, guardJ, guardDirection) {
		grid[guardI] = replaceAtIndex(grid[guardI], 'X', guardJ)
	}

	// CHeck how many X-es
	sum := 0
	for _, row := range grid {
		//fmt.Println(row)
		for _, el := range row {
			if el == 'X' {
				sum++
			}
		}
	}

	fmt.Println(sum)

	// Part 2 -- try putting obstructions on X-es and test if it's a loop
	sum = 0
	for i, row := range grid {
		//fmt.Println(row)
		for j, el := range row {
			if el == 'X' {
				tmp := grid[i]
				grid[i] = replaceAtIndex(grid[i], '#', j)
				// Test
				if !test(grid, initialGuardI, initialGuardJ) {
					sum++
				}
				// Swap back
				grid[i] = tmp
			}
		}
	}

	fmt.Println(sum)
}
