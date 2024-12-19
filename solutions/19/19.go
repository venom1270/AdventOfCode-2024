package solutions_day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput(filename string) ([]string, []string) {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var patterns []string
	var towels []string
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	patterns = append(patterns, strings.Split(line, ", ")...)

	scanner.Scan()
	for scanner.Scan() {
		line = scanner.Text()
		towels = append(towels, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return patterns, towels
}

func findPatterns(i int, towel string, patterns []string) bool {
	if i == len(towel) {
		return true
	}
	if i > len(towel) {
		return false
	}

	for _, pattern := range patterns {
		if len(pattern)+i > len(towel) {
			continue
		}
		if towel[i:i+len(pattern)] == pattern {
			if findPatterns(i+len(pattern), towel, patterns) {
				return true
			}
		}
	}
	return false
}

var memo map[int]int

func findPatternsNumMatches(i int, towel string, patterns []string) int {
	if i == len(towel) {
		return 1
	}
	if i > len(towel) {
		return 0
	}

	sum := 0
	if val, ok := memo[i]; ok {
		return val
	}

	for _, pattern := range patterns {
		if len(pattern)+i > len(towel) {
			continue
		}
		if towel[i:i+len(pattern)] == pattern {
			sum += findPatternsNumMatches(i+len(pattern), towel, patterns)
		}
	}

	memo[i] = sum

	return sum
}

func part1(patterns []string, towels []string) int {
	possible := 0
	for _, towel := range towels {
		if findPatterns(0, towel, patterns) {
			possible++
		}
	}
	return possible
}

func part2(patterns []string, towels []string) int {
	possible := 0
	for _, towel := range towels {
		//fmt.Println(towel)
		memo = map[int]int{}
		possible += findPatternsNumMatches(0, towel, patterns)
	}
	return possible
}

func Solve() {
	patterns, towels := readInput("solutions/19/input.txt")
	fmt.Println(patterns)
	fmt.Println(towels)
	fmt.Println(part1(patterns, towels))
	fmt.Println(part2(patterns, towels))
}
