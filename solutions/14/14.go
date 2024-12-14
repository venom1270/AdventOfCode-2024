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

type Robot struct {
	pos      Pos
	velocity Pos
}

func readInput(filename string) []Robot {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var robots []Robot
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, " ")

		s11 := strings.Split(s[0], "=")[1]
		s11s := strings.Split(s11, ",")

		px, _ := strconv.Atoi(s11s[0])
		py, _ := strconv.Atoi(s11s[1])

		s11 = strings.Split(s[1], "=")[1]
		s11s = strings.Split(s11, ",")

		vx, _ := strconv.Atoi(s11s[0])
		vy, _ := strconv.Atoi(s11s[1])

		robots = append(robots, Robot{Pos{px, py}, Pos{vx, vy}})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return robots
}

func getSafetyFactor(robots []Robot, W int, H int) int {
	var q1, q2, q3, q4 int
	for _, r := range robots {

		if W%2 == 1 {
			if W/2 == r.pos.x {
				continue
			}
		}
		if H%2 == 1 {
			if H/2 == r.pos.y {
				continue
			}
		}

		if r.pos.x < W/2 {
			if r.pos.y < H/2 {
				q1++
			} else {
				q4++
			}
		} else {
			if r.pos.y < H/2 {
				q2++
			} else {
				q3++
			}
		}
	}
	return q1 * q2 * q3 * q4
}

func robotOnPos(robots []Robot, pos Pos) bool {
	for _, r := range robots {
		//fmt.Println("Checking: ", r.pos.x, pos.x, r.pos.y, pos.y)
		if r.pos.x == pos.x && r.pos.y == pos.y {
			return true
		}
	}
	return false
}

func findTreeAtPos(robots []Robot, W int, H int, pos Pos) bool {
	x := pos.x
	y := pos.y
	if robotOnPos(robots, pos) {
		height := 1
		width := 1
		// Find triangle
		for {
			if x-width < 0 || x+width >= W || y+height >= H {
				break
			}
			end := false
			for xx := x - width; xx <= x+width; xx++ {
				if !robotOnPos(robots, Pos{xx, y + height}) {
					if xx == x-width {
						end = true
						break
					}
					//fmt.Println("ENDING", xx, height, width)
					return false
				}

			}
			if end {
				break
			}
			height++
			width++
		}
		//fmt.Println("wqer")
		if height < 2 {
			return false
		}

		// End of triangle, check if root is here
		if y+height+1 >= H {
			return false
		}

		if robotOnPos(robots, Pos{x, y + height}) && robotOnPos(robots, Pos{x, y + height + 1}) {
			// At least one+1 robot as root
			if height < 5 {
				return false
			}
			fmt.Println("Found! H W", height, width)
			return true
		}

	}
	return false
}

func findTree(robots []Robot, W int, H int) bool {
	for x := range W {
		for y := range H {
			if findTreeAtPos(robots, W, H, Pos{x, y}) {
				fmt.Println(x, y)
				return true
			}
		}
	}
	return false
}

func getRobotsAtTime(robots []Robot, W int, H int, time int) []Robot {
	var newPositions []Robot
	for _, r := range robots {

		vx := r.velocity.x
		vy := r.velocity.y
		if vx < 0 {
			vx = W + vx
		}
		if vy < 0 {
			vy = H + vy
		}

		newPos := Pos{(r.pos.x + vx*time) % W, (r.pos.y + vy*time) % H}
		newPositions = append(newPositions, Robot{newPos, r.velocity})
	}
	return newPositions
}

func Solve() {
	robots := readInput("solutions/14/input.txt")
	//fmt.Println(robots)

	// Part 1
	// Simulate movement
	TIME := 100
	WIDTH := 101  // 101
	HEIGHT := 103 // 103
	newPositions := getRobotsAtTime(robots, WIDTH, HEIGHT, TIME)
	fmt.Println(getSafetyFactor(newPositions, WIDTH, HEIGHT))

	// Part 2 - simulate movement until tree is found

	//robots = append(robots, Robot{Pos{50, 50}, Pos{1, 1}})
	//robots = append(robots, Robot{Pos{49, 51}, Pos{1, 1}})
	//robots = append(robots, Robot{Pos{50, 51}, Pos{1, 1}})
	//robots = append(robots, Robot{Pos{51, 51}, Pos{1, 1}})
	//robots = append(robots, Robot{Pos{50, 52}, Pos{1, 1}})

	//fmt.Println(findTreeAtPos(robots, WIDTH, HEIGHT, Pos{50, 50}))

	seconds := 0
	newPositions = getRobotsAtTime(robots, WIDTH, HEIGHT, 0)
	for !findTree(newPositions, WIDTH, HEIGHT) {
		seconds++
		fmt.Println("Advancing time...", seconds)
		newPositions = getRobotsAtTime(robots, WIDTH, HEIGHT, seconds)
	}

	fmt.Println("Tree found!")
}
