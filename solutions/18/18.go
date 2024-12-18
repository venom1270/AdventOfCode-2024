package solutions_day2

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	i int
	j int
}

func readInput(filename string) []Pos {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var positions []Pos
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ",")
		i, _ := strconv.Atoi(s[0])
		j, _ := strconv.Atoi(s[1])
		positions = append(positions, Pos{i, j})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return positions
}

type Item struct {
	pos      Pos
	priority int
	cost     int
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

func makeGrid(positions []Pos, N int, M int) [][]int {
	var grid [][]int

	for i := 0; i < N; i++ {
		var row []int
		for j := 0; j < M; j++ {
			row = append(row, 0)
		}
		grid = append(grid, row)
	}

	for _, pos := range positions {
		grid[pos.j][pos.i] = 2
	}

	return grid
}

func getHCost(position Pos, end Pos) int {
	return end.i - position.i + end.j - position.j
}

func part1(grid [][]int, start Pos, end Pos) int {
	pathLen := 0
	N := len(grid)
	M := len(grid[0])

	var memo map[Pos]int

	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		pos:      start,
		priority: 0,
		cost:     0,
		index:    0,
	}

	heap.Init(&pq)

	memo = map[Pos]int{}

	for pq.Len() > 0 {
		el := heap.Pop(&pq).(*Item)

		if el.pos.i < 0 || el.pos.i >= N || el.pos.j < 0 || el.pos.j >= M {
			// Out of bounds
			continue
		}
		if grid[el.pos.i][el.pos.j] == 2 {
			// Obstacle
			continue
		}

		if _, ok := memo[el.pos]; ok {
			continue
		}
		memo[el.pos] = 1

		if el.pos == end {
			return el.priority
		}

		// Move in all directions TODO: memo?
		pos1 := Pos{el.pos.i - 1, el.pos.j}
		pos2 := Pos{el.pos.i + 1, el.pos.j}
		pos3 := Pos{el.pos.i, el.pos.j - 1}
		pos4 := Pos{el.pos.i, el.pos.j + 1}
		heap.Push(&pq, &Item{
			pos:      pos1,
			cost:     el.cost + 1,
			priority: el.cost + 1 + getHCost(pos1, end),
		})
		heap.Push(&pq, &Item{
			pos:      pos2,
			cost:     el.cost + 1,
			priority: el.cost + 1 + getHCost(pos2, end),
		})
		heap.Push(&pq, &Item{
			pos:      pos3,
			cost:     el.cost + 1,
			priority: el.cost + 1 + getHCost(pos3, end),
		})
		heap.Push(&pq, &Item{
			pos:      pos4,
			cost:     el.cost + 1,
			priority: el.cost + 1 + getHCost(pos4, end),
		})

	}

	return pathLen
}

func part2(grid [][]int, start Pos, end Pos, bytes []Pos) Pos {
	var byteToAdd Pos
	for _, byteToAdd = range bytes {
		//fmt.Println("Adding: ", byteToAdd)
		grid[byteToAdd.j][byteToAdd.i] = 2
		pathCost := part1(grid, start, end)
		if pathCost == 0 {
			break
		}
	}
	return byteToAdd
}

func Solve() {
	positions := readInput("solutions/18/input.txt")
	//fmt.Println(positions)
	N := 71       // 7
	BYTES := 1024 //12
	grid := makeGrid(positions[:BYTES], N, N)
	//fmt.Println(grid)

	p1 := part1(grid, Pos{0, 0}, Pos{N - 1, N - 1})
	fmt.Println("Part 1: ", p1)

	p2 := part2(grid, Pos{0, 0}, Pos{N - 1, N - 1}, positions[BYTES:])
	fmt.Println("Part 2: ", p2)
}
