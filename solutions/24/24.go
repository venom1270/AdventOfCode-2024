package solutions_day2

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Expr struct {
	a, b  string
	op    string
	value int
}

func readInput(filename string) map[string]Expr {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	expressions := map[string]Expr{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		s := strings.Split(line, ": ")
		num, _ := strconv.Atoi(s[1])
		expressions[s[0]] = Expr{"", "", "", num}
	}

	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, " -> ")
		expr := s[0]
		key := s[1]
		s = strings.Split(expr, " ")
		expressions[key] = Expr{s[0], s[2], s[1], -1}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return expressions
}

func calculate(expressions map[string]Expr, e string) int {
	if swapMap != nil {
		if el, ok := swapMap[e]; ok {
			e = el
		}
	}
	exp := expressions[e]

	if exp.value == -2 {
		return -2
	}
	if exp.value != -1 {
		return exp.value
	}

	// Cycle detection
	exp.value = -2
	expressions[e] = exp

	a := calculate(expressions, exp.a)
	if a == -2 {
		return -2
	}
	b := calculate(expressions, exp.b)
	if b == -2 {
		return -2
	}

	value := 0
	switch exp.op {
	case "AND":
		value = a & b
	case "OR":
		value = a | b
	case "XOR":
		value = a ^ b
	default:
		panic("Wrong OP!!!")
	}

	exp.value = value

	expressions[e] = exp

	return value
}

func part1(expressions map[string]Expr) {
	val := 0
	for key, _ := range expressions {
		if key[0] == 'z' {
			fmt.Println(key)
			keyVal := calculate(expressions, key)
			shift, _ := strconv.Atoi(key[1:])
			val += keyVal << shift
		}
	}
	fmt.Println(val)
}

func checkValid(expressions map[string]Expr) bool {
	x, y, z := 0, 0, 0
	maxShift := 0
	for key, _ := range expressions {
		if key[0] == 'z' || key[0] == 'x' || key[0] == 'y' {
			keyVal := calculate(expressions, key)
			if keyVal == -2 {
				return false
			}
			shift, _ := strconv.Atoi(key[1:])
			if shift > maxShift {
				maxShift = shift
			}
			switch key[0] {
			case 'x':
				x += keyVal << shift
			case 'y':
				y += keyVal << shift
			case 'z':
				z += keyVal << shift
			}
		}
	}

	maxShift++
	sum := (x & y) //% (1 << maxShift)
	//fmt.Println(x, y, z, sum)
	if sum == z {
		return true
	} else {
		return false
	}
}

func checkValidForAll(expressions map[string]Expr) bool {

	// TODO: dirty heuristic check currently
	for i := range 4 {
		for j := i + 1; j <= 5; j++ {
			e := copy(expressions)
			key1 := "x0" + strconv.Itoa(i)
			key2 := "x1" + strconv.Itoa(j)
			e[key1] = Expr{"", "", "", 1}
			e[key2] = Expr{"", "", "", 1}

			if !checkValid(e) {
				return false
			}
		}
	}
	return true

}

func copy(e map[string]Expr) map[string]Expr {
	c := map[string]Expr{}
	for key, exp := range e {
		c[key] = Expr{exp.a, exp.b, exp.op, exp.value}
	}
	return c
}

func swapWireInPlace(e map[string]Expr, e1, e2 string) {
	tmp := e[e1]
	e[e1] = e[e2]
	e[e2] = tmp
}

func swapWire(e1, e2 string) {
	/*tmp1 := e[e1]
	tmp2 := e[e2]
	e[e1] = e[e2]
	e[e2] = tmp*/
	swapMap[e1] = e2
	swapMap[e2] = e1
}

var swapMap map[string]string

