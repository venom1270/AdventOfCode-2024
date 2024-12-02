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

func Solve() {
	reports := readInput("solutions/2/input.txt")
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
				break
			}

			if i+1 == len(line) {
				safeCount++
				fmt.Println(line)
				break
			}

			diff := difference(el, line[i+1])
			if diff < 1 || diff > 3 {
				dampener--
				if i <= 1 {
					// If index low, we can change isIncreasing...
					isIncreasing = 0
				}
				if i+2 >= len(line) {
					continue
				}

				// Should we remove current or next?
				if difference(el, line[i+2]) <= 3 && difference(el, line[i+2]) >= 1 && checkIncreasing(el, line[i+2], isIncreasing) {
					// remove next
					line[i+1] = el
					if el < line[i+2] {
						isIncreasing = 1
					} else {
						isIncreasing = -1
					}
				} else {
					// remove current if possible
					if i-1 >= 0 {
						if difference(line[i-1], line[i+1]) <= 3 && difference(line[i-1], line[i+1]) >= 1 && checkIncreasing(line[i-1], line[i+1], isIncreasing) {

							if line[i-1] < line[i+1] {
								isIncreasing = 1
							} else {
								isIncreasing = -1
							}

						} else if dampener >= 0 {
							break
						}
					}
				}

				continue
			}

			if isIncreasing == 0 {
				if line[i+1]-el > 0 {
					isIncreasing = 1
				} else {
					isIncreasing = -1
				}
			}
			if isIncreasing == 1 && line[i+1] < el || isIncreasing == -1 && line[i+1] > el {
				dampener--
				// Should we remove current or next?
				if i <= 1 {
					// If index low, we can change isIncreasing...
					isIncreasing = 0
				}
				if i+2 >= len(line) {
					continue
				}

				if i == 1 {
					// there is a chance we may have to remove previous el...
					if checkIncreasing(el, line[i+1], isIncreasing) {
						if el < line[i+1] {
							isIncreasing = 1
						} else {
							isIncreasing = -1
						}
						continue
					}
				}

				if difference(el, line[i+2]) <= 3 && difference(el, line[i+2]) >= 1 && checkIncreasing(el, line[i+2], isIncreasing) {
					// remove next
					if el < line[i+2] {
						isIncreasing = 1
					} else {
						isIncreasing = -1
					}
					line[i+1] = el
				} else {
					// remove current if possible
					if i-1 >= 0 {
						if difference(line[i-1], line[i+1]) <= 3 && difference(line[i-1], line[i+1]) >= 1 && checkIncreasing(line[i-1], line[i+1], isIncreasing) {

							if line[i-1] < line[i+1] {
								isIncreasing = 1
							} else {
								isIncreasing = -1
							}

						} else if dampener >= 0 {
							break
						}
					}
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
// CORRECT 348
*/
