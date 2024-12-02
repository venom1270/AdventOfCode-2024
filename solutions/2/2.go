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

	var list [][]int
	var tmp []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		tmp = nil
		splitLine := strings.Split(line, " ")
		for _, el := range splitLine {
			num, err := strconv.Atoi(el)
			if err != nil {
				panic(err)
			}
			tmp = append(tmp, num)
		}

		list = append(list, tmp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return list
}

func difference(x int, y int) int {
	d := x - y
	if d < 0 {
		return -d
	} else {
		return d
	}
}

func checkIncreasing(x int, y int, increasing int) bool {
	if increasing == 1 {
		return x < y
	} else if increasing == -1 {
		return x > y
	} else {
		return true
	}
}

func getIncreasing(x int, y int) int {
	if x < y {
		return 1
	} else {
		return -1
	}
}

func checkDifference(x int, y int) bool {
	diff := difference(x, y)
	if diff >= 1 && diff <= 3 {
		return true
	} else {
		return false
	}
}

func checkReport(report []int) bool {
	isIncreasing := 0
	for i, el := range report {
		if i+1 == len(report) {
			return true
		}

		el2 := report[i+1]
		if !checkDifference(el, el2) {
			break
		}

		if isIncreasing == 0 {
			isIncreasing = getIncreasing(el, el2)
		}
		if !checkIncreasing(el, el2, isIncreasing) {
			break
		}
	}

	return false
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func Solve() {
	reports := readInput("solutions/2/test.txt")
	//fmt.Println(reports)

	// Part 1
	safeCount := 0
	for _, line := range reports {
		if checkReport(line) {
			safeCount++
		}
	}

	fmt.Println(safeCount)

	// Part 2 -- dampener
	safeCount = 0
	for _, line := range reports {
		if checkReport(line) {
			safeCount++
		} else {
			// Try to remove each element
			tmp := make([]int, len(line))
			for i := range len(line) {
				copy(tmp, line)
				shortenedReport := append(tmp[0:i], tmp[min(i+1, len(line)):]...)
				//fmt.Println(shortenedReport)
				if checkReport(shortenedReport) {
					//fmt.Println(line)
					safeCount++
					break
				}
			}
		}
	}

	fmt.Println(safeCount)

}
