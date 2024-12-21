package solutions_day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput(filename string) []string {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var codes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		codes = append(codes, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return codes
}

type Pos struct {
	i, j int
}

func getGrid() [][]rune {
	grid := [][]rune{
		{'7', '8', '9'},
		{'4', '5', '6'},
		{'1', '2', '3'},
		{'X', '0', 'A'},
	}
	return grid
}

func distance(x, y Pos) int {
	dx := x.i - y.i
	dy := x.j - y.j
	if dx < 0 {
		dx *= -1
	}
	if dy < 0 {
		dy *= -1
	}
	return dx + dy
}

type Path struct {
	//pos Pos
	path  []Pos
	moves string
}

func copyPath(x []Pos) []Pos {
	var c []Pos
	for _, el := range x {
		c = append(c, el)
	}
	return c
}

type ShortestPathMapKey struct {
	start Pos
	end   Pos
}

var shortestPathMemo map[ShortestPathMapKey][]Path

func getShortestPaths(grid [][]rune, start Pos, end Pos) []Path {
	queue := make([]Path, 0)
	queue = append(queue, Path{[]Pos{start}, ""})

	shortestLength := distance(start, end) + 1

	var paths []Path

	if shortestPathMemo == nil {
		shortestPathMemo = map[ShortestPathMapKey][]Path{}
	} else {
		if m, ok := shortestPathMemo[ShortestPathMapKey{start, end}]; ok {
			return m
		}
	}

	for len(queue) > 0 {
		el := queue[0]
		queue = queue[1:]

		if len(el.path) > shortestLength {
			break
		}

		pos := el.path[len(el.path)-1]
		if grid[pos.i][pos.j] == 'X' {
			continue
		}

		if pos == end {
			paths = append(paths, el)
		}

		if end.i < pos.i {
			// Go up
			newPath := copyPath(el.path)
			newPath = append(newPath, Pos{pos.i - 1, pos.j})
			newEl := Path{newPath, el.moves + "^"}
			queue = append(queue, newEl)
		} else if end.i > pos.i {
			// Go down
			newPath := copyPath(el.path)
			newPath = append(newPath, Pos{pos.i + 1, pos.j})
			newEl := Path{newPath, el.moves + "v"}
			queue = append(queue, newEl)
		}

		if end.j < pos.j {
			// Go left
			newPath := copyPath(el.path)
			newPath = append(newPath, Pos{pos.i, pos.j - 1})
			newEl := Path{newPath, el.moves + "<"}
			queue = append(queue, newEl)
		} else if end.j > pos.j {
			// Go right
			newPath := copyPath(el.path)
			newPath = append(newPath, Pos{pos.i, pos.j + 1})
			newEl := Path{newPath, el.moves + ">"}
			queue = append(queue, newEl)
		}

	}

	shortestPathMemo[ShortestPathMapKey{start, end}] = paths

	//fmt.Println(paths)
	//fmt.Println(len(paths))

	return paths
}

func getShortestPathsKeypad(start rune, end rune) []string {
	if start == end {
		return []string{""}
	}
	if start == '^' {
		switch end {
		case '<':
			return []string{"v<"}
		case '>':
			return []string{"v>", ">v"}
			//return []string{"v>"}
		case 'v':
			return []string{"v"}
		case 'A':
			return []string{">"}
		}
	}
	if start == 'v' {
		switch end {
		case '<':
			return []string{"<"}
		case '>':
			return []string{">"}
		case '^':
			return []string{"^"}
		case 'A':
			return []string{"^>", ">^"}
			//return []string{"^>"}
		}
	}
	if start == '<' {
		switch end {
		case 'v':
			return []string{">"}
		case '>':
			return []string{">>"}
		case '^':
			return []string{">^"}
		case 'A':
			//return []string{">^>", ">>^"}
			return []string{">>^"}
		}
	}
	if start == '>' {
		switch end {
		case 'v':
			return []string{"<"}
		case '<':
			return []string{"<<"}
		case '^':
			return []string{"<^", "^<"}
			//return []string{"<^"}
		case 'A':
			return []string{"^"}
		}
	}
	if start == 'A' {
		switch end {
		case 'v':
			return []string{"<v", "v<"}
			//return []string{"<v"}
		case '>':
			return []string{"v"}
		case '^':
			return []string{"<"}
		case '<':
			//return []string{"v<<", "<v<"}
			return []string{"v<<"}
		}
	}
	fmt.Println(start, end)
	panic("wrong key!!")
}

func getDigitPos(grid [][]rune, digit rune) Pos {
	for i, row := range grid {
		for j, el := range row {
			if el == digit {
				return Pos{i, j}
			}
		}
	}
	panic("no digit!")
}

func getCodeValue(code string) int {
	val := 0
	for i, _ := range code {
		n, err := strconv.Atoi(code[i : i+1])
		if err == nil {
			val *= 10
			val += n
		}
	}
	return val
}

func getSequence(grid [][]rune, codes []string, numRobots int) {

	var shortestPathLengths []int

	paths := []string{""}
	var newPaths []string

	start := Pos{3, 2}
	for _, code := range codes {
		for _, digit := range code {
			end := getDigitPos(grid, digit)

			shortestPaths := getShortestPaths(grid, start, end)
			for _, p := range paths {
				for _, sp := range shortestPaths {
					np := p + sp.moves + "A"
					newPaths = append(newPaths, np)
				}
			}

			paths = newPaths
			newPaths = nil

			start = end
		}

		shortest := 999999999999999999
		memo = map[MemoKey]int{}
		for _, p := range paths {
			path := dfs(p, 'A', 0, numRobots)
			if path < shortest {
				shortest = path
			}
		}
		shortestPathLengths = append(shortestPathLengths, shortest)

		paths = []string{""}
		newPaths = nil
		start = Pos{3, 2}
	}

	//fmt.Println(shortestPathLengths)

	complexity := 0
	for i := range shortestPathLengths {
		complexity += shortestPathLengths[i] * getCodeValue(codes[i])
	}

	fmt.Println("Complexity: ", complexity)
}

type MemoKey struct {
	path  string
	depth int
	start rune
}

var memo map[MemoKey]int

func dfs(path string, start rune, depth int, maxDepth int) int {
	key := MemoKey{path, depth, start}

	if el, ok := memo[key]; ok {
		return el
	}

	if depth >= maxDepth {
		return len(path)
	}

	// Check all moves in path
	shortestPathLen := 0
	for _, p := range path {
		shortestPaths := getShortestPathsKeypad(start, p)
		shortestPathTmp := 999999999999999999
		for _, sp := range shortestPaths {
			tmp := dfs(sp+"A", 'A', depth+1, maxDepth)
			if tmp < shortestPathTmp {
				shortestPathTmp = tmp
			}
		}
		start = p
		shortestPathLen += shortestPathTmp
	}

	memo[key] = shortestPathLen
	return shortestPathLen
}

func Solve() {
	codes := readInput("solutions/21/input.txt")
	fmt.Println(codes)

	grid := getGrid()
	fmt.Println("PART 1")
	getSequence(grid, codes, 2)
	fmt.Println("PART 2")
	getSequence(grid, codes, 25)
}
