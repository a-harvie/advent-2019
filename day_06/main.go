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
			[]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"},
			[]string{"42"},
		},
	}

	for _, test := range tests {
		fmt.Printf("Test %v\nExpects %v\n", test[0], test[1])
		tree := buildTree(test[0])
		printTree(tree)
		fmt.Printf("Actual %v\n", treeTotals(tree))
		// descendant := "I"
		// anscestor := tree
		// fmt.Printf("%v is descendant of %v?: %v\n", descendant, anscestor.name, nodeHasDescendant(anscestor, descendant))
		// fmt.Printf("Path length from %v to %v: %v\n", anscestor.name, descendant, pathLengthToDescendant(anscestor, descendant))
		origin := getNodeFromTree(tree, "YOU").parent
		destination := getNodeFromTree(tree, "SAN").parent
		fmt.Printf("Transfers from %v to %v: %v\n", origin.name, destination.name, getOrbitTransfers(origin, destination.name))

		anscestor2 := "B"
		descendant2 := tree.children[0].children[0].children[0].children[1]
		fmt.Printf("Path length of %v to %v is %v\n", descendant2.name, anscestor2, pathLengthToAncestor(descendant2, anscestor2))
	}

	input, err := getInput()
	// fmt.Println(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tree := buildTree(input)
	fmt.Printf("Day 06 pt 1: %v\n", treeTotals(tree))
	origin := getNodeFromTree(tree, "YOU").parent
	destination := getNodeFromTree(tree, "SAN").parent
	fmt.Printf("Day 06 pt 2: %v\n", getOrbitTransfers(origin, destination.name))

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

func buildTree(orbitMap []string) *treeNode {
	root := makeNode(orbitMap, "COM", nil)

	return root
}

func makeNode(orbitMap []string, name string, parent *treeNode) *treeNode {
	// fmt.Printf("makeNode %v %v %v\n", len(orbitMap), name, parent)
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

	if len(unmatched) > 0 && len(node.children) > 0 {

		for i, child := range node.children {
			node.children[i] = makeNode(unmatched, child.name, &node)
		}
	}

	return &node
}

func getOrbitTransfers(origin *treeNode, destination string) int {
	if origin.name == destination {
		return 0
	}

	if nodeHasDescendant(origin, destination) {
		return pathLengthToDescendant(origin, destination)
	} else if origin.parent != nil {
		p := origin.parent
		count := 0
		for {
			if p == nil {
				break
			}
			count++
			if nodeHasDescendant(p, destination) {
				return pathLengthToDescendant(p, destination) + count
			}
			p = p.parent
		}
	}

	return 0
}

func pathLengthToAncestor(node *treeNode, anscestor string) int {
	count := 0
	if node.name == anscestor {
		return count
	}
	if node.parent == nil {
		return count
	}
	p := node.parent
	for {
		count++
		if p.name == anscestor {
			return count
		}
		if p.parent == nil {
			return 0
		} else {
			p = p.parent
		}
	}
}
func pathLengthToDescendant(node *treeNode, descendant string) int {
	return pltd(node, descendant, 0) - 1
}
func pltd(node *treeNode, descendant string, length int) int {
	if node.name == descendant {
		return 1
	}
	if len(node.children) == 0 {
		return 0
	}
	for _, child := range node.children {
		childLength := pltd(child, descendant, length)
		if childLength > 0 {
			return childLength + 1
		}
	}
	return 0
}

func nodeHasDescendant(node *treeNode, search string) bool {
	if node.name == search {
		return true
	}

	if len(node.children) == 0 {
		return false
	}

	found := false
	for _, child := range node.children {
		found = found || nodeHasDescendant(child, search)
	}

	return found
}

func getNodeFromTree(node *treeNode, search string) *treeNode {
	if node.name == search {
		return node
	}

	if len(node.children) == 0 {
		return nil
	}

	for _, child := range node.children {
		found := getNodeFromTree(child, search)
		if found != nil {
			return found
		}
	}

	return nil
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
