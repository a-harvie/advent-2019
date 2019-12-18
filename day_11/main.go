package main

import (
	"fmt"
	"math"

	"github.com/a-harvie/advent-2019/input"
	"github.com/a-harvie/advent-2019/intcode"
	"github.com/a-harvie/advent-2019/position"
)

const (
	left  float64 = 90.0
	right float64 = -90.0
	black int     = 0
	white int     = 1
)

type robot struct {
	program []int
	vec     position.Coord
	pos     position.Coord
	painted map[position.Coord]int
	camera  chan int
	output  chan int
}

func main() {
	program, _ := input.CommaSeparatedInts("input/input.txt")
	fmt.Println(program[0])

	r := newRobot(program[0])
	r.run(black)
	fmt.Println(len(r.painted))

	r2 := newRobot(program[0])
	r2.run(white)
	render(r2.painted)
}

func newRobot(program []int) robot {
	r := robot{
		program: program,
		vec:     position.Coord{X: 0, Y: 1},
		pos:     position.Coord{X: 0, Y: 0},
		painted: make(map[position.Coord]int),
		camera:  make(chan int),
		output:  make(chan int),
	}

	return r
}

func (r robot) run(startColour int) {
	// Hack to get around race condition - need a better way to signal computer is done :/
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Guess the robot is done now ¯\\_(ツ)_/¯:", r)
		}
	}()
	errors := make(chan error)

	// start the program and supply it with the starting camera input
	go intcode.ChannelCompute(r.program, r.camera, r.output, errors)
	r.camera <- startColour

	var outputs [2]int
	outputIndex := 0
	for {
		select {
		case x, open := <-r.output:
			if !open {
				fmt.Println("Robot finished - outputs closed")
				break
			}
			outputs[outputIndex] = x
			outputIndex++
		case err, open := <-errors:
			if !open {
				fmt.Println("Robot finished - errors closed")
				break
			}
			fmt.Printf("Robot broke? %v\n", err)
		}

		if outputIndex > 1 {
			outputIndex = 0
			r.painted[r.pos] = outputs[0]
			// fmt.Printf("Robot has now painted: %v\n", r.painted)
			if outputs[1] == 0 {
				// fmt.Println("turning left")
				r.vec.Rotate(left)
			} else {
				// fmt.Println("turning right")
				r.vec.Rotate(right)
			}
			r.pos.Translate(r.vec)
			// fmt.Printf("Robot now at: %v\n", r.pos)

			newColour, found := r.painted[r.pos]
			if !found {
				newColour = black
			}
			r.camera <- newColour
		}
	}
}

func render(painted map[position.Coord]int) {
	minX, maxX, minY, maxY := getBounds(painted)

	// Y is backwards, I guess?
	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			pos := position.Coord{X: x, Y: y}
			colour, found := painted[pos]
			if !found {
				colour = black
			}
			if colour == white {
				fmt.Print("□")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func getBounds(painted map[position.Coord]int) (int, int, int, int) {
	minX := math.MaxInt64
	minY := math.MaxInt64
	maxX := math.MinInt64
	maxY := math.MinInt64

	for k := range painted {
		if k.X > maxX {
			maxX = k.X
		}
		if k.X < minX {
			minX = k.X
		}
		if k.Y > maxY {
			maxY = k.Y
		}
		if k.Y < minY {
			minY = k.Y
		}
	}

	return minX, maxX, minY, maxY
}
