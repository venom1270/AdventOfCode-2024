package solutions_day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func readInput(filename string) ([][]byte, []Direction) {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 3 {
			break
		}

		var row []byte
		for i := range len(line) {
			row = append(row, line[i])
		}
		grid = append(grid, row)
	}

	var dirs []Direction
	for scanner.Scan() {
		line := scanner.Text()
		for i := range len(line) {
			var d Direction
			switch line[i] {
			case '<':
				d = Left
			case '>':
				d = Right
			case '^':
				d = Up
			case 'v':
				d = Down
			default:
				fmt.Println("Parsing error or newline")
			}
			dirs = append(dirs, d)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid, dirs
}

func getRobotPos(grid [][]byte) (int, int) {
	for i, el := range grid {
		for j, b := range el {
			if b == '@' {
				return i, j
			}
		}
	}
	return -1, -1
}

func moveRobot(grid [][]byte, dir Direction, rI int, rJ int) (int, int) {
	if dir == Up {
		i := rI - 1
		for grid[i][rJ] == 'O' {
			i--
		}
		if grid[i][rJ] == '#' {
			// Invalid move
			//fmt.Println("Invalid move", dir)
			return rI, rJ
		}
		// Move boxes and robot
		grid[i][rJ] = 'O'
		grid[rI-1][rJ] = '@'
		grid[rI][rJ] = '.'
		rI = rI - 1
	} else if dir == Down {
		i := rI + 1
		for grid[i][rJ] == 'O' {
			i++
		}
		if grid[i][rJ] == '#' {
			// Invalid move
			//fmt.Println("Invalid move", dir)
			return rI, rJ
		}
		// Move boxes and robot
		grid[i][rJ] = 'O'
		grid[rI+1][rJ] = '@'
		grid[rI][rJ] = '.'
		rI = rI + 1
	} else if dir == Left {
		j := rJ - 1
		for grid[rI][j] == 'O' {
			j--
		}
		if grid[rI][j] == '#' {
			// Invalid move
			//fmt.Println("Invalid move", dir)
			return rI, rJ
		}
		// Move boxes and robot
		grid[rI][j] = 'O'
		grid[rI][rJ-1] = '@'
		grid[rI][rJ] = '.'
		rJ = rJ - 1
	} else if dir == Right {
		j := rJ + 1
		for grid[rI][j] == 'O' {
			j++
		}
		if grid[rI][j] == '#' {
			// Invalid move
			//fmt.Println("Invalid move", dir)
			return rI, rJ
		}
		// Move boxes and robot
		grid[rI][j] = 'O'
		grid[rI][rJ+1] = '@'
		grid[rI][rJ] = '.'
		rJ = rJ + 1
	}

	return rI, rJ
}

func moveRobotPart2(grid [][]byte, dir Direction, rI int, rJ int) (int, int) {
	if dir == Up {
		i := rI - 1

		if grid[i][rJ] == '#' {
			//fmt.Println("Invalid move")
			return rI, rJ
		}

		var boxesToMove [][]int
		if grid[i][rJ] == '[' {
			boxesToMove = append(boxesToMove, []int{i, rJ})
			//boxesToMove = append(boxesToMove, []int{i, rJ + 1})
		} else if grid[i][rJ] == ']' {
			boxesToMove = append(boxesToMove, []int{i, rJ - 1})
			//boxesToMove = append(boxesToMove, []int{i, rJ})
		}

		currBox := 0
		for currBox < len(boxesToMove) {
			// Check each box
			box := boxesToMove[currBox]
			bi, bj := box[0], box[1]

			if grid[bi-1][bj] == '#' || grid[bi-1][bj+1] == '#' {
				//fmt.Println("Invalid move!", dir)
				return rI, rJ
			}

			//fmt.Println(bi, bj)

			if grid[bi-1][bj] == '[' {
				boxesToMove = append(boxesToMove, []int{bi - 1, bj})
			}
			if grid[bi-1][bj] == ']' {
				boxesToMove = append(boxesToMove, []int{bi - 1, bj - 1})
			}
			if grid[bi-1][bj+1] == '[' {
				boxesToMove = append(boxesToMove, []int{bi - 1, bj + 1})
			}

			currBox++
		}

		//fmt.Println("Boxes to move: ", len(boxesToMove))

		// Move boxes
		for i = len(boxesToMove) - 1; i >= 0; i-- {
			box := boxesToMove[i]
			bi, bj := box[0], box[1]
			grid[bi-1][bj] = '['
			grid[bi-1][bj+1] = ']'
			grid[bi][bj] = '.'
			grid[bi][bj+1] = '.'
		}

		// Move robot
		grid[rI-1][rJ] = '@'
		grid[rI][rJ] = '.'
		rI = rI - 1
	} else if dir == Down {
		i := rI + 1

		if grid[i][rJ] == '#' {
			//fmt.Println("Invalid move")
			return rI, rJ
		}

		var boxesToMove [][]int
		if grid[i][rJ] == '[' {
			boxesToMove = append(boxesToMove, []int{i, rJ})
			//boxesToMove = append(boxesToMove, []int{i, rJ + 1})
		} else if grid[i][rJ] == ']' {
			boxesToMove = append(boxesToMove, []int{i, rJ - 1})
			//boxesToMove = append(boxesToMove, []int{i, rJ})
		}

		currBox := 0
		for currBox < len(boxesToMove) {
			// Check each box
			box := boxesToMove[currBox]
			bi, bj := box[0], box[1]

			if grid[bi+1][bj] == '#' || grid[bi+1][bj+1] == '#' {
				//fmt.Println("Invalid move!", dir)
				return rI, rJ
			}

			if grid[bi+1][bj] == '[' {
				boxesToMove = append(boxesToMove, []int{bi + 1, bj})
			}
			if grid[bi+1][bj] == ']' {
				boxesToMove = append(boxesToMove, []int{bi + 1, bj - 1})
			}
			if grid[bi+1][bj+1] == '[' {
				boxesToMove = append(boxesToMove, []int{bi + 1, bj + 1})
			}

			currBox++
		}

		// Move boxes
		for i = len(boxesToMove) - 1; i >= 0; i-- {
			box := boxesToMove[i]
			bi, bj := box[0], box[1]
			grid[bi+1][bj] = '['
			grid[bi+1][bj+1] = ']'
			grid[bi][bj] = '.'
			grid[bi][bj+1] = '.'
		}

		// Move robot
		grid[rI+1][rJ] = '@'
		grid[rI][rJ] = '.'
		rI = rI + 1
	} else if dir == Left {
		j := rJ - 1
		for grid[rI][j] == ']' {
			j -= 2
		}
		if grid[rI][j] == '#' {
			// Invalid move
			//fmt.Println("Invalid move", dir)
			return rI, rJ
		}
		// Move boxes and robot
		grid[rI][j] = '['
		j++
		for j < rJ {
			if grid[rI][j] == '[' {
				grid[rI][j] = ']'
			} else if grid[rI][j] == ']' {
				grid[rI][j] = '['
			}
			j++
		}
		grid[rI][rJ-1] = '@'
		grid[rI][rJ] = '.'
		rJ = rJ - 1
	} else if dir == Right {
		j := rJ + 1
		for grid[rI][j] == '[' {
			j += 2
		}
		if grid[rI][j] == '#' {
			// Invalid move
			//fmt.Println("Invalid move", dir)
			return rI, rJ
		}
		// Move boxes and robot
		grid[rI][j] = ']'
		j--
		for j > rJ {
			if grid[rI][j] == '[' {
				grid[rI][j] = ']'
			} else if grid[rI][j] == ']' {
				grid[rI][j] = '['
			}
			j--
		}
		grid[rI][rJ+1] = '@'
		grid[rI][rJ] = '.'
		rJ = rJ + 1
	}

	return rI, rJ
}

func simulateRobot(grid [][]byte, dirs []Direction, part1 bool) {
	rI, rJ := getRobotPos(grid)

	for _, d := range dirs {
		//fmt.Println("Moving ", d)
		if part1 {
			rI, rJ = moveRobot(grid, d, rI, rJ)
		} else {
			rI, rJ = moveRobotPart2(grid, d, rI, rJ)
		}

	}
}

func getGpsSum(grid [][]byte) int {
	sum := 0
	for i, row := range grid {
		for j, el := range row {
			if el == 'O' || el == '[' {
				gps := (100*i + j)
				//fmt.Println("GPS ", i, j, gps)
				sum += gps
			}
		}
	}
	return sum
}

func getSecondWarehouse(grid [][]byte) [][]byte {
	var grid2 [][]byte
	for _, row := range grid {
		var r []byte
		for j := range row {
			switch row[j] {
			case '.':
				r = append(r, '.')
				r = append(r, '.')
			case '#':
				r = append(r, '#')
				r = append(r, '#')
			case 'O':
				r = append(r, '[')
				r = append(r, ']')
			case '@':
				r = append(r, '@')
				r = append(r, '.')
			default:
				fmt.Println("Error creating second warehouse")
			}
		}
		grid2 = append(grid2, r)
	}
	return grid2
}

func Solve() {
	warehouse, dirs := readInput("solutions/15/input.txt")
	warehouse2 := getSecondWarehouse(warehouse)

	// Part 1
	simulateRobot(warehouse, dirs, true)
	fmt.Println(getGpsSum(warehouse))

	// Part 2
	//fmt.Println(warehouse2)
	simulateRobot(warehouse2, dirs, false)
	//fmt.Println(warehouse2)
	fmt.Println(getGpsSum(warehouse2))

}
