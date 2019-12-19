package position

import "math"

// Coord is a 2-dimensional coordinate
type Coord struct {
	X int
	Y int
}

// Rotate rotates a Coord by the given angle in degrees
func (c *Coord) Rotate(theta float64) {
	theta = theta / 180.0 * math.Pi
	X := float64(c.X)*math.Cos(theta) - float64(c.Y)*math.Sin(theta)
	Y := float64(c.X)*math.Sin(theta) + float64(c.Y)*math.Cos(theta)
	c.X = int(X)
	c.Y = int(Y)
}

// Translate translates a Coord by the given Coord
func (c *Coord) Translate(v Coord) {
	c.X += v.X
	c.Y += v.Y
}

// Vec3 is a 3-element vector
type Vec3 struct {
	X int
	Y int
	Z int
}

// Translate translates a Vec3 by the given Vec3
func (v *Vec3) Translate(v2 Vec3) {
	v.X += v2.X
	v.Y += v2.Y
	v.Z += v2.Z
}

// Equals compares all elements of v against v2 and returns true if they are all equal
func (v *Vec3) Equals(v2 Vec3) bool {
	return v.X == v2.X && v.Y == v2.Y && v.Z == v2.Z
}

// Copy returns a copy of a Vec3
func (v *Vec3) Copy() Vec3 {
	return Vec3{X: v.X, Y: v.Y, Z: v.Z}
}
