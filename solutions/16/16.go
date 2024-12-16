package solutions_day2

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Pos struct {
	i int
	j int
}

func readInput(filename string) ([][]byte, Pos, Pos) {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]byte
	j := 0
	scanner := bufio.NewScanner(file)
	var start, end Pos
	for scanner.Scan() {
		line := scanner.Text()
		var row []byte
		for i := range len(line) {
			row = append(row, line[i])
			if line[i] == 'S' {
				start = Pos{j, i}
			}
			if line[i] == 'E' {
				end = Pos{j, i}
			}
		}
		j++
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid, start, end
}

type Item struct {
	pos      Pos
	dir      Direction
	priority int
	path     []Pos
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

func getKey(el *Item) MapKey {
	return MapKey{el.pos.i, el.pos.j, el.dir}
}

type MapKey struct {
	i   int
	j   int
	dir Direction
}

func part1(grid [][]byte, start Pos, end Pos) int {
	score := 0

	pq := make(PriorityQueue, 1)

	var memo map[MapKey]int
	memo = map[MapKey]int{}

	pq[0] = &Item{
		pos:      start,
		dir:      Right,
		priority: 0,
		index:    0,
	}

	heap.Init(&pq)

	for pq.Len() > 0 {
		el := heap.Pop(&pq).(*Item)

		if el.pos == end {
			return el.priority
		}
		key := getKey(el)
		if _, ok := memo[key]; ok {
			if memo[key] <= el.priority {
				continue
			}
		}
		memo[key] = el.priority

		//fmt.Println(el.priority, el.pos, el.dir)

		// Move
		switch el.dir {
		case Up:
			if grid[el.pos.i-1][el.pos.j] != '#' {
				heap.Push(&pq, &Item{
					pos:      Pos{el.pos.i - 1, el.pos.j},
					dir:      el.dir,
					priority: el.priority + 1,
				})
			}
		case Down:
			if grid[el.pos.i+1][el.pos.j] != '#' {
				heap.Push(&pq, &Item{
					pos:      Pos{el.pos.i + 1, el.pos.j},
					dir:      el.dir,
					priority: el.priority + 1,
				})
			}
		case Left:
			if grid[el.pos.i][el.pos.j-1] != '#' {
				heap.Push(&pq, &Item{
					pos:      Pos{el.pos.i, el.pos.j - 1},
					dir:      el.dir,
					priority: el.priority + 1,
				})
			}
		case Right:
			if grid[el.pos.i][el.pos.j+1] != '#' {
				heap.Push(&pq, &Item{
					pos:      Pos{el.pos.i, el.pos.j + 1},
					dir:      el.dir,
					priority: el.priority + 1,
				})
			}
		}

		// Rotate
		heap.Push(&pq, &Item{
			pos:      el.pos,
			dir:      (el.dir + 1) % 4,
			priority: el.priority + 1000,
		})
		tmp := el.dir
		if el.dir == 0 {
			tmp = 4
		}
		heap.Push(&pq, &Item{
			pos:      el.pos,
			dir:      tmp - 1,
			priority: el.priority + 1000,
		})

	}

	return score
}

func copy(x []Pos) []Pos {
	var c []Pos
	for _, el := range x {
		c = append(c, el)
	}
	return c
}

func markPath(grid [][]byte, path []Pos) {
	for _, el := range path {
		//fmt.Println("Marking", el)
		grid[el.i][el.j] = 'O'
	}
}

func part2(grid [][]byte, start Pos, end Pos) int {
	score := 999999999

	pq := make(PriorityQueue, 1)

	var memo map[MapKey]int
	memo = map[MapKey]int{}
	var path []Pos
	path = append(path, start)

	pq[0] = &Item{
		pos:      start,
		dir:      Right,
		path:     path,
		priority: 0,
		index:    0,
	}

	heap.Init(&pq)

	for pq.Len() > 0 {
		el := heap.Pop(&pq).(*Item)

		if el.pos == end {
			//fmt.Println("Found path!")
			score = el.priority
			markPath(grid, el.path)
			continue
		}

		if el.priority > score {
			break
		}

		key := getKey(el)
		if _, ok := memo[key]; ok {
			if memo[key] < el.priority {
				continue
			}
		}
		memo[key] = el.priority

		//fmt.Println(el.priority, el.pos, el.dir)

		// Move
		switch el.dir {
		case Up:
			if grid[el.pos.i-1][el.pos.j] != '#' {
				newPos := Pos{el.pos.i - 1, el.pos.j}
				heap.Push(&pq, &Item{
					pos:      newPos,
					dir:      el.dir,
					path:     append(copy(el.path), newPos),
					priority: el.priority + 1,
				})
			}
		case Down:
			if grid[el.pos.i+1][el.pos.j] != '#' {
				newPos := Pos{el.pos.i + 1, el.pos.j}
				heap.Push(&pq, &Item{
					pos:      newPos,
					dir:      el.dir,
					path:     append(copy(el.path), newPos),
					priority: el.priority + 1,
				})
			}
		case Left:
			if grid[el.pos.i][el.pos.j-1] != '#' {
				newPos := Pos{el.pos.i, el.pos.j - 1}
				heap.Push(&pq, &Item{
					pos:      newPos,
					dir:      el.dir,
					path:     append(copy(el.path), newPos),
					priority: el.priority + 1,
				})
			}
		case Right:
			if grid[el.pos.i][el.pos.j+1] != '#' {
				newPos := Pos{el.pos.i, el.pos.j + 1}
				heap.Push(&pq, &Item{
					pos:      newPos,
					dir:      el.dir,
					path:     append(copy(el.path), newPos),
					priority: el.priority + 1,
				})
			}
		}

		// Rotate
		heap.Push(&pq, &Item{
			pos:      el.pos,
			path:     el.path,
			dir:      (el.dir + 1) % 4,
			priority: el.priority + 1000,
		})
		tmp := el.dir
		if el.dir == 0 {
			tmp = 4
		}
		heap.Push(&pq, &Item{
			pos:      el.pos,
			path:     el.path,
			dir:      tmp - 1,
			priority: el.priority + 1000,
		})

	}

	// Count path elements
	count := 0
	for _, row := range grid {
		for _, el := range row {
			if el == 'O' {
				count++
			}
		}
	}

	return count
}

func Solve() {
	grid, start, end := readInput("solutions/16/input.txt")
	//fmt.Println(grid, start, end)
	fmt.Println(part1(grid, start, end))
	fmt.Println(part2(grid, start, end))
}
