package main

import (
	"./utils"
	"strconv"
)

/*
 * The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
 * Find the sum of all the primes below two million.
 */

func Solution010() string {

	MAX := 2000000
	answer := 0
	utils.BuildPrimesSieve(MAX)
	for _, p := range utils.GetPrimes() {
		answer += p
	}
	return strconv.Itoa(answer)

}
