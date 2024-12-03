package solutions_day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput(filename string) []string {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var list []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		list = append(list, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return list
}

func Solve() {
	lines := readInput("solutions/3/input.txt")

	//fmt.Println(lines)

	// Part 1 (sum) and Part 2 (sum2)
	sum := 0
	sum2 := 0
	enabled := true
	for _, line := range lines {
		i := 0
		for i < len(line) {
			valid := true
			if i+3 < len(line) && line[i:i+4] == "mul(" {
				j := i + 4
				num1 := ""
				num2 := ""
				for j < len(line) && line[j] >= '0' && line[j] <= '9' {
					num1 += string([]byte{line[j]})
					j++
				}
				if len(num1) > 3 {
					valid = false
				}
				if line[j] != ',' {
					valid = false
				}
				j++
				for line[j] >= '0' && line[j] <= '9' {
					num2 += string([]byte{line[j]})
					j++
				}
				if len(num2) > 3 {
					valid = false
				}
				if line[j] != ')' {
					valid = false
				}

				if valid {
					i = j
					n1, _ := strconv.Atoi(num1)
					n2, _ := strconv.Atoi(num2)
					mul := n1 * n2
					sum += mul
					if enabled {
						sum2 += mul
					}
				}
			} else if i+3 < len(line) && line[i:i+4] == "do()" {
				enabled = true
				i += 3
			} else if i+6 < len(line) && line[i:i+7] == "don't()" {
				enabled = false
				i += 6
			}

			i++
		}
	}

	fmt.Println(sum)
	fmt.Println(sum2)

}
