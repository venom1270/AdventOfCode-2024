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

func evenDigits(num int) bool {
	even := false
	for num/10 != 0 {
		num /= 10
		even = !even
	}
	return even
}

type Pos struct {
	num          int
	depth        int
	nextNumCount int
	depthDiff    int
}

var memo map[int]Pos

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

func memoization2(num int, depth int) int {
	// TODO next time

	if depth >= MAX_DEPTH {
		return 1
	}

	if val, ok := memo[num]; ok {
		// CYCLE
		fmt.Println("Cycle", num, depth, val.depth)
		val.depthDiff = depth - val.depth
		memo[num] = val
		return 0
	} else {
		m := Pos{num, depth, 0, 0}
		memo[num] = m
	}

	sum := 0

	if num == 0 {
		sum = memoization2(1, depth+1)
	} else if s := strconv.Itoa(num); len(s)%2 == 0 {
		num1, _ := strconv.Atoi(s[:len(s)/2])
		num2, _ := strconv.Atoi(s[len(s)/2:])
		sum = memoization2(num1, depth+1) + naive(num2, depth+1)
	} else {
		sum = memoization2(num*2024, depth+1)
	}

	val, _ := memo[num]
	if val.depthDiff == 0 {
		return sum
	} else {
		fmt.Println("Working")
	}
	val.nextNumCount = sum
	memo[num] = val
	numCycles := (MAX_DEPTH - depth) / val.depthDiff

	fmt.Println("cycle of num with depthDiff depth numCycles", num, val.depthDiff, depth, numCycles)

	sum = sum*numCycles + naive(num, depth+numCycles*val.depthDiff)

	return sum

}

var simpleMemo map[Pos]int

func memoization(num int, depth int) int {
	// TODO next time

	if depth >= MAX_DEPTH {
		return 1
	}

	key := Pos{num, depth, 0, 0}
	if val, ok := simpleMemo[key]; ok {
		return val
	}

	sum := 0

	if num == 0 {
		sum = memoization(1, depth+1)
	} else if s := strconv.Itoa(num); len(s)%2 == 0 {
		num1, _ := strconv.Atoi(s[:len(s)/2])
		num2, _ := strconv.Atoi(s[len(s)/2:])
		sum = memoization(num1, depth+1) + naive(num2, depth+1)
	} else {
		sum = memoization(num*2024, depth+1)
	}

	simpleMemo[key] = sum

	return sum

}

func Solve() {
	nums := readInput("solutions/11/test.txt")
	fmt.Println(nums)

	// Part 1
	sum := 0
	MAX_DEPTH = 20
	for _, num := range nums {
		sum += naive(num, 0)
	}
	fmt.Println("Part 1: ", sum)

	// Part 2
	sum = 0
	MAX_DEPTH = 20
	memo = map[int]Pos{}
	simpleMemo = map[Pos]int{}
	for _, num := range nums {
		sum += memoization2(num, 0)
	}

	fmt.Println("Part 2: ", sum)

}
