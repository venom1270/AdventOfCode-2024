package solutions_day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput(filename string) []int {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		numbers = append(numbers, n)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return numbers
}

func mix(secret int, number int) int {
	return secret ^ number
}

func prune(secret int) int {
	return secret % 16777216
}

func getNextSecretNumber(number int) int {
	number = mix(number, number*64)
	number = prune(number)
	number = mix(number, number/32)
	number = prune(number)
	number = mix(number, number*2048)
	number = prune(number)
	return number
}

func part1(numbers []int, iterations int) ([][]int, [][]int) {
	sum := 0
	var prices [][]int
	var changes [][]int
	for _, num := range numbers {
		var p []int
		var c []int
		prev := 0
		for range iterations {
			num = getNextSecretNumber(num)
			price := num % 10
			p = append(p, price)
			c = append(c, price-prev)
			prev = price
		}
		sum += num
		prices = append(prices, p)
		changes = append(changes, c)
	}

	fmt.Println("Part 1: ", sum)

	return prices, changes
}

type Sequence struct {
	a, b, c, d int
}

func findBestSequence(prices, changes [][]int) {

	var sequences map[Sequence][]int
	sequences = map[Sequence][]int{}

	// Loop through all the buyers to get all sequences
	for buyerI := range len(prices) {
		buyerPrices := prices[buyerI]
		buyerChanges := changes[buyerI]

		seenSequences := map[Sequence]bool{}

		// Get all sequences of this buyer and append price to list
		for j := 0; j < len(buyerChanges)-3; j++ {
			s := Sequence{buyerChanges[j], buyerChanges[j+1], buyerChanges[j+2], buyerChanges[j+3]}

			if _, ok := seenSequences[s]; ok {
				continue
			}
			seenSequences[s] = true

			if el, ok := sequences[s]; ok {
				sequences[s] = append(el, buyerPrices[j+3])
			} else {
				sequences[s] = []int{buyerPrices[j+3]}
			}
		}
	}

	// Loop through all sequences and sum up their yield - get the best one
	maxYield := 0
	//var bestS Sequence
	for _, yield := range sequences {
		sum := 0
		for i := range len(yield) {
			sum += yield[i]
		}
		if sum > maxYield {
			maxYield = sum
			//bestS = s
		}
	}

	fmt.Println("Part 2: ", maxYield)
	//fmt.Println(bestS)

}

func Solve() {
	numbers := readInput("solutions/22/input.txt")
	//fmt.Println(numbers)
	// Part 1
	p, c := part1(numbers, 2000)

	// Part 2
	findBestSequence(p, c)
}
