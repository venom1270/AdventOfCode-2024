package solutions_day2

/*
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
		fmt.Println(line)
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

func solvr(report []int, i int, isIncreasing int, dampener int) bool {
	if dampener <= 0 {
		return false
	}

	if i+1 == len(report) {
		return true
	}

	if isIncreasing == 0 {
		if report[i] < report[i+1] {
			isIncreasing = 1
		} else {
			isIncreasing = -1
		}
	}

	diff := difference(report[i], report[i+1])
	if diff <= 3 && diff > +1 {
		if isIncreasing == 1 && report[i] < report[i+1] || isIncreasing == -1 && report[i] > report[i+1] {
			// OK...
			solvr(report, i+1, isIncreasing, dampener)
		} else {
			// Wrong
			// Remove this or next
		}
	} else {
		// remove this or next

	}

	return true
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

func Solve() {
	reports := readInput("solutions/2/t.txt")
	fmt.Println(reports)

	// Part 1
	safeCount := 0
	for _, line := range reports {
		isIncreasing := 0
		for i, el := range line {
			if i+1 == len(line) {
				safeCount++
				break
			}

			diff := difference(el, line[i+1])
			if diff < 1 || diff > 3 {
				break
			}

			if isIncreasing == 0 {
				if line[i+1]-el > 0 {
					isIncreasing = 1
				} else {
					isIncreasing = -1
				}
			}
			if isIncreasing == 1 && line[i+1] < el || isIncreasing == -1 && line[i+1] > el {
				break
			}
		}
	}

	fmt.Println(safeCount)

	// Part 2 -- dampener
	safeCount = 0
	for _, line := range reports {
		isIncreasing := 0
		dampener := 1
		for i, el := range line {
			if dampener < 0 {
				//fmt.Println(line)
				break
			}

			if i+2 == len(line) {
				safeCount++
				fmt.Println(line)
				break
			}

			el2 := line[i+1]
			el3 := line[i+2]

			if isIncreasing == 0 {
				if el < el2 && el2 < el3 {
					isIncreasing = 1
				} else if el > el2 && el2 > el3 {
					isIncreasing = -1
				} else {
					dampener--
					// Check which one to remove

					if checkDifference(el, el2) {
						// Remove third
						line[i+2] = el2
						line[i+1] = el
						isIncreasing = getIncreasing(el, el2)
					} else if checkDifference(el, el3) {
						// Remove second
						line[i+1] = el
						isIncreasing = getIncreasing(el, el3)
					} else {
						// First... do nothing
						isIncreasing = getIncreasing(el2, el3)
					}

					continue
				}
			}

			if !checkDifference(el, el2) || !checkIncreasing(el, el2, isIncreasing) {
				dampener--
				// Remove first or second
				if checkDifference(el, el3) && checkIncreasing(el, el3, isIncreasing) && checkDifference(el2, el3) && checkIncreasing(el2, el3, isIncreasing) {
					// Remove first...
				} else if checkDifference(el, el3) && checkIncreasing(el, el3, isIncreasing) {
					// Remove second
					line[i+1] = el
				} else if checkDifference(el2, el3) && checkIncreasing(el2, el3, isIncreasing) {
					// Remove first
				} else {
					// No possibilities
					break
				}
			} else if !checkDifference(el2, el3) || !checkIncreasing(el2, el3, isIncreasing) {
				dampener--
				// Remove second or third
				if checkDifference(el, el3) && checkIncreasing(el, el3, isIncreasing) && checkDifference(el, el2) && checkIncreasing(el, el2, isIncreasing) {
					// Remove first...
					line[i+2] = el2
					line[i+1] = el
				} else if checkDifference(el, el3) && checkIncreasing(el, el3, isIncreasing) {
					// Remove second
					line[i+1] = el
				} else if checkDifference(el, el2) && checkIncreasing(el, el2, isIncreasing) {
					// Remove third
					line[i+2] = el2
					line[i+1] = el
				} else {
					// No possibilities
					break
				}
			}

		}
	}

	fmt.Println(safeCount)

}

// 337 too low
// 338 wrong
// 341 wrong
// 358 wrong
// 363
// 347, 348
// CORRECT 348
*/
