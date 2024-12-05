package solutions_day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readInput(filename string) ([][]int, [][]int) {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var rules [][]int
	var pages [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 1 {
			break
		}
		s := strings.Split(line, "|")
		num1, _ := strconv.Atoi(s[0])
		num2, _ := strconv.Atoi(s[1])
		tmp := []int{num1, num2}
		rules = append(rules, tmp)

	}
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ",")
		var tmp []int
		for _, p := range s {
			num, _ := strconv.Atoi(p)
			tmp = append(tmp, num)
		}
		pages = append(pages, tmp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rules, pages
}

func check(m map[int][]int, p int, page []int) (bool, int) {
	rule := m[p]
	if rule == nil {
		return true, -1
	}

	for i, curPage := range page {
		if curPage == p {
			break
		}
		if slices.Contains(rule, curPage) {
			return false, i
		}
	}

	return true, -1
}

func insertAt(arr []int, i int, j int) {
	x := arr[j]
	k := j
	for k > i {
		arr[k] = arr[k-1]
		k--
	}
	arr[i] = x
}

func Solve() {
	rules, pages := readInput("solutions/5/input.txt")

	fmt.Println(rules)
	fmt.Println(pages)

	// Part 1
	var m map[int][]int
	m = make(map[int][]int)
	for _, rule := range rules {
		key := rule[0]
		val := rule[1]
		if _, ok := m[key]; ok {
			m[key] = append(m[key], val)
		} else {
			m[key] = []int{val}
		}
	}

	sum := 0
	var unorderedPages [][]int
	for _, page := range pages {
		ok := true
		for _, p := range page {
			if b, _ := check(m, p, page); !b {
				ok = false
				break
			}
		}

		if ok {
			sum += page[len(page)/2]
		} else {
			unorderedPages = append(unorderedPages, page)
		}
	}

	fmt.Println(sum)

	// Part 2 -- order the incorrectly ordered pages
	sum = 0
	for _, page := range unorderedPages {
		for currIndex, p := range page {
			ok, i := check(m, p, page)
			if !ok {
				insertAt(page, i, currIndex)
			}
		}
		sum += page[len(page)/2]
		//fmt.Println(page)
	}

	fmt.Println(sum)

}
