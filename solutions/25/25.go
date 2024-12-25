package solutions_day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(filename string) [][]string {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input [][]string
	scanner := bufio.NewScanner(file)
	var schematic []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			input = append(input, schematic)
			schematic = []string{}
			continue
		}
		schematic = append(schematic, line)
	}

	input = append(input, schematic)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}

func splitKeysLocks(schematics [][]string) ([][]int, [][]int) {
	var keys, locks [][]int
	for _, s := range schematics {
		var nums []int
		if s[0] == "#####" {
			for col := 0; col < len(s[0]); col++ {
				var row int
				for row = 0; row < len(s) && s[row][col] == '#'; row++ {
				}
				nums = append(nums, row-1)
			}
			locks = append(locks, nums)
		} else {
			for col := 0; col < len(s[0]); col++ {
				var row int
				for row = len(s) - 1; row >= 0 && s[row][col] == '#'; row-- {
				}
				nums = append(nums, len(s)-2-row)
			}
			keys = append(keys, nums)
		}
	}

	return keys, locks
}

func part1(locks, keys [][]int, height int) {

	sum := 0
	for _, lock := range locks {
		for _, key := range keys {
			ok := true
			for i := 0; i < len(lock); i++ {
				if lock[i]+key[i] > height {
					ok = false
					break
				}
			}
			if ok {
				sum++
			}
		}
	}

	fmt.Println("Part 1:", sum)
}

func Solve() {
	schematics := readInput("solutions/25/input.txt")

	//fmt.Println(schematics)

	keys, locks := splitKeysLocks(schematics)

	//fmt.Println(keys)
	//fmt.Println(locks)

	height := len(schematics[0]) - 2
	part1(locks, keys, height)
}
