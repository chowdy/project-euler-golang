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

func IsPallindrome(n int) bool {

	nAsString := strconv.Itoa(n)
	nLen := len(nAsString)

	for i := 0; i < len(nAsString); i++ {
		if nAsString[i] != nAsString[nLen-i-1] {
			return false
		}
	}

	return true

}