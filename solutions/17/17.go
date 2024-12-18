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
		commandString += strconv.Itoa(el) + ","
	}
	commandString = commandString[:len(commandString)-1]

	var numsToCheck []int
	numsToCheck = append(numsToCheck, 0)
	numsToCheck = append(numsToCheck, 1)
	numsToCheck = append(numsToCheck, 2)
	numsToCheck = append(numsToCheck, 3)
	numsToCheck = append(numsToCheck, 4)
	numsToCheck = append(numsToCheck, 5)
	numsToCheck = append(numsToCheck, 6)
	numsToCheck = append(numsToCheck, 7)

	count := 1
	for len(numsToCheck) > 0 {
		if len(numsToCheck) == 0 {
			fmt.Println("FAIL!!!")
		}
		var newNums []int
		for i := 0; i <= 7; i++ {
			//fmt.Println(i)

			for _, num := range numsToCheck {
				aVal := num*pow(2, 3) + i
				outStr := simulate(aVal, b, c, commands)
				//fmt.Println("Checking ", outStr, commandString)
				if outStr == commandString {
					return aVal
				}
				if len(outStr)/2 == count && len(outStr) <= len(commandString) && (outStr == commandString[:len(outStr)] || outStr == commandString[len(commandString)-len(outStr):]) {
					//fmt.Println("SAME!")
					newNums = append(newNums, aVal)
				}
			}

		}
		numsToCheck = newNums
		//fmt.Println(numsToCheck)
		count++
	}

	return -1
}

func Solve() {
	commands, a, b, c := readInput("solutions/17/input.txt")
	fmt.Println(commands, a, b, c)
	// Part 1
	fmt.Println(simulate(a, b, c, commands))
	// Part 2
	fmt.Println(findCopy(a, b, c, commands))
}
