package solutions_day2

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
)

func readInput(filename string) ([][]rune, Pos, Pos) {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	i := 0
	var start, end Pos
	for scanner.Scan() {
		line := scanner.Text()
		var row []rune
		for j, el := range line {
			row = append(row, el)
			if el == 'S' {
				start = Pos{i, j}
			} else if el == 'E' {
				end = Pos{i, j}
			}
		}
		grid = append(grid, row)
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid, start, end
}

type Pos struct {
	i int
	j int
}

type Item struct {
	pos      Pos
	priority int
	path     []Pos
	//hasCheated bool
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

var memo map[Pos]int

func getAtPos(grid [][]rune, pos Pos) rune {
	return grid[pos.i][pos.j]
}

func posValid(pos Pos, n int, m int) bool {
	if pos.i < 0 || pos.j < 0 || pos.i >= n || pos.j >= m {
		return false
	} else {
		return true
	}
}

func findShortestPath(grid [][]rune, start Pos, end Pos) (int, []Pos) {

	N := len(grid)
	M := len(grid[0])

	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		pos: start,
		//hasCheated: false,
		priority: 0,
		index:    0,
		path:     []Pos{},
	}

	heap.Init(&pq)

	memo = map[Pos]int{}

	for pq.Len() > 0 {
		el := heap.Pop(&pq).(*Item)

		if !posValid(el.pos, N, M) {
			// Out of bounds
			continue
		}
		if getAtPos(grid, el.pos) == '#' {
			// Obstacle
			continue
		}

		el.path = append(copy(el.path), el.pos)

		if el.pos.i == end.i && el.pos.j == end.j {
			return el.priority, el.path
		}

		if _, ok := memo[el.pos]; ok {
			continue
		}
		memo[el.pos] = 1

		// Move in all directions
		pos1 := Pos{el.pos.i - 1, el.pos.j}
		pos2 := Pos{el.pos.i + 1, el.pos.j}
		pos3 := Pos{el.pos.i, el.pos.j - 1}
		pos4 := Pos{el.pos.i, el.pos.j + 1}
		heap.Push(&pq, &Item{
			pos:      pos1,
			priority: el.priority + 1,
			path:     el.path,
		})
		heap.Push(&pq, &Item{
			pos:      pos2,
			priority: el.priority + 1,
			path:     el.path,
		})
		heap.Push(&pq, &Item{
			pos:      pos3,
			priority: el.priority + 1,
			path:     el.path,
		})
		heap.Push(&pq, &Item{
			pos:      pos4,
			priority: el.priority + 1,
			path:     el.path,
		})

	}

	return -1, []Pos{}
}

var cheatPaths map[int]int

func copy(x []Pos) []Pos {
	var c []Pos
	for _, i := range x {
		c = append(c, i)
	}
	return c
}

func checkIfShortcut(path []Pos, pos Pos, i int) bool {
	for j := range i {
		if path[j] == pos {
			return false
		}
	}
	return true
}

func getSaved(path []Pos, pos Pos, i int) int {

	saved := -999999

	for j := range len(path) {
		if path[j] == pos {
			saved = j - i - 2
			break
		}
	}

	if cp, b := cheatPaths[saved]; b {
		cheatPaths[saved] = cp + 1
	} else {
		cheatPaths[saved] = 1
	}

	return -1
}

func checkCheatsPart1(grid [][]rune, path []Pos) {
	// Mark path
	for _, p := range path {
		grid[p.i][p.j] = 'X'
	}

	cheatPaths = map[int]int{}

	N := len(grid)
	M := len(grid[0])
	count := 0
	for i, p := range path[:len(path)-1] {

		pos1 := Pos{p.i - 1, p.j}
		pos2 := Pos{p.i + 1, p.j}
		pos3 := Pos{p.i, p.j - 1}
		pos4 := Pos{p.i, p.j + 1}

		pos1C := Pos{p.i - 2, p.j}
		pos2C := Pos{p.i + 2, p.j}
		pos3C := Pos{p.i, p.j - 2}
		pos4C := Pos{p.i, p.j + 2}

		if posValid(pos1, N, M) && getAtPos(grid, pos1) == '#' && posValid(pos1C, N, M) && getAtPos(grid, pos1C) == 'X' && checkIfShortcut(path, pos1C, i) {
			getSaved(path, pos1C, i)
			count++
		}
		if posValid(pos2, N, M) && getAtPos(grid, pos2) == '#' && posValid(pos2C, N, M) && getAtPos(grid, pos2C) == 'X' && checkIfShortcut(path, pos2C, i) {
			getSaved(path, pos2C, i)
			count++
		}
		if posValid(pos3, N, M) && getAtPos(grid, pos3) == '#' && posValid(pos3C, N, M) && getAtPos(grid, pos3C) == 'X' && checkIfShortcut(path, pos3C, i) {
			getSaved(path, pos3C, i)
			count++
		}
		if posValid(pos4, N, M) && getAtPos(grid, pos4) == '#' && posValid(pos4C, N, M) && getAtPos(grid, pos4C) == 'X' && checkIfShortcut(path, pos4C, i) {
			getSaved(path, pos4C, i)
			count++
		}

	}

	//fmt.Println("COunt = ", count)
	//fmt.Println(cheatPaths)

	// Get all that save at least 100 seconds
	countCheats := 0
	for saving, count := range cheatPaths {
		if saving >= 100 {
			countCheats += count
		}
	}

	fmt.Println("Part 1: ", countCheats)
}

func distance(p1 Pos, p2 Pos) int {
	x := p1.i - p2.i
	y := p1.j - p2.j
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	return x + y
}

func checkCheatsPart2(grid [][]rune, path []Pos) {
	// Mark path
	for _, p := range path {
		grid[p.i][p.j] = 'X'
	}

	cheatPaths = map[int]int{}

	for i, p := range path[:len(path)-1] {

		// Cheat at poition p
		// Count all path nodes with index>i if they are withnin radious
		for j, pj := range path[i+2:] {
			index := j + i + 2
			d := distance(p, pj)
			if d <= 20 && d < index-i {
				saved := j + 2 - d //- 2 + d
				if cp, b := cheatPaths[saved]; b {
					cheatPaths[saved] = cp + 1
				} else {
					cheatPaths[saved] = 1
				}
			}
		}
	}

	//fmt.Println("COunt = ", count)
	//fmt.Println(cheatPaths)

	// Get all that save at least 100 seconds
	countCheats := 0
	for saving, count := range cheatPaths {
		if saving >= 100 {
			countCheats += count
		}
	}

	fmt.Println("Part 2: ", countCheats)
}

func Solve() {
	grid, start, end := readInput("solutions/20/input.txt")
	//fmt.Println(grid, start, end)
	shortestPath, path := findShortestPath(grid, start, end)
	fmt.Println("Shortest path length: ", shortestPath)

	checkCheatsPart1(grid, path)
	checkCheatsPart2(grid, path)
}
