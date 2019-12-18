package position

import "math"

type Coord struct {
	X int
	Y int
}

func (c *Coord) Rotate(theta float64) {
	theta = theta / 180.0 * math.Pi
	X := float64(c.X)*math.Cos(theta) - float64(c.Y)*math.Sin(theta)
	Y := float64(c.X)*math.Sin(theta) + float64(c.Y)*math.Cos(theta)
	c.X = int(X)
	c.Y = int(Y)
}

func (c *Coord) Translate(v Coord) {
	c.X += v.X
	c.Y += v.Y
}
