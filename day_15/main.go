package main

import (
	"fmt"

	"github.com/a-harvie/advent-2019/input"
	"github.com/a-harvie/advent-2019/intcode"
	"github.com/a-harvie/advent-2019/position"
)

const (
	// direction codes
	north int = 1
	south int = 2
	west  int = 3
	east  int = 4

	// status codes
	wall    int = 0
	moved   int = 1
	foundO2 int = 2
)

// type direction int

type droid struct {
	program []int
	// pos           position.Coord
	// visited       map[position.Coord]bool
	envMap      envMap
	envRoot     *envNode
	currentNode *envNode
	lastMove    int
	foundGoal   bool
	debug       bool
}

type envMap map[position.Coord]*envNode

type envNode struct {
	pos      position.Coord
	parent   *envNode
	children map[int]*envNode
	isWall   bool
	visited  bool
}

func main() {
	program, _ := input.CommaSeparatedInts("input/input.txt")
	d := newDroid(program[0])
	d.debug = true
	d.run()
	fmt.Printf("Distance to origin: %v\n", d.getDistanceToOrigin())
}

func newDroid(program []int) *droid {
	d := droid{
		program:  program,
		lastMove: 0,
	}
	envRoot := newEnvNode(position.Coord{X: 0, Y: 0}, nil, false)
	envRoot.visited = true
	d.envRoot = &envRoot
	d.currentNode = &envRoot
	return &d
}

func newEnvNode(pos position.Coord, parent *envNode, isWall bool) envNode {
	e := envNode{
		pos:      pos,
		parent:   parent,
		children: make(map[int]*envNode),
		isWall:   isWall,
		visited:  false,
	}
	return e
}

func (d *droid) run() {
	input := make(chan int)
	waiting := make(chan bool)
	output := make(chan int)
	errors := make(chan error)
	go intcode.ChannelComputeV2(d.program, input, waiting, output, errors)
	for {
		select {
		case _, open := <-waiting:
			d.log(fmt.Sprintln("Droid awaiting input"))
			if !open {
				fmt.Println("Droid terminated on waiting channel")
				return
			}
			nextInput := d.getNextInput()
			d.log(fmt.Sprintf("Input %v to program\n", nextInput))
			input <- nextInput
		case o, open := <-output:
			if !open {
				fmt.Println("Droid terminated on output channel")
				return
			}
			d.log(fmt.Sprintf("Got output %v\n", o))
			d.processOutput(o)
			if d.foundGoal {
				fmt.Println("Found O2! Terminating.")
				return
			}
		case err, open := <-errors:
			if !open {
				fmt.Println("Droid terminated on errors channel")
				return
			}
			fmt.Printf("Oh noes poor droid: %v\n", err)
		}

		if d.debug {
			fmt.Println("Debug mode: pausing until the enter key is pressed")
			fmt.Printf("Droid: lastMove: %v foundGoal: %v\n", d.lastMove, d.foundGoal)
			fmt.Printf("Droid is at: %v %v\n", &d.currentNode, d.currentNode)
			var input string
			fmt.Scanln(&input)
		}
	}
}

func (d *droid) log(msg string) {
	if d.debug {
		fmt.Print(msg)
	}
}

func (d *droid) getNextInput() int {
	var dir int

	dir = d.nextMoveDirection()

	return dir
}

func (d *droid) nextMoveDirection() int {
	dir := 0
	for i := north; i <= east; i++ {

		if i == d.getDirectionToPreviousNode() {
			d.log(fmt.Sprintf("Nextmovedir: droid came from %v, skipping\n", i))
			continue
		}
		child, exists := d.currentNode.children[i]
		d.log(fmt.Sprintf("Nextmovedir: %v %v %v\n", i, child, exists))

		if !exists {
			dir = i
			break
		}
		if !child.isWall && !child.visited {
			dir = i
			break
		}
	}

	if dir == 0 {
		dir = d.getDirectionToPreviousNode()
		d.log(fmt.Sprintf("Dead end: droid backtracking to %v\n", dir))
	}

	d.lastMove = dir
	return dir
}

func (d *droid) getNextCoordinateInDirection(direction int) position.Coord {
	var newCoordinate position.Coord
	currentPos := d.currentNode.pos
	switch direction {
	case north:
		newCoordinate = position.Coord{X: currentPos.X, Y: currentPos.Y - 1}
	case south:
		newCoordinate = position.Coord{X: currentPos.X, Y: currentPos.Y + 1}
	case west:
		newCoordinate = position.Coord{X: currentPos.X - 1, Y: currentPos.Y}
	case east:
		newCoordinate = position.Coord{X: currentPos.X + 1, Y: currentPos.Y}

	}
	return newCoordinate
}

func (d *droid) getDirectionToPreviousNode() int {
	parent := d.currentNode.parent
	if parent == nil {
		return 0
	}

	dir := 0
	x1 := d.currentNode.pos.X
	x2 := d.currentNode.parent.pos.X
	if x1 < x2 {
		dir = east
	} else if x1 > x2 {
		dir = west
	}

	y1 := d.currentNode.pos.Y
	y2 := d.currentNode.parent.pos.Y
	if y1 < y2 {
		dir = south
	} else if y1 > y2 {
		dir = north
	}

	return dir
}

func (d *droid) getBacktrackDirection() int {
	dir := 0
	switch d.lastMove {
	case north:
		dir = south
	case south:
		dir = north
	case west:
		dir = east
	case east:
		dir = west
	}
	return dir
}

func (d *droid) processOutput(output int) {
	switch output {
	case wall:
		d.addWall()
	case moved:
		d.updatePosition()
	case foundO2:
		d.updatePosition()
		d.foundGoal = true
	}
}

func (d *droid) addWall() {
	wallCoordinate := d.getNextCoordinateInDirection(d.lastMove)
	currentNode := d.currentNode
	wallNode := newEnvNode(wallCoordinate, currentNode, true)
	currentNode.children[d.lastMove] = &wallNode
}

func (d *droid) updatePosition() {
	newCoordinate := d.getNextCoordinateInDirection(d.lastMove)
	currentNode := d.currentNode
	newNode := newEnvNode(newCoordinate, currentNode, false)
	newNode.visited = true
	currentNode.children[d.lastMove] = &newNode
	d.currentNode = &newNode
}

func (e *envNode) hasUnvisitedChildren() bool {
	maxChildren := 3
	if e.parent == nil {
		maxChildren = 4
	}
	return len(e.children) < maxChildren
}

func (d *droid) getDistanceToOrigin() int {
	distance := 0
	current := d.currentNode
	parent := current.parent
	for parent != nil {
		distance++
		parent = parent.parent
		current = parent
	}
	return distance
}
