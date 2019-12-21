package main

import (
	"fmt"
	"strconv"

	"github.com/a-harvie/advent-2019/input"
)

type stochioList []([]string)

type treeNode struct {
	name        string
	quantity    int
	oreQuantity int
	parent      *treeNode
	children    []*treeNode
}

// maps chemical names to the ore ratios to produce them
type leafChemicals map[string][]int

func main() {
	for i := 1; i <= 5; i++ {
		testInput, err := input.StochiometricList(fmt.Sprintf("input/test_0%v.txt", i))
		if err != nil {
			fmt.Printf("Oh noes at %v: %v\n", i, err)
		}
		fmt.Println(testInput)
		tree, err := buildTree(testInput)
		if err != nil {
			fmt.Printf("Oh noes at %v: %v\n", i, err)
		}
		printTree(tree)
		leafChem := getLeafChemicals(tree, make(leafChemicals, 0))
		// fmt.Printf("Leafchem: %v\n", leafChem)
		fmt.Printf("Ore needed: %v\n", getOreNeeded(tree, leafChem))
	}
}

func buildTree(input stochioList) (*treeNode, error) {
	return makeNode(input, "FUEL", 1, nil)
}

func makeNode(list stochioList, name string, quantity int, parent *treeNode) (*treeNode, error) {
	// fmt.Printf("makeNode %v %v %v\n", len(orbitMap), name, parent)

	// find data for this node, remove it from the list
	nodeData, list, err := findNodeDataForName(list, name)
	if err != nil {
		return nil, err
	}
	node := treeNode{
		name:     name,
		quantity: quantity,
		parent:   parent,
		children: make([]*treeNode, 0),
	}

	for i := 2; i < len(nodeData); i += 2 {
		childName := nodeData[i]
		childQuantity, _ := strconv.Atoi(nodeData[i+1])
		if childName == "ORE" {
			quantity, _ := strconv.Atoi(nodeData[1])
			child := treeNode{
				name:        "ORE",
				quantity:    quantity,
				oreQuantity: childQuantity,
				parent:      &node,
			}
			node.children = append(node.children, &child)
		} else {
			child, err := makeNode(list, childName, childQuantity, &node)
			if err != nil {
				return nil, err
			}
			node.children = append(node.children, child)
		}
	}

	return &node, nil
}

func findNodeDataForName(list stochioList, name string) ([]string, stochioList, error) {
	for i, entry := range list {
		if entry[0] == name {
			return entry, append(list[:i], list[i:]...), nil
		}
	}

	return nil, nil, fmt.Errorf("%v not found in list", name)
}
func getLeafChemicals(node *treeNode, lc leafChemicals) leafChemicals {
	if node.name == "ORE" {
		_, found := lc[node.parent.name]
		if !found {
			lc[node.parent.name] = []int{node.quantity, node.oreQuantity}
		}
	} else {
		for _, child := range node.children {
			childLC := getLeafChemicals(child, lc)
			for k, v := range childLC {
				if _, found := lc[k]; !found {
					lc[k] = v
				}
			}
		}
	}
	return lc
}

func getOreNeeded(tree *treeNode, leafChem leafChemicals) int {
	total := 0
	leafChemTally := getLeafChemTally(tree, leafChem)
	for k, v := range leafChemTally {
		// fmt.Printf("%v: %v; %v\n", k, v, leafChem[k])
		min := getMinimumQuantityForRatio(v, leafChem[k])
		// fmt.Printf("Min amt of ore: %v\n", min)
		total += min
	}
	return total
}

func getMinimumQuantityForRatio(q int, ratio []int) int {
	produced := ratio[0]
	consumed := ratio[1]
	amt := produced
	for amt < q {
		amt += produced
	}
	if amt/consumed > 0 {
		return amt / produced * consumed
	} else {
		return amt
	}
}

func getLeafChemTally(tree *treeNode, leafChem leafChemicals) map[string]int {
	tally := make(map[string]int)
	for k := range leafChem {
		tally[k] = 0
	}
	tallyLeafChem(tree, tally)
	return tally
}

func tallyLeafChem(node *treeNode, tally map[string]int) {
	amt, isLeaf := tally[node.name]
	// fmt.Printf("tallyLeafChem: %v %v %v %v\n", node.name, amt, isLeaf, tally)
	if isLeaf {
		multiplier := 1
		if node.parent != nil {
			multiplier = node.parent.quantity
		}
		tally[node.name] = amt + (node.quantity * multiplier)
	} else {
		for _, child := range node.children {
			tallyLeafChem(child, tally)
		}
	}
}

func printTree(node *treeNode) {
	fmt.Printf("%v %v is produced by:", node.quantity, node.name)

	for _, child := range node.children {
		if child.name == "ORE" {
			fmt.Printf(" %v:%v ore;", child.quantity, child.oreQuantity)
		} else {
			fmt.Printf(" %v %v;", child.quantity, child.name)
		}
	}

	if node.parent != nil {
		fmt.Printf(" and produces: %v\n", node.parent.name)
	} else {
		fmt.Println(" and produces nothing.")
	}

	for _, child := range node.children {
		printTree(child)
	}
}
