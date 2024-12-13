package solutions_day2

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
}

type Machine struct {
	prizePos Pos
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
	//if buttonPresses.x > 100 || buttonPresses.y > 100 {
	//	return -1
	//}

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
		fmt.Println("qwe")
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
			fmt.Println("Machine has price ", p, m)
			minPrice += p
		} else {
			fmt.Println("Machine has no possible combinations", m)
		}
	}
	return minPrice
}

// An Item is something we manage in a priority queue.
type Item struct {
	clawPos       Pos // The value of the item; arbitrary.
	buttonPresses Pos
	price         int
	priority      int // The priority of the item in the queue. -- PRICE
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	//return pq[i].priority > pq[j].priority
	return pq[i].priority < pq[j].priority // Return lowest first
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, clawPos Pos, buttonPresses Pos, priority int) {
	item.clawPos = clawPos
	item.buttonPresses = buttonPresses
	item.priority = priority
	heap.Fix(pq, item.index)
}

func distance(p1 Pos, p2 Pos) int {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	if dx < 0 {
		dx *= -1
	}
	if dy < 0 {
		dy *= -1
	}
	return int(math.Sqrt(float64(dx*dx + dy*dy)))
}

func aStarMachine(machine Machine) int {
	minPrice := 0

	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		clawPos:       Pos{0, 0},
		buttonPresses: Pos{0, 0},
		price:         0,
		priority:      0,
		index:         0,
	}
	heap.Init(&pq)

	// Take the items out; they arrive in increasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		if item.buttonPresses.x > 100 || item.buttonPresses.y > 100 {
			continue
		}

		if item.clawPos == machine.prizePos {
			return item.priority
		}

		if item.clawPos.x >= machine.prizePos.x || item.clawPos.y >= machine.prizePos.y {
			continue
		}

		fmt.Println(item)

		if item.clawPos.x != 0 && item.clawPos.y != 0 && machine.prizePos.x%item.clawPos.x == 0 && machine.prizePos.y%item.clawPos.y == 0 && machine.prizePos.x/item.clawPos.x == machine.prizePos.y/item.clawPos.y {
			skip := machine.prizePos.x / item.clawPos.x
			fmt.Println("qwe")
			return item.priority * skip
		}

		pressA := &Item{
			clawPos:       Pos{item.clawPos.x + machine.buttonA.x, item.clawPos.y + machine.buttonA.y},
			buttonPresses: Pos{item.buttonPresses.x + 1, item.buttonPresses.y},
			price:         item.price + 3,
			priority:      item.price + 3 + distance(item.clawPos, machine.prizePos),
		}
		pressB := &Item{
			clawPos:       Pos{item.clawPos.x + machine.buttonB.x, item.clawPos.y + machine.buttonB.y},
			buttonPresses: Pos{item.buttonPresses.x, item.buttonPresses.y + 1},
			price:         item.price + 1,
			priority:      item.price + 1 + distance(item.clawPos, machine.prizePos),
		}
		heap.Push(&pq, pressA)
		heap.Push(&pq, pressB)
	}

	return minPrice
}

func aStar(machines []Machine) int {
	minPrice := 0
	for _, m := range machines {
		p := aStarMachine(m)
		minPrice += p
	}

	return minPrice
}

func Solve() {
	machines := readInput("solutions/13/test.txt")
	fmt.Println(machines)

	// Part 1
	fmt.Println(dfs(machines))

	// Part 2
	for i, _ := range machines {
		machines[i].prizePos.x += 10000000000000
		machines[i].prizePos.y += 10000000000000
	}
	fmt.Println(aStar(machines))
}
