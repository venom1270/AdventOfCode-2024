package solutions_day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func Solve() {
	readInput("solutions/1/input.txt")
	fmt.Println("Solving...")

}
