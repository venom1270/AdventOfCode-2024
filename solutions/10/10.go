package solutions_day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(filename string) [][]int {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var line string
	var grid [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		var l []int
		for _, el := range line {
			l = append(l, int(el-'0'))
		}
		grid = append(grid, l)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid
}

func findScore(grid [][]int, i int, j int, height int) int {

	if height == 9 {
		grid[i][j] = -1
		return 1
	}

	sum := 0
	if i-1 >= 0 && grid[i-1][j] == height+1 {
		sum += findScore(grid, i-1, j, height+1)
	}
	if i+1 < len(grid) && grid[i+1][j] == height+1 {
		sum += findScore(grid, i+1, j, height+1)
	}
	if j-1 >= 0 && grid[i][j-1] == height+1 {
		sum += findScore(grid, i, j-1, height+1)
	}
	if j+1 < len(grid[0]) && grid[i][j+1] == height+1 {
		sum += findScore(grid, i, j+1, height+1)
	}

	return sum
}

func findScore2(grid [][]int, i int, j int, height int) int {

	if height == 9 {
		return 1
	}

	sum := 0
	if i-1 >= 0 && grid[i-1][j] == height+1 {
		sum += findScore2(grid, i-1, j, height+1)
	}
	if i+1 < len(grid) && grid[i+1][j] == height+1 {
		sum += findScore2(grid, i+1, j, height+1)
	}
	if j-1 >= 0 && grid[i][j-1] == height+1 {
		sum += findScore2(grid, i, j-1, height+1)
	}
	if j+1 < len(grid[0]) && grid[i][j+1] == height+1 {
		sum += findScore2(grid, i, j+1, height+1)
	}

	return sum
}

func reset(grid [][]int) {
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == -1 {
				grid[i][j] = 9
			}
		}
	}
}

func Solve() {
	grid := readInput("solutions/10/input.txt")
	//fmt.Println(grid)

	// Part 1 and 2
	sum := 0
	sum2 := 0
	for i, row := range grid {
		for j, el := range row {
			if el == 0 {
				score := findScore(grid, i, j, 0)
				reset(grid)
				sum += score
				sum2 += findScore2(grid, i, j, 0)
			}
		}
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}
