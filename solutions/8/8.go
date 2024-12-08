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

type Coords struct {
	i, j int
}

func validCoords(c Coords, N int, M int) bool {
	if c.i < 0 || c.j < 0 || c.i >= N || c.j >= M {
		return false
	}
	return true
}

func contains(arr []Coords, c Coords) bool {
	for _, el := range arr {
		if el.i == c.i && el.j == c.j {
			return true
		}
	}
	return false
}

func Solve() {
	grid := readInput("solutions/8/input.txt")
	//fmt.Println(grid)

	m := make(map[rune][]Coords)

	N, M := len(grid), len(grid[0])
	// For part 1
	var antinodes []Coords
	// For part 2
	var antinodes2 []Coords

	for i, row := range grid {
		for j, el := range row {
			if el != '.' {
				if m[el] != nil {
					indexes := Coords{i, j}
					m[el] = append(m[el], indexes)
				} else {
					m[el] = []Coords{{i, j}}
				}
				antinodes2 = append(antinodes2, Coords{i, j})
			}
		}
	}

	for _, v := range m {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				a1 := v[i]
				a2 := v[j]

				distanceI := a2.i - a1.i
				distanceJ := a2.j - a1.j
				if distanceJ < 0 {
					distanceJ *= -1
				}

				var antinode1, antinode2 Coords
				var a1Diff, a2Diff Coords

				if a1.j < a2.j {
					antinode1 = Coords{a1.i - distanceI, a1.j - distanceJ}
					antinode2 = Coords{a2.i + distanceI, a2.j + distanceJ}
					a1Diff = Coords{-distanceI, -distanceJ}
					a2Diff = Coords{distanceI, distanceJ}
				} else {
					antinode1 = Coords{a1.i - distanceI, a1.j + distanceJ}
					antinode2 = Coords{a2.i + distanceI, a2.j - distanceJ}
					a1Diff = Coords{-distanceI, distanceJ}
					a2Diff = Coords{distanceI, -distanceJ}
				}

				//fmt.Println("Coords: ", a1, a2, " have antinodes: ", antinode1, antinode2)

				if validCoords(antinode1, N, M) && !contains(antinodes, antinode1) {
					antinodes = append(antinodes, antinode1)
				}
				if validCoords(antinode2, N, M) && !contains(antinodes, antinode2) {
					antinodes = append(antinodes, antinode2)
				}

				// Part 2
				// We use tmp because saome values are getting overwritten otherwise (reference)
				tmp := Coords{antinode1.i, antinode1.j}
				for validCoords(tmp, N, M) {
					if !contains(antinodes2, tmp) {
						antinodes2 = append(antinodes2, Coords{tmp.i, tmp.j})
					}
					tmp = Coords{tmp.i + a1Diff.i, tmp.j + a1Diff.j}
				}

				tmp = Coords{antinode2.i, antinode2.j}
				for validCoords(tmp, N, M) {
					if !contains(antinodes2, tmp) {
						antinodes2 = append(antinodes2, Coords{tmp.i, tmp.j})
					}
					tmp = Coords{tmp.i + a2Diff.i, tmp.j + a2Diff.j}
				}

			}
		}
	}

	fmt.Println(len(antinodes))
	fmt.Println(len(antinodes2))

	/*for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if contains(antinodes2, Coords{i, j}) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}*/

}
