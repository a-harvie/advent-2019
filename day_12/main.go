package main

import (
	"fmt"

	"github.com/a-harvie/advent-2019/input"
	sonofmath "github.com/a-harvie/advent-2019/mathforrobpike"
	"github.com/a-harvie/advent-2019/position"
)

type moon struct {
	pos        position.Vec3
	vel        position.Vec3
	grav       position.Vec3
	initialPos position.Vec3
	initialVel position.Vec3
}

type moonSystem []*moon

func main() {
	testInput, err := input.Vectors("input/testinput.txt")
	if err != nil {
		fmt.Printf("Oh noes: %v\n", err)
	}
	moons := makeMoons(testInput[:4])
	steps := 10
	moons.simulateNth(steps)
	fmt.Printf("State after %v\n", steps)
	moons.print()
	fmt.Printf("Total energy after %v: %v\n", steps, moons.totalEnergy())
	moons = makeMoons(testInput[:4])
	cycles := moons.findCycleCount()
	fmt.Printf("Moons cycle after %v\n\n", cycles)

	moons = makeMoons(testInput[4:])
	steps = 100
	moons.simulateNth(steps)
	fmt.Printf("State after %v\n", steps)
	moons.print()
	fmt.Printf("Total energy after %v: %v\n", steps, moons.totalEnergy())
	moons = makeMoons(testInput[4:])
	cycles = moons.findCycleCount()
	fmt.Printf("Moons cycle after %v\n\n", cycles)

	input, err := input.Vectors("input/input.txt")
	if err != nil {
		fmt.Printf("Oh noes: %v\n", err)
	}
	moons = makeMoons(input)
	steps = 1000
	moons.simulateNth(steps)
	fmt.Printf("State after %v\n", steps)
	moons.print()
	fmt.Printf("Total energy after %v: %v\n", steps, moons.totalEnergy())
	moons = makeMoons(input)
	cycles = moons.findCycleCount()
	fmt.Printf("Moons cycle after %v\n", cycles)
}

func makeMoons(startingPositions []([]int)) moonSystem {
	moons := make([]*moon, 0)
	for _, startingPos := range startingPositions {
		m := moon{
			pos: position.Vec3{
				X: startingPos[0],
				Y: startingPos[1],
				Z: startingPos[2],
			},
			vel:  position.Vec3{X: 0, Y: 0, Z: 0},
			grav: position.Vec3{X: 0, Y: 0, Z: 0},
			initialPos: position.Vec3{
				X: startingPos[0],
				Y: startingPos[1],
				Z: startingPos[2],
			},
			initialVel: position.Vec3{X: 0, Y: 0, Z: 0},
		}
		moons = append(moons, &m)
	}

	return moons
}

func (moons moonSystem) simulateNth(steps int) {
	for i := 0; i < steps; i++ {
		moons.simulate()
	}
}
func (moons moonSystem) simulate() {
	moons.gravitate()
	moons.move()
}

func (moons moonSystem) findCycleCount() int {

	xChan := make(chan int)
	yChan := make(chan int)
	zChan := make(chan int)
	go getDimensionCycle(moons, "X", xChan)
	go getDimensionCycle(moons, "Y", yChan)
	go getDimensionCycle(moons, "Z", zChan)
	xCycle := <-xChan
	yCycle := <-yChan
	zCycle := <-zChan

	// fmt.Printf("Got some cycles: %v %v %v\n", xCycle, yCycle, zCycle)

	return sonofmath.LCM(sonofmath.LCM(xCycle, yCycle), zCycle)
}

func getDimensionCycle(moons moonSystem, dimension string, ch chan int) {
	cycles := 1
	cycler := moons.copy()
	cycler.simulate()
	for !cycler.initalDimensionState(dimension) {
		cycler.simulate()
		cycles++
	}
	ch <- cycles
}

func (moons moonSystem) print() {
	for _, moon := range moons {
		fmt.Printf(
			"pos=<x=%v, y=%v, z=%v>, vel=<x=%v, y=%v, z=%v>\n",
			moon.pos.X,
			moon.pos.Y,
			moon.pos.Z,
			moon.vel.X,
			moon.vel.Y,
			moon.vel.Z,
		)
	}
}

func (moons moonSystem) copy() moonSystem {
	newMoonSystem := make(moonSystem, len(moons))
	for i, m := range moons {
		newMoonSystem[i] = &moon{
			pos:        m.pos.Copy(),
			vel:        m.vel.Copy(),
			grav:       m.grav.Copy(),
			initialPos: m.initialPos.Copy(),
			initialVel: m.initialVel.Copy(),
		}
	}
	return newMoonSystem
}

func (moons moonSystem) gravitate() {
	for _, m1 := range moons {
		for _, m2 := range moons {
			if m1 == m2 {
				continue
			}
			gravitateMoons(m1, m2)
		}
		m1.gravitate()
	}
}

func (moons moonSystem) totalEnergy() int {
	energy := 0
	for _, m := range moons {
		energy += m.energy()
	}
	return energy
}

func (moons moonSystem) initialState() bool {
	for _, m := range moons {
		if !m.pos.Equals(m.initialPos) || m.vel.Equals(m.initialVel) {
			return false
		}
	}
	return true
}

func (moons moonSystem) initalDimensionState(dimension string) bool {

	switch dimension {
	case "X":
		for _, m := range moons {
			if m.pos.X != m.initialPos.X || m.vel.X != m.initialVel.X {
				return false
			}
		}
	case "Y":
		for _, m := range moons {
			if m.pos.Y != m.initialPos.Y || m.vel.Y != m.initialVel.Y {
				return false
			}
		}
	case "Z":
		for _, m := range moons {
			if m.pos.Z != m.initialPos.Z || m.vel.Z != m.initialVel.Z {
				return false
			}
		}
	}
	return true
}

func gravitateMoons(m1 *moon, m2 *moon) {
	if m1.pos.X > m2.pos.X {
		m1.grav.X--
	} else if m1.pos.X < m2.pos.X {
		m1.grav.X++
	}

	if m1.pos.Y > m2.pos.Y {
		m1.grav.Y--
	} else if m1.pos.Y < m2.pos.Y {
		m1.grav.Y++
	}

	if m1.pos.Z > m2.pos.Z {
		m1.grav.Z--
	} else if m1.pos.Z < m2.pos.Z {
		m1.grav.Z++
	}
}

func (moons moonSystem) move() {
	for _, moon := range moons {
		moon.move()
	}
}
