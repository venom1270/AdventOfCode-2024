package solutions_day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func readInput(filename string) [][]string {
	fmt.Println("Reading input...")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var connections [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "-")
		connections = append(connections, []string{s[0], s[1]})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return connections
}

type Graph map[string][]string

func makeGraph(connections [][]string) Graph {
	g := Graph{}

	for _, conn := range connections {
		if n, ok := g[conn[0]]; ok {
			g[conn[0]] = append(n, conn[1])
		} else {
			g[conn[0]] = []string{conn[1]}
		}
		if n, ok := g[conn[1]]; ok {
			g[conn[1]] = append(n, conn[0])
		} else {
			g[conn[1]] = []string{conn[0]}
		}
	}

	return g
}

func part1(g Graph) {
	var trios [][]string
	for key, connections := range g {
		for _, secondConn := range connections {
			for _, thirdConn := range g[secondConn] {
				if thirdConn != key {
					for _, firstConn := range connections {
						if thirdConn == firstConn {
							// Check for duplicates
							add := true
							for _, trio := range trios {
								t1, t2, t3 := trio[0], trio[1], trio[2]
								if t1 == key && t2 == secondConn && t3 == thirdConn || t1 == secondConn && t2 == thirdConn && t3 == key || t1 == thirdConn && t2 == key && t3 == secondConn || t1 == key && t2 == thirdConn && t3 == secondConn || t1 == secondConn && t2 == key && t3 == thirdConn || t1 == thirdConn && t2 == secondConn && t3 == key {
									add = false
									break
								}
							}
							if add {
								trios = append(trios, []string{key, secondConn, thirdConn})
							}
							break
						}
					}
				}
			}
		}
	}

	fmt.Println("Num trios: ", len(trios))

	sum := 0
	for _, trio := range trios {
		for _, t := range trio {
			if t[0] == 't' {
				sum++
				break
			}
		}
	}

	fmt.Println("Part 1:", sum)
}

var visited map[string]bool

func dfs(g Graph, nodes []string, n string) []string {

	sort.Strings(nodes)
	key := strings.Join(nodes, "")
	if _, ok := visited[key]; ok {
		return []string{}
	} else {
		visited[key] = true
	}

	var groups [][]string
	nodes = append(nodes, n)
	for _, conn := range g[n] {
		if areConnected(g, nodes[:len(nodes)-1], conn) {
			//fmt.Println("ARE CONNECTED")
			group := dfs(g, copy(nodes), conn)
			if len(group) > 0 {
				groups = append(groups, group)
			}
		} else {
			groups = append(groups, copy(nodes))
			//fmt.Println("NOT cONNECTED")
		}
	}

	// Go through groups and return longest
	//fmt.Println(groups)
	longestGroup := []string{}
	for _, group := range groups {
		if len(group) > len(longestGroup) {
			longestGroup = group
		}
	}

	return longestGroup
}

func part2(g Graph) {

	visited = map[string]bool{}
	longestGroup := []string{}
	for firstConn, connections := range g {
		for _, n := range connections {
			lg := dfs(g, []string{firstConn}, n)
			if len(lg) > len(longestGroup) {
				longestGroup = lg
			}
		}
	}

	//fmt.Println(longestGroup)
	sort.Strings(longestGroup)
	//fmt.Println(longestGroup)
	fmt.Println("Part 2:", strings.Join(longestGroup, ","))

}

func copy(x []string) []string {
	var c []string
	for _, el := range x {
		c = append(c, el)
	}
	return c
}

func areConnected(g Graph, a []string, b string) bool {
	for _, n := range a {
		if !isConnected(g, n, b) {
			return false
		}
	}
	return true
}

func isConnected(g Graph, a, b string) bool {
	if a == b {
		return false // Special case
	}
	for _, c := range g[a] {
		if c == b {
			return true
		}
	}
	return false
}

func Solve() {
	connections := readInput("solutions/23/input.txt")
	//fmt.Println(connections)
	g := makeGraph(connections)
	//fmt.Println(g)
	part1(g) // Part 1 is a bit unoptipal, but we get an answer in about 5 seconds so it's fine :)
	part2(g)
}
