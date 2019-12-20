package main

import (
	"fmt"

	"github.com/a-harvie/advent-2019/input"
	"github.com/a-harvie/advent-2019/intcode"
	"github.com/a-harvie/advent-2019/position"
)

const (
	outputLength int = 3
	empty        int = 0
	wall         int = 1
	block        int = 2
	paddle       int = 3
	ball         int = 4
)

type game struct {
	program []int
	tiles   []gameTile
	input   chan int
	output  chan int
}

type gameTile struct {
	pos    position.Coord
	tileID int
}

func main() {
	program, _ := input.CommaSeparatedInts("input/input.txt")

	g := newGame(program[0])
	g.run()
	fmt.Println(g.tiles)
	fmt.Println(g.numBlockTiles())
}

func newGame(program []int) game {
	r := game{
		program: program,
		tiles:   make([]gameTile, 0),
		input:   make(chan int),
		output:  make(chan int),
	}

	return r
}
func (g *game) run() {
	errors := make(chan error)
	go intcode.ChannelCompute(g.program, g.input, g.output, errors)

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
		case err, open := <-errors:
			if !open {
				fmt.Println("Game finished - errors closed")
				return
			}
			fmt.Printf("Robot broke? %v\n", err)
		}

		if outputIndex == outputLength {
			fmt.Println("processing output", outputs)
			outputIndex = 0
			gt := gameTile{
				pos: position.Coord{
					X: outputs[0],
					Y: outputs[1],
				},
				tileID: outputs[2],
			}
			g.tiles = append(g.tiles, gt)
		}
	}
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

// func render(painted map[position.Coord]int) {
// 	minX, maxX, minY, maxY := getBounds(painted)

// 	// Y is backwards, I guess?
// 	for y := maxY; y >= minY; y-- {
// 		for x := minX; x <= maxX; x++ {
// 			pos := position.Coord{X: x, Y: y}
// 			colour, found := painted[pos]
// 			if !found {
// 				colour = black
// 			}
// 			if colour == white {
// 				fmt.Print("□")
// 			} else {
// 				fmt.Print(" ")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

// func getBounds(painted map[position.Coord]int) (int, int, int, int) {
// 	minX := math.MaxInt64
// 	minY := math.MaxInt64
// 	maxX := math.MinInt64
// 	maxY := math.MinInt64

// 	for k := range painted {
// 		if k.X > maxX {
// 			maxX = k.X
// 		}
// 		if k.X < minX {
// 			minX = k.X
// 		}
// 		if k.Y > maxY {
// 			maxY = k.Y
// 		}
// 		if k.Y < minY {
// 			minY = k.Y
// 		}
// 	}

// 	return minX, maxX, minY, maxY
// }
