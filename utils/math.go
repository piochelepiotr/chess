package utils

// Absolute returns the absolute value of an int
func Absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Sign returns the +-1 depending on the sign of the number
func Sign(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}
