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

// How many solutions can I get away with using this for?
func IsPrimeNaive(n int) bool {

	if n < 2 {
		return false
	}

	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

// 0 and 1 are not prime, but 2 is!
var sieve = []bool{ false, false, true }
func IsPrime(n int) bool {

	// If the sieve slice is too small, then grow it and rediscover the primes
	if len(sieve) < n {
		// TODO: This
	}

	return sieve[n]
}