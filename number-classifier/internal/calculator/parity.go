package calculator

// IsEven returns true if the number is even
func IsEven(n int) bool {
	return n%2 == 0
}

// IsOdd returns true if the number is odd
func IsOdd(n int) bool {
	return !IsEven(n)
}
