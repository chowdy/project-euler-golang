package utils

import (
	"strconv"
)

// Returns a slice of prime factors of n
func PrimeFactors(n int) []int {

	pfs := make([]int, 0)

	for i := 2; n > 1; i++ {
		if n % i == 0 {
			pfs = append(pfs, i)
			n /= i
		}
	}

	return pfs
}

func IsPallindromeInt(n int) bool {
	return IsPallindromeString(strconv.Itoa(n))
}

func IsPallindromeString(str string) bool {
	nLen := len(str)
	for i := 0; i < len(str); i++ {
		if str[i] != str[nLen-i-1] {
			return false
		}
	}
	return true
}