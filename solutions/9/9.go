package solutions_day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(filename string) string {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return line
}

func part1(diskMap []int) []int {
	var compact []int
	idCount := 0
	idCountRight := len(diskMap) / 2

	left := 0
	right := len(diskMap) - 1

	for left < right {
		// Left pointer to free space
		for i := 0; i < diskMap[left]; i++ {
			compact = append(compact, idCount)
		}
		idCount++
		left++

		// Left pointer points to free space - add content from right pointer
		for diskMap[right] > 0 && diskMap[left] > 0 {
			// If right still has content and left is still free, then MOVE
			diskMap[right]--
			diskMap[left]--
			compact = append(compact, idCountRight)
		}

		// If loop ended bacuse lack of content, move the right pointer to the left
		for diskMap[right] == 0 {
			right -= 2 // TODO: not sure, we skip free space... i think it should be ok though
			idCountRight--

			if left >= right {
				break
			}

			// TODO: make nicer!!!
			for diskMap[right] > 0 && diskMap[left] > 0 {
				// If right still has content and left is still free, then MOVE
				diskMap[right]--
				diskMap[left]--
				compact = append(compact, idCountRight)
			}
		}

		// If lopp eneded because lack of free space on the left, continue normally
		left++
	}

	if left == right {
		// On same file, add if any parts left
		for i := 0; i < diskMap[left]; i++ {
			compact = append(compact, idCount)
		}
	}

	return compact
}

type DiskPart struct {
	size int
	id   int
}

func part2(diskMap []int) []int {
	var compact []int
	idCount := 0

	right := len(diskMap) - 1

	// Make an easier representation to work with
	var diskPartMap []DiskPart
	for i, el := range diskMap {
		if i%2 != 0 {
			// Free space
			diskPartMap = append(diskPartMap, DiskPart{el, -1})
		} else {
			// A file
			diskPartMap = append(diskPartMap, DiskPart{el, idCount})
			idCount++
		}
	}

	//fmt.Println(diskMap)
	//fmt.Println(diskPartMap)

	for right > 0 {
		// Try to move the file at index 'right'
		file := diskPartMap[right]
		if file.id == -1 {
			right--
			continue
		}

		//fmt.Println("Moving file ", file.id)

		for i := 0; i < right; i++ {
			// Go from left to right and try to find a place for this file
			if diskPartMap[i].id == -1 && diskPartMap[i].size >= file.size {
				// Found a suitable spot. Put file here
				remainingSpace := diskPartMap[i].size - file.size
				diskPartMap[i] = file
				//fmt.Println(file)
				diskPartMap[right] = DiskPart{file.size, -1}
				if remainingSpace > 0 {
					// a = append(a[:i], append(b, a[i:]...)...)
					diskPartMap = insertAt(diskPartMap, DiskPart{remainingSpace, -1}, i)
				}

				break
			}
		}

		right--
	}

	// Make a 'compact' array
	for _, el := range diskPartMap {
		id := 0
		if el.id != -1 {
			id = el.id
		}
		for i := 0; i < el.size; i++ {
			compact = append(compact, id)
		}
	}

	//fmt.Println(diskPartMap)

	return compact
}

func insertAt(a []DiskPart, x DiskPart, i int) []DiskPart {
	return append(a[:i+1], append([]DiskPart{x}, a[i+1:]...)...)
}

func getIntArray(s string) []int {
	var diskMap []int
	for _, el := range s {
		diskMap = append(diskMap, int(el-'0'))
	}
	return diskMap
}

func Solve() {
	diskMapString := readInput("solutions/9/input.txt")

	compact1 := part1(getIntArray(diskMapString))
	compact2 := part2(getIntArray(diskMapString))

	sum1 := 0
	for i := 0; i < len(compact1); i++ {
		sum1 += (i * compact1[i])
	}

	sum2 := 0
	for i := 0; i < len(compact2); i++ {
		sum2 += (i * compact2[i])
	}

	fmt.Println(sum1)
	fmt.Println(sum2)

}
