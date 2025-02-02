package calculator

import "math"

// IsEven returns true if the number is even
func IsEven(n int) bool {
	return n%2 == 0
}

// IsOdd returns true if the number is odd
func IsOdd(n int) bool {
	return !IsEven(n)
}

// IsPrime returns true if the number is prime
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// IsPerfect returns true if n equals the sum of its proper divisors
func IsPerfect(n int) bool {
	if n <= 1 {
		return false
	}
	sum := 1
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum == n
}

// GetDigitSum returns the sum of all digits in n
func GetDigitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// IsArmstrong returns true if n is an Armstrong number
func IsArmstrong(n int) bool {
	if n <= 0 {
		return false
	}
	
	// Count digits
	temp := n
	numDigits := 0
	for temp > 0 {
		numDigits++
		temp /= 10
	}
	
	// Calculate sum of digits raised to power
	temp = n
	sum := 0
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(numDigits)))
		temp /= 10
	}
	
	return sum == n
}
