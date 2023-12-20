package main

import (
	h "day8/helpers"
	"day8/models"
	"fmt"
	"regexp"
	"sync"
)

func loadInstructions(fileName string) *[]string {
	return h.LoadFile[[]string](fileName)
}

func loadInput(fileName string) *[]models.Node {
	return h.LoadFile[[]models.Node](fileName)
}

func resolveNodes(nodes *[]models.Node) {
	var wg sync.WaitGroup
	for i := 0; i < len(*nodes); i++ {
		wg.Add(1)
		go func(node *models.Node) {
			defer wg.Done()
			for j := 0; j < len(*nodes); j++ {
				leftOrRight := &(*nodes)[j]
				if (*node).Left == (*leftOrRight).Id {
					(*node).LeftNode = leftOrRight
				}
				if (*node).Right == (*leftOrRight).Id {
					(*node).RightNode = leftOrRight
				}
				if (*node).LeftNode != nil && (*node).RightNode != nil {
					break
				}
			}
		}(&(*nodes)[i])
	}
	wg.Wait()
}

func getFirstNode(nodes *[]models.Node) *models.Node {
	for _, n := range *nodes {
		if n.Id == "AAA" {
			return &n
		}
	}

	return nil
}

func findStartingNodes(nodes *[]models.Node) *[]models.Node {
	starters := make([]models.Node, 0, 1000)
	re := regexp.MustCompile(`..A`)
	for _, n := range *nodes {
		if re.MatchString(n.Id) {
			starters = append(starters, n)
		}
	}

	return &starters
}

func findSteps(instructions []string, nodes *[]models.Node) int {
	i, counter := 0, 0
	node := getFirstNode(nodes)
	for {
		counter++
		node = h.Ternary(instructions[i] == "L", node.LeftNode, node.RightNode)
		if (*node).Id == "ZZZ" {
			break
		}
		i++
		if i == len(instructions) {
			i = 0
		}
	}
	return counter
}

func isThisTheEnd(nodes *[]models.Node) bool {
	re := regexp.MustCompile(`..Z`)
	fmt.Println(nodes)
	for i := 0; i < len(*nodes); i++ {
		if !(re.MatchString((*nodes)[i].Id)) {
			return false
		}
	}
	return true
}

// func findConcurrentSteps(instructions []string, ns *[]models.Node) int {
// 	i, counter := 0, 0
// 	nodes := *ns
// 	for {
// 		counter++
// 		for j := 0; j < len(nodes); j++ {
// 			nodes[j] = *h.Ternary(instructions[i] == "L", nodes[j].LeftNode, nodes[j].RightNode)
// 		}
// 		if isThisTheEnd(&nodes) {
// 			break
// 		}
// 		i++
// 		if i == len(instructions) {
// 			i = 0
// 		}
// 	}
// 	return counter
// }

func findConcurrentSteps(instructions []string, starterNodes *[]models.Node) []int {
	var steps []int
	var wg sync.WaitGroup
	re := regexp.MustCompile(`..Z`)

	for _, n := range *starterNodes {
		wg.Add(1)
		go func(startNode models.Node, currentSteps *[]int) {
			defer wg.Done()
			i, counter := 0, 0
			node := &startNode
			for {
				counter++
				node = h.Ternary(instructions[i] == "L", node.LeftNode, node.RightNode)
				if re.MatchString((*node).Id) {
					break
				}
				i++
				if i == len(instructions) {
					i = 0
				}
			}

			steps = append(steps, counter)
		}(n, &steps)
	}
	wg.Wait()
	return steps
}

func main() {
	instructions := loadInstructions("instructions.json")
	nodes := loadInput("input.json")

	resolveNodes(nodes)

	steps1 := findSteps(*instructions, nodes)

	fmt.Println("Part 1 Steps:", steps1)

	starters := findStartingNodes(nodes)

	steps2 := findConcurrentSteps(*instructions, starters)

	a, b, integers := steps2[0:1][0], steps2[1:2][0], steps2[2:]

	fmt.Println("Part 2 Steps:", h.FindLCM(a, b, integers...))
}
