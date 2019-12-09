package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	tests := []([]([]string)){
		{
			[]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"},
			[]string{"42"},
		},
	}

	for _, test := range tests {
		fmt.Printf("Test %v\nExpects %v\n", test[0], test[1])
		fmt.Printf("Actual %v\n", totalOrbits(test[0]))
	}

	input, err := getInput()
	// fmt.Println(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Day 06 pt 1: %v\n", totalOrbits(input))

}

func getInput() ([]string, error) {
	file, err := os.Open("input/input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	input := make([]string, 0)
	for {
		fileLine, err := reader.ReadString('\n')

		input = append(input, strings.Trim(fileLine, " \t\n"))
		if err != nil {
			break
		}
	}

	return input, nil
}

type treeNode struct {
	name     string
	parent   *treeNode
	children []*treeNode
}

func totalOrbits(orbitMap []string) int {
	root := buildTree(orbitMap)
	printTree(root)
	return treeTotals(root)
}

// func buildTree(orbitMap []string) treeNode {
// 	orbits := make(map[string]string)
// 	// root := treeNode{"COM", nil, make([]*treeNode, 0)}
// 	treeNodes := make(map[string]*treeNode)

// 	for _, orbitDesc := range orbitMap {
// 		parts := strings.Split(orbitDesc, ")")
// 		orbits[parts[1]] = parts[0]

// 	}
// 	for
// 	fmt.Println(orbits)
// 	return treeNode{"", nil, nil}
// }

func buildTree(orbitMap []string) *treeNode {
	root := getNode(orbitMap, "COM", nil)

	return root
}

func getNode(orbitMap []string, name string, parent *treeNode) *treeNode {
	fmt.Printf("getNode %v %v %v\n", len(orbitMap), name, parent)
	node := treeNode{name, parent, make([]*treeNode, 0)}
	unmatched := make([]string, 0)
	for _, orbitDesc := range orbitMap {
		parts := strings.Split(orbitDesc, ")")
		if parts[0] == name {
			// fmt.Printf("Matched %v %v", parts[0], name)
			child := treeNode{parts[1], &node, make([]*treeNode, 0)}
			node.children = append(node.children, &child)
		} else {
			// fmt.Printf("Nope %v %v", parts[0], name)
			unmatched = append(unmatched, orbitDesc)
		}
	}
	fmt.Printf("Umatched %v Children %v\n", len(unmatched), len(node.children))
	if len(unmatched) > 0 && len(node.children) > 0 {

		for i, child := range node.children {
			node.children[i] = getNode(unmatched, child.name, &node)
		}
	}

	return &node
}

func treeTotals(node *treeNode) int {
	total := 0

	p := node.parent
	for {
		if p == nil {
			break
		}

		p = p.parent

		total++
	}

	for _, child := range node.children {
		total += treeTotals(child)
	}

	return total
}

func printTree(node *treeNode) {
	fmt.Printf("%v is orbited by:", node.name)
	for _, child := range node.children {
		fmt.Printf(" %v", child.name)
	}
	if node.parent != nil {
		fmt.Printf(" and orbits: %v\n", node.parent.name)
	} else {
		fmt.Println(" and orbits nothing.")
	}

	for _, child := range node.children {
		printTree(child)
	}
}
