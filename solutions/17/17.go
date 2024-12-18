package solutions_day2

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) ([]int, int, int, int) {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var commands []int
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var a, b, c int
	line := scanner.Text()
	a, _ = strconv.Atoi(strings.Split(line, ": ")[1])
	scanner.Scan()
	line = scanner.Text()
	b, _ = strconv.Atoi(strings.Split(line, ": ")[1])
	scanner.Scan()
	line = scanner.Text()
	c, _ = strconv.Atoi(strings.Split(line, ": ")[1])
	scanner.Scan()
	scanner.Scan()
	line = scanner.Text()
	for _, el := range strings.Split(strings.Split(line, ": ")[1], ",") {
		num, _ := strconv.Atoi(el)
		commands = append(commands, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return commands, a, b, c
}

func getComboValue(combo int, a int, b int, c int) int {
	if combo <= 3 {
		return combo
	}
	if combo == 4 {
		return a
	} else if combo == 5 {
		return b
	} else {
		return c
	}
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func step(a int, b int, c int, commands []int, iPtr int) (int, int, int, int, string) {
	output := ""
	comboVal := getComboValue(commands[iPtr+1], a, b, c)
	literalVal := commands[iPtr+1]
	//fmt.Println("COmmands: ", commands[iPtr], literalVal)
	switch commands[iPtr] {
	case 0: // adv
		a = a / pow(2, comboVal)
	case 1: // bxl
		b = b ^ literalVal
	case 2: // bst
		b = comboVal % 8
	case 3: // jnz
		if a != 0 {
			iPtr = literalVal - 2
		}
	case 4: // bxc
		b = b ^ c
	case 5: // out
		out := comboVal % 8
		output = strconv.Itoa(out)
		//fmt.Println(out)
	case 6: // bdv
		b = a / pow(2, comboVal)
	case 7: // cdv
		c = a / pow(2, comboVal)
	}

	//fmt.Println("Registers: ", a, b, c)

	iPtr += 2

	return a, b, c, iPtr, output
}

func simulate(a int, b int, c int, commands []int) string {
	outputString := ""
	var out string
	iPtr := 0
	for iPtr < len(commands) {
		a, b, c, iPtr, out = step(a, b, c, commands, iPtr)
		if out != "" {
			outputString += out + ","
		}
	}

	if len(outputString) > 1 {
		outputString = outputString[:len(outputString)-1]
	}

	return outputString
}

func findCopy(a, b, c int, commands []int) int {

	// Transform commands to string for easier comparison
	commandString := ""
	for _, el := range commands {
		commandString += strconv.Itoa(el)
	}

	fmt.Println(len(commandString))
	fmt.Println(len(simulate(0, b, c, commands)))
	fmt.Println(len(simulate(10000, b, c, commands)))
	fmt.Println(len(simulate(1000000, b, c, commands)))
	fmt.Println(len(simulate(1000000000, b, c, commands)))
	fmt.Println(len(simulate(10000000000, b, c, commands)))
	fmt.Println(len(simulate(100000000000, b, c, commands)))
	fmt.Println(len(simulate(1000000000000, b, c, commands)))
	fmt.Println(len(simulate(30000000000000, b, c, commands)))
	fmt.Println(len(simulate(290000000000000, b, c, commands)))
	fmt.Println(len(simulate(100000000000000, b, c, commands)))
	//return 123

	var ai int
	// 10000000000
	// 13710000000
	//200000000

	start := 30000000000000
	end := 290000000000000

	for ai = start; ai < end; ai += 8 {
		if ai%10000000000 == 0 {
			fmt.Println(float64(ai-start) * 100 / float64(end-start))
		}
		aTest, bTest, cTest := ai, b, c
		iPtr := 0
		out := ""
		outStr := ""
		for iPtr < len(commands) {
			aTest, bTest, cTest, iPtr, out = step(aTest, bTest, cTest, commands, iPtr)
			if out != "" {
				outStr += out
				if outStr != commandString[:len(outStr)] {
					break
				}
			}
		}

		if outStr == commandString {
			fmt.Println(outStr, commandString)
			break
		}
	}

	return ai
}

func Solve() {
	commands, a, b, c := readInput("solutions/17/input.txt")
	fmt.Println(commands, a, b, c)
	// Part 1
	fmt.Println(simulate(a, b, c, commands))
	// Part 2
	fmt.Println(findCopy(a, b, c, commands))
}