func getErrorWires(expressions map[string]Expr) ([]string, int) {
	x, y, z := 0, 0, 0
	maxShift := 0
	/*for key, _ := range expressions {
		if key[0] == 'z' || key[0] == 'x' || key[0] == 'y' {
			keyVal := calculate(expressions, key)
			shift, _ := strconv.Atoi(key[1:])
			if shift > maxShift {
				maxShift = shift
			}
			switch key[0] {
			case 'x':
				x += keyVal << shift
			case 'y':
				y += keyVal << shift
			case 'z':
				z += keyVal << shift
			}
		}
	}*/

	carry := 0
	for i := 0; i <= 44; i++ {
		var xStr, yStr, zStr string
		if i < 10 {
			xStr = "x0" + strconv.Itoa(i)
			yStr = "y0" + strconv.Itoa(i)
			zStr = "z0" + strconv.Itoa(i)
		} else {
			xStr = "x" + strconv.Itoa(i)
			yStr = "y" + strconv.Itoa(i)
			zStr = "z" + strconv.Itoa(i)
		}

		xVal, yVal, zVal := calculate(expressions, xStr), calculate(expressions, yStr), calculate(expressions, zStr)
		fmt.Println(xVal, yVal, zVal, expressions[zStr], carry)
		if xVal^yVal^carry != zVal {
			fmt.Println("Error", i)
			return getDependantWires(expressions, zStr), 0
		}
		carry = xVal ^ yVal ^ carry ^ 1
	}

	var wires []string
	for maxShift >= 0 {
		var _, _, e3 string
		if maxShift < 10 {
			//e1 = "x0" + strconv.Itoa(maxShift)
			//e2 = "y0" + strconv.Itoa(maxShift)
			e3 = "z0" + strconv.Itoa(maxShift)
		} else {
			//e1 = "x" + strconv.Itoa(maxShift)
			//e2 = "y" + strconv.Itoa(maxShift)
			e3 = "z" + strconv.Itoa(maxShift)
		}
		//if z&(1<<maxShift) != expressions[e1].value^expressions[e2].value {
		if z&(1<<maxShift) != (x&(1<<maxShift))^(y&(1<<maxShift)) {
			wires = append(wires, e3)
		}
		maxShift--
	}

	//fmt.Println(wires)

	var allWires []string
	for _, w := range wires {
		dw := getDependantWires(expressions, w)
		if dw == nil {
			return nil, 999999
		}
		for _, el := range dw {
			if !slices.Contains(allWires, el) {
				allWires = append(allWires, el)
			}
		}
		//allWires = append(allWires, dw...)
	}

	//fmt.Println(allWires)
	//fmt.Println(len(allWires))
	//return removeDuplicate(allWires), len(wires)
	//fmt.Println(len(wires))
	return allWires, len(wires)

}

func getErrorWiresDecimalCheck(expressions map[string]Expr) ([]string, int) {
	x, y, z := 0, 0, 0
	maxShift := 0
	for key, _ := range expressions {
		if key[0] == 'z' || key[0] == 'x' || key[0] == 'y' {
			keyVal := calculate(expressions, key)
			shift, _ := strconv.Atoi(key[1:])
			if shift > maxShift {
				maxShift = shift
			}
			switch key[0] {
			case 'x':
				x += keyVal << shift
			case 'y':
				y += keyVal << shift
			case 'z':
				z += keyVal << shift
			}
		}
	}

	var wires []string
	//sum := x + y
	for maxShift >= 0 {
		if z-(1<<maxShift) > 0 {
			// TODO i guess
		}
	}

	//fmt.Println(wires)

	var allWires []string
	for _, w := range wires {
		dw := getDependantWires(expressions, w)
		if dw == nil {
			return nil, 999999
		}
		for _, el := range dw {
			if !slices.Contains(allWires, el) {
				allWires = append(allWires, el)
			}
		}
		//allWires = append(allWires, dw...)
	}

	//fmt.Println(allWires)
	//fmt.Println(len(allWires))
	//return removeDuplicate(allWires), len(wires)
	//fmt.Println(len(wires))
	return allWires, len(wires)

}

func removeDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func getDependantWires(expressions map[string]Expr, wire string) []string {
	e := expressions[wire]

	if e.op == "X" {
		return nil
	}

	if e.op == "" {
		return []string{}
	}

	tmpOp := e.op
	e.op = "X"
	expressions[wire] = e

	e1 := getDependantWires(expressions, e.a)
	e2 := getDependantWires(expressions, e.b)

	e.op = tmpOp
	expressions[wire] = e

	if e1 == nil || e2 == nil {
		return nil
	}

	var wires []string
	wires = append(wires, e1...)
	wires = append(wires, e2...)
	wires = append(wires, wire)

	return wires
}

