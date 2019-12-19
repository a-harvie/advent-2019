package mathforrobpike

// Abs is easy to implement for ints so why would the stdlib ever include it
func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// GCD is the greatest common denominator
func GCD(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// LCM is the least common multiple
func LCM(x int, y int) int {
	return Abs(x*y) / GCD(x, y)
}
