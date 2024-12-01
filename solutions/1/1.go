package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput(filename string) ([]int, []int) {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var list1 []int
	var list2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		splitLine := strings.Split(line, "   ")
		num1, err := strconv.Atoi(splitLine[0])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(splitLine[1])
		if err != nil {
			panic(err)
		}
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return list1, list2
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func Solve1() {
	list1, list2 := readInput("solutions/1/input.txt")
	fmt.Println("Solving...")

	sort.Ints(list1)
	sort.Ints(list2)

	// Part 1

	var sum int = 0
	for i, _ := range list1 {
		sum += abs(list2[i] - list1[i])
	}

	fmt.Println(sum)

	// Part 2
	sum = 0
	var pointer int = 0
	for i, _ := range list1 {
		// Find the number in the second list
		number := list1[i]
		for pointer < len(list2) && list2[pointer] < number {
			pointer++
		}
		var p2 = pointer
		if list2[pointer] == number {
			for p2 < len(list2) && list2[p2] == number {
				sum += number
				p2++
			}
		}
	}

	fmt.Println(sum)

}