func part2(expressions map[string]Expr) {
	// Calculate all
	var allKeys []string
	for key, exp := range expressions {
		if exp.value == -1 {
			allKeys = append(allKeys, key)
		}
	}

	// Cache all wires where key is included
	keyIncluded := map[string][]string{}
	for key, exp := range expressions {
		if exp.value != -1 {
			continue
		}
		if _, ok := keyIncluded[exp.a]; ok {
			keyIncluded[exp.a] = append(keyIncluded[exp.a], key)
		} else {
			keyIncluded[exp.a] = []string{key}
		}
		if _, ok := keyIncluded[exp.b]; ok {
			keyIncluded[exp.b] = append(keyIncluded[exp.b], key)
		} else {
			keyIncluded[exp.b] = []string{key}
		}
	}

	// Get all z keys where there is an error
	allKeys, _ = getErrorWires(expressions)

	fmt.Println(len(allKeys))

	for s11 := 0; s11 < len(allKeys)-1; s11++ {
		for s12 := s11 + 1; s12 < len(allKeys); s12++ {
			fmt.Println(s12)
			// We have first swap
			for s21 := s12 + 1; s21 < len(allKeys)-1; s21++ {
				for s22 := s21 + 1; s22 < len(allKeys); s22++ {
					fmt.Println(s21, s22)
					// Now we have two swaps
					for s31 := s22 + 1; s31 < len(allKeys)-1; s31++ {
						for s32 := s31 + 1; s32 < len(allKeys); s32++ {
							fmt.Println(s11, s12, s21, s22, s31, s32)
							// We have three swaps
							for s41 := s32 + 1; s41 < len(allKeys)-1; s41++ {
								for s42 := s41 + 1; s42 < len(allKeys); s42++ {
									// We have four swaps

									swapMap = map[string]string{}
									e11, e12 := allKeys[s11], allKeys[s12]
									e21, e22 := allKeys[s21], allKeys[s22]
									e31, e32 := allKeys[s31], allKeys[s32]
									e41, e42 := allKeys[s41], allKeys[s42]
									swapWire(e11, e12)
									swapWire(e21, e22)
									swapWire(e31, e32)
									swapWire(e41, e42)
									// Set flag to recalculate
									setRecalculateFlag(expressions, keyIncluded, []string{e11, e12, e21, e22, e31, e32, e41, e42})
									//fmt.Println("Swapping ", e11, e12, e21, e22)
									valid := checkValid(expressions)
									if valid {
										//fmt.Println(e)
										fmt.Println(e11, e12, e21, e22, e31, e32, e41, e42)
									}
									setRecalculateFlag(expressions, keyIncluded, []string{e11, e12, e21, e22, e31, e32, e41, e42})
								}
							}
						}
					}

				}
			}

		}
	}

	//swap++
	//}

	// CHeck if ok
	//fmt.Println(checkValid(expressions))
}

func setRecalculateFlag(expressions map[string]Expr, keyIncluded map[string][]string, s []string) {
	for _, el := range s {
		for _, e := range keyIncluded[el] {
			tmp := expressions[e]
			tmp.value = -1
			expressions[e] = tmp
		}
	}
}

type SwapCache struct {
	expressions  map[string]Expr
	swaps        []string
	errors       int
	nextFrontier []string
}

