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

	var line string
	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid
}

func find(grid []string, newGrid [][]int, i int, j int, group byte, id int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return 0
	}

	if newGrid[i][j] != 0 {
		return 0
	}

	if grid[i][j] == group {
		sum := 1
		newGrid[i][j] = id
		// Optimize!!
		sum += find(grid, newGrid, i+1, j, group, id)
		sum += find(grid, newGrid, i-1, j, group, id)
		sum += find(grid, newGrid, i, j+1, group, id)
		sum += find(grid, newGrid, i, j-1, group, id)
		return sum
	}

	return 0
}

func findGroups(grid []string) ([][]int, []int) {
	// Init new array
	newGrid := make([][]int, len(grid))
	for i := range grid {
		newGrid[i] = make([]int, len(grid[0]))
	}

	id := 1
	var areaSizes []int
	for i := range grid {
		for j := range grid[i] {
			if newGrid[i][j] == 0 {
				areaSize := find(grid, newGrid, i, j, grid[i][j], id)
				areaSizes = append(areaSizes, areaSize)
				id++
			}
		}
	}

	return newGrid, areaSizes
}

var perimiterMap map[int]int

func calculatePerimiter(grid [][]int) {
	perimiterMap = map[int]int{}
	for i, row := range grid {
		for j, _ := range row {
			p := 0
			el := grid[i][j]
			if i-1 < 0 || grid[i-1][j] != el {
				p++
			}
			if j-1 < 0 || grid[i][j-1] != el {
				p++
			}
			if i+1 >= len(grid) || grid[i+1][j] != el {
				p++
			}
			if j+1 >= len(grid[0]) || grid[i][j+1] != el {
				p++
			}
			if m, ok := perimiterMap[el]; ok {
				perimiterMap[el] = m + p
			} else {
				perimiterMap[el] = p
			}
		}
	}
}

func calculateBulkPerimiter(grid [][]int) {
	perimiterMap = map[int]int{}
	N := len(grid)
	M := len(grid[0])
	for i, row := range grid {
		for j, _ := range row {
			p := 0
			el := grid[i][j]

			// Outer corners
			var testI, testJ bool
			testI = i-1 < 0 || grid[i-1][j] != el
			testJ = j+1 >= M || grid[i][j+1] != el
			if testI && testJ {
				p++
			}
			testI = i+1 >= N || grid[i+1][j] != el
			testJ = j+1 >= M || grid[i][j+1] != el
			if testI && testJ {
				p++
			}
			testI = i+1 >= N || grid[i+1][j] != el
			testJ = j-1 < 0 || grid[i][j-1] != el
			if testI && testJ {
				p++
			}
			testI = i-1 < 0 || grid[i-1][j] != el
			testJ = j-1 < 0 || grid[i][j-1] != el
			if testI && testJ {
				p++
			}

			// Inner corner
			if i-1 >= 0 && j+1 < M && grid[i-1][j] == el && grid[i][j+1] == el && grid[i-1][j+1] != el {
				p++
			}
			if i+1 < N && j+1 < M && grid[i][j+1] == el && grid[i+1][j] == el && grid[i+1][j+1] != el {
				p++
			}
			if i+1 < N && j-1 >= 0 && grid[i+1][j] == el && grid[i][j-1] == el && grid[i+1][j-1] != el {
				p++
			}
			if i-1 >= 0 && j-1 >= 0 && grid[i][j-1] == el && grid[i-1][j] == el && grid[i-1][j-1] != el {
				p++
			}

			if m, ok := perimiterMap[el]; ok {
				perimiterMap[el] = m + p
			} else {
				perimiterMap[el] = p
			}
		}
	}
}

func Solve() {
	grid := readInput("solutions/12/input.txt")
	//fmt.Println(grid)

	newGrid, areaSizes := findGroups(grid)
	//fmt.Println(newGrid)
	//fmt.Println(areaSizes)

	// Calculate price - part 1
	calculatePerimiter(newGrid)

	price := 0
	for k, v := range perimiterMap {
		price += (v * areaSizes[k-1])
	}

	fmt.Println(price)

	// Calculate price - part 2
	calculateBulkPerimiter(newGrid)

	price = 0
	for k, v := range perimiterMap {
		price += (v * areaSizes[k-1])
	}

	fmt.Println(price)

}
