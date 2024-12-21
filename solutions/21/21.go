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
		[]rune{'7', '8', '9'},
		[]rune{'4', '5', '6'},
		[]rune{'1', '2', '3'},
		[]rune{'X', '0', 'A'},
	}
	return grid
}

func getKeypad() [][]rune {
	keypad := [][]rune{
		[]rune{'X', 'U', 'A'},
		[]rune{'L', 'D', 'R'},
	}
	return keypad
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
			//return []string{"v>", ">v"}
			return []string{"v>"}
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
			//return []string{"^>", ">^"}
			return []string{"^>"}
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
			//return []string{"<^", "^<"}
			return []string{"<^"}
		case 'A':
			return []string{"^"}
		}
	}
	if start == 'A' {
		switch end {
		case 'v':
			//return []string{"<v", "v<"}
			return []string{"<v"}
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

func filterShortestPaths(paths []string) []string {
	shortestPaths := []string{}
	// Find shortest
	shortestLen := 9999999999
	for _, p := range paths {
		if len(p) < shortestLen {
			shortestLen = len(p)
		}
	}
	// Only take shortest
	for _, p := range paths {
		if len(p) == shortestLen {
			shortestPaths = append(shortestPaths, p)
		}
	}
	return shortestPaths
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

func getScore(path string) int {
	score := 0
	for i := range len(path) - 1 {
		if path[i] == path[i+1] {
			score++
		}
	}
	return score
}

func filterBestScore(paths []string) []string {
	bestPaths := []string{}
	// Find best score
	bestScore := 0
	for _, p := range paths {
		score := getScore(p)
		if score > bestScore {
			bestScore = score
		}
	}
	// Only take best paths
	for _, p := range paths {
		if getScore(p) == bestScore {
			bestPaths = append(bestPaths, p)
		}
	}
	return bestPaths
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

		//fmt.Println(paths)

		// Go through all the robots...
		for i := range numRobots {
			fmt.Println(i)
			for _, p := range paths {
				translatedStr := []string{""}
				var tmp []string
				s := 'A'
				for _, pi := range p {
					for _, sp := range getShortestPathsKeypad(s, pi) {
						for _, ts := range translatedStr {
							tmp = append(tmp, ts+sp+"A")
						}

					}
					translatedStr = tmp
					tmp = nil
					s = pi
				}
				newPaths = append(newPaths, translatedStr...)
			}
			paths = newPaths
			newPaths = nil

			// Find shortest path and filter out paths that are longer....

			//fmt.Println("Before filter: ", len(paths))
			paths = filterShortestPaths(paths)
			fmt.Println("Before best score: ", len(paths))
			paths = filterBestScore(paths)
			fmt.Println("After: ", len(paths))
			//fmt.Println(len(paths))

		}

		//fmt.Println(paths)
		/*for _, p := range paths {
			if p == "<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A" {
				fmt.Println("OK!!!")
			}
		}*/

		fmt.Println("Shortest path len:", len(paths[0]))
		shortestPathLengths = append(shortestPathLengths, len(paths[0]))

		paths = []string{""}
		newPaths = nil
		start = Pos{3, 2}
	}

	fmt.Println(shortestPathLengths)

	complexity := 0
	for i := range shortestPathLengths {
		complexity += shortestPathLengths[i] * getCodeValue(codes[i])
	}

	fmt.Println("Part 1 (complexity): ", complexity)
}

func Solve() {
	codes := readInput("solutions/21/input.txt")
	fmt.Println(codes)

	grid := getGrid()
	getSequence(grid, codes, 2)

	getSequence(grid, codes, 25)

}
