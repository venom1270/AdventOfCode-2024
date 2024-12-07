package solutions_day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) [][]int {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var nums [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ": ")
		var row []int
		n1, _ := strconv.Atoi(s[0])
		row = append(row, n1)
		for _, el := range strings.Split(s[1], " ") {
			n, _ := strconv.Atoi(el)
			row = append(row, n)
		}
		nums = append(nums, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nums
}

type Operation int

const (
	ADD Operation = iota
	MULTIPLY
	CONCATENATE
)

func testOperation(nums []int, currentIndex int, result int, op Operation, part2 bool) bool {
	if currentIndex >= len(nums) {
		return result == nums[0]
	}
	if result > nums[0] {
		return false
	}

	switch op {
	case ADD:
		result = result + nums[currentIndex]
	case MULTIPLY:
		result = result * nums[currentIndex]
	case CONCATENATE:
		result, _ = strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(nums[currentIndex]))
	default:
		fmt.Println("WHAT??")
	}

	if testOperation(nums, currentIndex+1, result, ADD, part2) {
		return true
	}
	if testOperation(nums, currentIndex+1, result, MULTIPLY, part2) {
		return true
	}
	if part2 {
		if testOperation(nums, currentIndex+1, result, CONCATENATE, part2) {
			return true
		}
	}

	return false
}

func initTest(nums []int, part2 bool) bool {
	if testOperation(nums, 2, nums[1], ADD, part2) {
		return true
	}
	if testOperation(nums, 2, nums[1], MULTIPLY, part2) {
		return true
	}
	if part2 {
		if testOperation(nums, 2, nums[1], CONCATENATE, part2) {
			return true
		}
	}

	return false
}

func Solve() {
	nums := readInput("solutions/7/input.txt")
	fmt.Println(nums)

	// Part 1
	sum := 0
	sum2 := 0
	for _, row := range nums {
		if initTest(row, false) {
			sum += row[0]
		}
		if initTest(row, true) {
			sum2 += row[0]
		}
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}