func part2Smart(expressions map[string]Expr) {
	allKeys, numErrors := getErrorWires(expressions)

	/*allKeys = []string{}
	for key, exp := range expressions {
		if exp.op != "" {
			allKeys = append(allKeys, key)
		}
	}*/

	// Cache all wires where key is included
	/*keyIncluded := map[string][]string{}
	for key, exp := range expressions {
		if exp.value != -1 {
			continue
		}
		if _, ok := keyIncluded[exp.a]; ok {
			keyIncluded[exp.a] = append(keyIncluded[exp.a], key)
		} else {
			keyIncluded[exp.a] = []string{key}
		}
		if _, ok := keyIncluded[exp.b]; ok {
			keyIncluded[exp.b] = append(keyIncluded[exp.b], key)
		} else {
			keyIncluded[exp.b] = []string{key}
		}
	}*/

	// Find pairs that fix stuff...
	//fmt.Println("sdf")
	//dfsCheck(expressions, allKeys, []int{}, 0)

	nextIter := []SwapCache{}
	minErr := 99999
	for i := 0; i < len(allKeys)-1; i++ {
		for j := i + 1; j < len(allKeys); j++ {

			//for k := i + 1; k < len(allKeys); k++ {
			//	if k == j {
			//		continue
			//	}
			//	for l := k + 1; l < len(allKeys); l++ {
			//		if l == j {
			//			continue
			//		}

			//fmt.Println("TESTING", allKeys[i], allKeys[j], allKeys[k], allKeys[l])

			e := copy(expressions)
			swapWireInPlace(e, allKeys[i], allKeys[j])
			//swapWireInPlace(e, allKeys[k], allKeys[l])
			//swapMap = map[string]string{}
			//swapWire(allKeys[i], allKeys[j])
			//setRecalculateFlag(expressions, keyIncluded, []string{allKeys[i], allKeys[j]})
			nextFrontier, err := getErrorWires(e)
			//fmt.Println(err, numErrors)
			if err < numErrors {
				if err < minErr {
					minErr = err
				}
				fmt.Println(err, numErrors, len(nextFrontier))
				fmt.Println(allKeys[i], allKeys[j]) //, allKeys[k], allKeys[l])
				nextIter = append(nextIter, SwapCache{e, []string{allKeys[i], allKeys[j]}, err, nextFrontier})
			}
			//setRecalculateFlag(expressions, keyIncluded, []string{allKeys[i], allKeys[j]})
			//}
			//}
		}
	}

	fmt.Println("SECOND PASS", minErr)

	for nii, ni := range nextIter {
		fmt.Println(nii, len(nextIter))
		for i := 0; i < len(ni.nextFrontier)-1; i++ {
			for j := i + 1; j < len(ni.nextFrontier); j++ {

				if slices.Contains(ni.swaps, ni.nextFrontier[i]) || slices.Contains(ni.swaps, ni.nextFrontier[j]) {
					continue
				}

				//for k := i + 1; k < len(allKeys); k++ {
				//	if k == j {
				//		continue
				//	}
				//	for l := k + 1; l < len(allKeys); l++ {
				//		if l == j {
				//			continue
				//		}

				//fmt.Println("TESTING", allKeys[i], allKeys[j], allKeys[k], allKeys[l])

				e := copy(ni.expressions)
				swapWireInPlace(e, allKeys[i], allKeys[j])
				//swapWireInPlace(e, allKeys[k], allKeys[l])
				//swapMap = map[string]string{}
				//swapWire(allKeys[i], allKeys[j])
				//setRecalculateFlag(expressions, keyIncluded, []string{allKeys[i], allKeys[j]})
				nextFrontier, err := getErrorWires(e)
				//fmt.Println(err, numErrors)
				if err < ni.errors {
					if err < minErr {
						fmt.Println(err)
						minErr = err
					}
					//fmt.Println(err, numErrors, len(nextFrontier))
					//fmt.Println(allKeys[i], allKeys[j]) //, allKeys[k], allKeys[l])
					nextIter = append(nextIter, SwapCache{e, []string{ni.swaps[0], ni.swaps[1], ni.nextFrontier[i], ni.nextFrontier[j]}, err, nextFrontier})
				}
				//setRecalculateFlag(expressions, keyIncluded, []string{allKeys[i], allKeys[j]})
				//}
				//}
			}
		}
	}

	fmt.Println(minErr)

	fmt.Println("END")

}

func dfsCheck(exp map[string]Expr, allKeys []string, swaps []int, depth int) {
	if depth == 4 {
		e := copy(exp)
		for i := 0; i < len(swaps); i += 2 {
			swapWireInPlace(e, allKeys[i], allKeys[i+1])
		}
		_, err := getErrorWires(e)
		if err != 999999 && err < 25 {
			fmt.Println(err)
		}

		fmt.Println(swaps)
		if err == 0 {
			fmt.Println(err)
			fmt.Println("FOUND SWAPS:", swaps)
		}
		return
	}

	for i := 0; i < len(allKeys)-1; i++ {
		if slices.Contains(swaps, i) {
			continue
		}
		for j := i + 1; j < len(allKeys); j++ {
			if slices.Contains(swaps, j) {
				continue
			}
			dfsCheck(exp, allKeys, append(swaps, []int{i, j}...), depth+1)
		}
	}

}

func Solve() {
	expressions := readInput("solutions/24/input.txt")
	//fmt.Println(expressions)
	//part1(expressions)
	//part2Smart(expressions)
	//part2astar(expressions)
	part2ByHand(expressions)
}

func setValue(e map[string]Expr, s string, val int) {
	el := e[s]
	el.value = val
	e[s] = el
}

