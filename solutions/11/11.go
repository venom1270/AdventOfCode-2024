package solutions_day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) []int {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var line string
	var stones []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		for _, el := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(el)
			stones = append(stones, num)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return stones
}

var MAX_DEPTH int

func naive(num int, depth int) int {
	if depth == MAX_DEPTH {
		return 1
	}

	if num == 0 {
		return naive(1, depth+1)
	} else if s := strconv.Itoa(num); len(s)%2 == 0 {
		num1, _ := strconv.Atoi(s[:len(s)/2])
		num2, _ := strconv.Atoi(s[len(s)/2:])
		return naive(num1, depth+1) + naive(num2, depth+1)
	} else {
		return naive(num*2024, depth+1)
	}
}

type SimplePos struct {
	val   int
	depth int
}

var simpleMemo map[SimplePos]int

func memoization(num int, depth int, sum int) int {

	if depth >= MAX_DEPTH {
		return sum + 1
	}

	key := SimplePos{num, depth}
	if val, ok := simpleMemo[key]; ok {
		return val
	}

	if num == 0 {
		sum = memoization(1, depth+1, sum)
	} else if s := strconv.Itoa(num); len(s)%2 == 0 {
		num1, _ := strconv.Atoi(s[:len(s)/2])
		num2, _ := strconv.Atoi(s[len(s)/2:])
		sum = memoization(num1, depth+1, sum) + memoization(num2, depth+1, sum)
	} else {
		sum = memoization(num*2024, depth+1, sum)
	}

	simpleMemo[key] = sum

	return sum

}

func Solve() {
	nums := readInput("solutions/11/input.txt")
	fmt.Println(nums)

	// Part 1
	sum := 0
	MAX_DEPTH = 25
	for _, num := range nums {
		sum += naive(num, 0)
	}
	fmt.Println("Part 1: ", sum)

	// Part 2
	sum = 0
	MAX_DEPTH = 75
	simpleMemo = map[SimplePos]int{}
	for _, num := range nums {
		sum += memoization(num, 0, 0)
	}

	fmt.Println("Part 2: ", sum)

}
