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

	var list []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		list = append(list, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return list
}

func check1(grid []string, i int, j int) int {
	if grid[i][j] != 'X' {
		return 0
	}

	N := len(grid)
	M := len(grid[0])
	xmasCount := 0

	// Check right
	if j+3 < M && grid[i][j+1] == 'M' && grid[i][j+2] == 'A' && grid[i][j+3] == 'S' {
		xmasCount++
	}
	// Check left
	if j-3 >= 0 && grid[i][j-1] == 'M' && grid[i][j-2] == 'A' && grid[i][j-3] == 'S' {
		xmasCount++
	}
	// Check up
	if i-3 >= 0 && grid[i-1][j] == 'M' && grid[i-2][j] == 'A' && grid[i-3][j] == 'S' {
		xmasCount++
	}
	// CHeck down
	if i+3 < N && grid[i+1][j] == 'M' && grid[i+2][j] == 'A' && grid[i+3][j] == 'S' {
		xmasCount++
	}
	// Check diagonals...
	if i+3 < N && j+3 < M && grid[i+1][j+1] == 'M' && grid[i+2][j+2] == 'A' && grid[i+3][j+3] == 'S' {
		xmasCount++
	}
	if i+3 < N && j-3 >= 0 && grid[i+1][j-1] == 'M' && grid[i+2][j-2] == 'A' && grid[i+3][j-3] == 'S' {
		xmasCount++
	}
	if i-3 >= 0 && j+3 < M && grid[i-1][j+1] == 'M' && grid[i-2][j+2] == 'A' && grid[i-3][j+3] == 'S' {
		xmasCount++
	}
	if i-3 >= 0 && j-3 >= 0 && grid[i-1][j-1] == 'M' && grid[i-2][j-2] == 'A' && grid[i-3][j-3] == 'S' {
		xmasCount++
	}

	return xmasCount
}

func check2(grid []string, i int, j int) int {
	if grid[i][j] != 'A' {
		return 0
	}

	N := len(grid)
	M := len(grid[0])

	if i < 1 || j < 1 || i+1 >= N || j+1 >= M {
		return 0
	}

	xmasCount := 0

	if (grid[i-1][j-1] == 'M' && grid[i+1][j+1] == 'S' || grid[i-1][j-1] == 'S' && grid[i+1][j+1] == 'M') &&
		(grid[i-1][j+1] == 'M' && grid[i+1][j-1] == 'S' || grid[i-1][j+1] == 'S' && grid[i+1][j-1] == 'M') {
		xmasCount++
	}

	return xmasCount
}

func Solve() {
	grid := readInput("solutions/4/input.txt")

	fmt.Println(grid)

	xmasCount1 := 0
	xmasCount2 := 0
	for i, row := range grid {
		for j, _ := range row {
			xmasCount1 += check1(grid, i, j)
			xmasCount2 += check2(grid, i, j)
		}
	}

	fmt.Println(xmasCount1)
	fmt.Println(xmasCount2)

}
