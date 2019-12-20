package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"

	"github.com/a-harvie/advent-2019/input"
	"github.com/a-harvie/advent-2019/intcode"
	"github.com/a-harvie/advent-2019/position"
)

const (
	outputLength      int = 3
	empty             int = 0
	wall              int = 1
	block             int = 2
	paddle            int = 3
	ball              int = 4
	numQuartersToPlay int = 2
)

type game struct {
	program        []int
	tiles          tileSet
	inputReady     chan bool
	input          chan int
	output         chan int
	ballPosition   position.Coord
	paddlePosition position.Coord
	score          int
}

type gameTile struct {
	pos    position.Coord
	tileID int
}

type tileSet map[position.Coord]gameTile

func main() {
	program, _ := input.CommaSeparatedInts("input/input.txt")

	g := newGame(program[0])
	g.run()
	fmt.Printf("Day 13 pt 1: %v\n", g.numBlockTiles())

	g = newGame(program[0])
	g.runWithQuarters()
	fmt.Printf("Day 13 pt 2: %v\n", g.score)
}

func clearScreen() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func newGame(program []int) game {
	r := game{
		program:    program,
		tiles:      make(tileSet, 0),
		input:      make(chan int),
		inputReady: make(chan bool),
		output:     make(chan int),
	}

	return r
}

func (g *game) runWithQuarters() {
	g.program[0] = numQuartersToPlay
	g.run()
}

func (g *game) run() {
	quarters := g.program[0]
	fmt.Printf("Running with quarters: %v\n", quarters)
	errors := make(chan error)
	go intcode.ChannelComputeV2(g.program, g.input, g.inputReady, g.output, errors)

	var outputs [outputLength]int
	outputIndex := 0
	for {
		select {
		case x, open := <-g.output:
			if !open {
				fmt.Println("Game finished - outputs closed")
				return
			}
			outputs[outputIndex] = x
			outputIndex++
			if outputIndex == outputLength {
				outputIndex = 0
				// fmt.Printf("Outputs %v\n", outputs)
				g.handleOutput(outputs)

				// TODO: render less glitchily?
				// clearScreen()
				// time.Sleep(50 * time.Millisecond) // this doesn't work at all
				// g.render()
			}

		case err, open := <-errors:
			if !open {
				fmt.Println("Game finished - errors closed")
				return
			}
			fmt.Printf("Game broke? %v\n", err)
		case _, open := <-g.inputReady:
			if !open {
				fmt.Println("Game finished - inputReady closed")
				return
			}
			nextInput := g.getNextInput()
			g.input <- nextInput

		}
	}
}

func (g *game) handleOutput(outputs [3]int) {
	// fmt.Printf("handleOutput %v\n", outputs)
	x := outputs[0]
	y := outputs[1]

	if x == -1 && y == 0 {
		// fmt.Printf("Set score to %v\n", outputs[2])
		g.score = outputs[2]
		return
	}

	obj := outputs[2]

	gt := gameTile{
		pos: position.Coord{
			X: x,
			Y: y,
		},
		tileID: obj,
	}
	g.tiles[gt.pos] = gt

	switch obj {
	case ball:
		g.ballPosition.X = x
		g.ballPosition.Y = y
	case paddle:
		g.paddlePosition.X = x
		g.paddlePosition.Y = y
	}
}

func (g *game) getNextInput() int {
	if g.ballPosition.X < g.paddlePosition.X {
		return -1
	}
	if g.ballPosition.X > g.paddlePosition.X {
		return 1
	}
	return 0
}

func (g *game) numBlockTiles() int {
	count := 0
	for _, tile := range g.tiles {
		if tile.tileID == block {
			count++
		}
	}

	return count
}

func (g *game) render() {
	fmt.Printf("Score: %v\n", g.score)
	minX, maxX, minY, maxY := g.getBounds()
	// Y is backwards, I guess?
	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			pos := position.Coord{X: x, Y: y}
			var draw int
			tile, found := g.tiles[pos]
			if !found {
				draw = empty
			} else {
				draw = tile.tileID
			}
			switch draw {
			case empty:
				fmt.Print(" ")
			case wall:
				fmt.Print("#")
			case block:
				fmt.Print("□")
			case paddle:
				fmt.Print("_")
			case ball:
				fmt.Print("•")
			}
		}
		fmt.Println()
	}
}

func (g *game) getBounds() (int, int, int, int) {
	minX := math.MaxInt64
	minY := math.MaxInt64
	maxX := math.MinInt64
	maxY := math.MinInt64

	for _, tile := range g.tiles {
		if tile.pos.X > maxX {
			maxX = tile.pos.X
		}
		if tile.pos.X < minX {
			minX = tile.pos.X
		}
		if tile.pos.Y > maxY {
			maxY = tile.pos.Y
		}
		if tile.pos.Y < minY {
			minY = tile.pos.Y
		}
	}

	return minX, maxX, minY, maxY
}

// func (r robot) run(startColour int) {
// 	// Hack to get around race condition - need a better way to signal computer is done :/
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Guess the robot is done now ¯\\_(ツ)_/¯:", r)
// 		}
// 	}()
// 	errors := make(chan error)

// 	// start the program and supply it with the starting camera input
// 	go intcode.ChannelCompute(g.program, g.camera, g.output, errors)
// 	g.camera <- startColour

// 	var outputs [2]int
// 	outputIndex := 0
// 	for {
// 		select {
// 		case x, open := <-g.output:
// 			if !open {
// 				fmt.Println("Robot finished - outputs closed")
// 				break
// 			}
// 			outputs[outputIndex] = x
// 			outputIndex++
// 		case err, open := <-errors:
// 			if !open {
// 				fmt.Println("Robot finished - errors closed")
// 				break
// 			}
// 			fmt.Printf("Robot broke? %v\n", err)
// 		}

// 		if outputIndex > 1 {
// 			outputIndex = 0
// 			g.painted[g.pos] = outputs[0]
// 			// fmt.Printf("Robot has now painted: %v\n", g.painted)
// 			if outputs[1] == 0 {
// 				// fmt.Println("turning left")
// 				g.vec.Rotate(left)
// 			} else {
// 				// fmt.Println("turning right")
// 				g.vec.Rotate(right)
// 			}
// 			g.pos.Translate(g.vec)
// 			// fmt.Printf("Robot now at: %v\n", g.pos)

// 			newColour, found := g.painted[g.pos]
// 			if !found {
// 				newColour = black
// 			}
// 			g.camera <- newColour
// 		}
// 	}
// }

//
