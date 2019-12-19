package main

import (
	mmmmath "github.com/a-harvie/advent-2019/mathforrobpike"
)

func (m moon) energy() int {
	potential := mmmmath.Abs(m.pos.X) + mmmmath.Abs(m.pos.Y) + mmmmath.Abs(m.pos.Z)
	kinetic := mmmmath.Abs(m.vel.X) + mmmmath.Abs(m.vel.Y) + mmmmath.Abs(m.vel.Z)
	return potential * kinetic
}

func (m *moon) gravitate() {
	m.vel.X += m.grav.X
	m.vel.Y += m.grav.Y
	m.vel.Z += m.grav.Z
	m.grav.X = 0
	m.grav.Y = 0
	m.grav.Z = 0
}

func (m *moon) move() {
	m.pos.X += m.vel.X
	m.pos.Y += m.vel.Y
	m.pos.Z += m.vel.Z
}

func (m *moon) positionInDimension(dimension string) int {
	position := 0
	switch dimension {
	case "X":
		position = m.pos.X
	case "Y":
		position = m.pos.Y
	case "Z":
		position = m.pos.Z
	}

	return position
}