func part2ByHand(exppressions map[string]Expr) {

	e := copy(exppressions)

	swapWireInPlace(e, "z07", "nqk")
	swapWireInPlace(e, "fgt", "pcp") // Za tole nism zihr...
	swapWireInPlace(e, "z24", "fpq")
	swapWireInPlace(e, "z32", "srn")

	arr := []string{"z07", "nqk", "fgt", "pcp", "z24", "fpq", "z32", "srn"}
	sort.Strings(arr)
	fmt.Println("Final result:", strings.Join(arr, ","))

	setValue(e, "x00", 0)
	setValue(e, "x01", 0)

	//x, y, z := 0, 0, 0
	//maxShift := 0
	carry := 0
	for i := 0; i <= 44; i++ {
		var xStr, yStr, zStr string
		if i < 10 {
			xStr = "x0" + strconv.Itoa(i)
			yStr = "y0" + strconv.Itoa(i)
			zStr = "z0" + strconv.Itoa(i)
		} else {
			xStr = "x" + strconv.Itoa(i)
			yStr = "y" + strconv.Itoa(i)
			zStr = "z" + strconv.Itoa(i)
		}

		xVal, yVal, zVal := calculate(e, xStr), calculate(e, yStr), calculate(e, zStr)
		fmt.Println(xVal, yVal, zVal, e[zStr], carry)
		if xVal^yVal^carry != zVal {
			fmt.Println("Error", i)

			//fmt.Println(getDependantWires(e, zStr))

		}

		if xVal+yVal == 0 {
			carry = 0
		} else if xVal+yVal > 1 {
			carry = 1
		}

		//carry = xVal ^ yVal ^ carry ^ 1

	}

}

func part2astar(expressions map[string]Expr) {
	allKeys, numErrors := getErrorWires(expressions)

	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		el:       SwapCache{expressions, []string{}, numErrors, allKeys},
		priority: 0,
		index:    0,
	}

	heap.Init(&pq)

	//memo = map[Pos]int{}

	for pq.Len() > 0 {
		el := heap.Pop(&pq).(*Item)

		/*if _, ok := memo[el.pos]; ok {
			continue
		}
		memo[el.pos] = 1*/

		if len(el.el.swaps) == 8 {
			if el.el.errors == 0 {
				fmt.Println("FOUND!!!", el.el.swaps)
			}
			continue
		}

		fmt.Println(el.el.errors, el.el.swaps)
		frontier := el.el.nextFrontier
		for i := 0; i < len(frontier)-1; i++ {
			if slices.Contains(el.el.swaps, el.el.nextFrontier[i]) {
				continue
			}
			for j := i + 1; j < len(frontier); j++ {
				if slices.Contains(el.el.swaps, el.el.nextFrontier[j]) {
					continue
				}
				e := copy(el.el.expressions)
				swapWireInPlace(e, frontier[i], frontier[j])
				nextFrontier, err := getErrorWires(e)

				if err < el.el.errors {
					heap.Push(&pq, &Item{
						el:       SwapCache{e, append(el.el.swaps, []string{el.el.nextFrontier[i], el.el.nextFrontier[j]}...), err, nextFrontier},
						priority: err,
					})
				}

			}
		}

	}
}

type Item struct {
	el       SwapCache
	priority int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
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

// SWAP 2
/*
for s11 := 0; s11 < len(allKeys)-1; s11++ {
		for s12 := s11 + 1; s12 < len(allKeys); s12++ {
			fmt.Println(s12)
			// We have first swap
			for s21 := s11 + 1; s21 < len(allKeys)-1; s21++ {
				for s22 := s21 + 1; s22 < len(allKeys); s22++ {
					if s11 == s21 || s12 == s22 {
						continue
					}
					swapMap = map[string]string{}
					// Now we have two swaps
					//e := copy(expressions) // TOO EXPENDIVE!!!
					e11, e12 := allKeys[s11], allKeys[s12]
					e21, e22 := allKeys[s21], allKeys[s22]
					swapWire(e11, e12)
					swapWire(e21, e22)
					// Set flag to recalculate
					setRecalculateFlag(expressions, keyIncluded, []string{e11, e12, e21, e22})
					//fmt.Println("Swapping ", e11, e12, e21, e22)
					valid := checkValid(expressions)
					if valid {
						//fmt.Println(e)
						fmt.Println(e11, e12, e21, e22)
					}
					setRecalculateFlag(expressions, keyIncluded, []string{e11, e12, e21, e22})
				}
			}

		}
	}
*/
