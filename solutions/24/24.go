package solutions_day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func swapWire(e1, e2 string) {
	/*tmp1 := e[e1]
	tmp2 := e[e2]
	e[e1] = e[e2]
	e[e2] = tmp*/
	swapMap[e1] = e2
	swapMap[e2] = e1
}

var swapMap map[string]string

func part2(expressions map[string]Expr) {
	// Calculate all
	var allKeys []string
	for key, exp := range expressions {
		if exp.value == -1 {
			allKeys = append(allKeys, key)
		}
	}

	//swap := 2
	//for swap <= 4 {

	/*for s1 := 0; s1 < len(allKeys)-1; s1++ {
		for s2 := s1 + 1; s2 < len(allKeys); s2++ {
			// Swap s1 and s2
			e := copy(expressions)
			e1, e2 := allKeys[s1], allKeys[s2]
			swapWire(e, e1, e2)
			valid := checkValid(e)
			if valid && checkValidForAll(e) {
				fmt.Println(e)
				fmt.Println(e1, e2)
			}
		}
	}*/

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

	fmt.Println(len(allKeys))

	for s11 := 0; s11 < len(allKeys)-1; s11++ {
		for s12 := s11 + 1; s12 < len(allKeys); s12++ {
			fmt.Println(s12)
			// We have first swap
			for s21 := s11 + 1; s21 < len(allKeys)-1; s21++ {
				if s12 == s21 {
					continue
				}
				for s22 := s21 + 1; s22 < len(allKeys); s22++ {
					if s12 == s22 {
						continue
					}
					//fmt.Println(s21, s22)
					// Now we have two swaps
					for s31 := s21 + 1; s31 < len(allKeys)-1; s31++ {
						if s31 == s22 || s31 == s12 {
							continue
						}
						for s32 := s31 + 1; s32 < len(allKeys); s32++ {
							if s32 == s22 || s32 == s12 {
								continue
							}
							fmt.Println(s21, s22, s31, s32)
							// We have three swaps
							for s41 := s31 + 1; s41 < len(allKeys)-1; s41++ {
								if s41 == s32 || s41 == s22 || s41 == s12 {
									continue
								}
								for s42 := s41 + 1; s42 < len(allKeys); s42++ {
									if s42 == s32 || s42 == s22 || s42 == s12 {
										continue
									}
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

func Solve() {
	expressions := readInput("solutions/24/input.txt")
	fmt.Println(expressions)
	//part1(expressions)
	part2(expressions)
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
