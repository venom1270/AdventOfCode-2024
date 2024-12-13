package solutions_day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
}

type BigPos struct {
	x int64
	y int64
}

type Machine struct {
	prizePos Pos
	buttonA  Pos
	buttonB  Pos
}

type BigMachine struct {
	prizePos BigPos
	buttonA  Pos
	buttonB  Pos
}

func readInput(filename string) []Machine {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var machines []Machine
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()

		s := strings.Split(line1, " ")
		xTmp := strings.Split(s[2], "+")[1]
		ax, _ := strconv.Atoi(xTmp[:len(xTmp)-1])
		ay, _ := strconv.Atoi(strings.Split(s[3], "+")[1])

		s = strings.Split(line2, " ")
		xTmp = strings.Split(s[2], "+")[1]
		bx, _ := strconv.Atoi(xTmp[:len(xTmp)-1])
		by, _ := strconv.Atoi(strings.Split(s[3], "+")[1])

		s = strings.Split(line3, " ")
		xTmp = strings.Split(s[1], "=")[1]
		xVal, _ := strconv.Atoi(xTmp[:len(xTmp)-1])
		yVal, _ := strconv.Atoi(strings.Split(s[2], "=")[1])

		machines = append(machines, Machine{Pos{xVal, yVal}, Pos{ax, ay}, Pos{bx, by}})

		scanner.Scan()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return machines
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func dfs_rec(machine Machine, clawPos Pos, price int, buttonPresses Pos) int {
	if buttonPresses.x > 100 || buttonPresses.y > 100 {
		return -1
	}

	if clawPos == machine.prizePos {
		return price
	}
	if clawPos.x > machine.prizePos.x || clawPos.y > machine.prizePos.y {
		return -1
	}

	pressA := -1
	pressB := -1

	if clawPos.x != 0 && clawPos.y != 0 && machine.prizePos.x%clawPos.x == 0 && machine.prizePos.y%clawPos.y == 0 && machine.prizePos.x/clawPos.x == machine.prizePos.y/clawPos.y {
		skip := machine.prizePos.x / clawPos.x
		pressA = dfs_rec(machine, Pos{clawPos.x * skip, clawPos.y * skip}, price*skip, Pos{buttonPresses.x * skip, buttonPresses.y * skip})
	} else {
		el, ok := memo[clawPos]
		if ok {
			if el <= price {
				return -1
			} else {
				fmt.Println(el, price)
				memo[clawPos] = price
			}
		} else {
			memo[clawPos] = price
		}

		pressA = dfs_rec(machine, Pos{clawPos.x + machine.buttonA.x, clawPos.y + machine.buttonA.y}, price+3, Pos{buttonPresses.x + 1, buttonPresses.y})
		pressB = dfs_rec(machine, Pos{clawPos.x + machine.buttonB.x, clawPos.y + machine.buttonB.y}, price+1, Pos{buttonPresses.x, buttonPresses.y + 1})
	}

	if pressA == -1 && pressB == -1 {
		return -1
	} else if pressA == -1 {
		return pressB
	} else if pressB == -1 {
		return pressA
	} else {
		return min(pressA, pressB)
	}
}

var memo map[Pos]int

func dfs(machines []Machine) int {
	minPrice := 0
	for _, m := range machines {
		memo = map[Pos]int{}
		p := dfs_rec(m, Pos{0, 0}, 0, Pos{0, 0})
		if p != -1 {
			//fmt.Println("Machine has price ", p, m)
			minPrice += p
		} else {
			//fmt.Println("Machine has no possible combinations", m)
		}
	}
	return minPrice
}

func cr(machine BigMachine) int {

	var x int64 = 0
	var y int64 = 0

	D := int64(machine.buttonA.x*machine.buttonB.y - machine.buttonA.y*machine.buttonB.x)
	Dx := machine.prizePos.x*int64(machine.buttonB.y) - machine.prizePos.y*int64(machine.buttonB.x)
	Dy := machine.prizePos.y*int64(machine.buttonA.x) - machine.prizePos.x*int64(machine.buttonA.y)

	x = Dx / D
	y = Dy / D

	checkX := x*int64(machine.buttonA.x) + y*int64(machine.buttonB.x)
	checkY := x*int64(machine.buttonA.y) + y*int64(machine.buttonB.y)

	if checkX == machine.prizePos.x && checkY == machine.prizePos.y {
		return int(x*3 + y)
	} else {
		return 0
	}

}

func cramersRule(machines []BigMachine) int {
	minPrice := 0
	for _, m := range machines {
		p := cr(m)
		if p != 0 {
			//fmt.Println("Machine has price ", p, m)
		} else {
			//fmt.Println("Machine has no possible combinations", m)
		}
		minPrice += p
	}

	return minPrice
}

func Solve() {
	machines := readInput("solutions/13/input.txt")
	//fmt.Println(machines)

	// Part 1 -- DFS search
	fmt.Println(dfs(machines))

	// Part 2 -- linear equasion solve - Cramer rule
	var bigMachines []BigMachine
	for i := range machines {
		var p2 int64 = 10000000000000
		bigMachines = append(bigMachines, BigMachine{BigPos{int64(machines[i].prizePos.x) + p2, int64(machines[i].prizePos.y) + p2}, machines[i].buttonA, machines[i].buttonB})

	}
	fmt.Println(cramersRule(bigMachines))
}
