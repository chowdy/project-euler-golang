package main

import (
	"project-euler-golang/utils"
	"strconv"
)

/*
 * The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
 * Find the sum of all the primes below two million.
 */

func Solution010() string {

	MAX := 2000000
	answer := 0

	for i := 0; i < MAX; i++ {
		if utils.IsPrimeNaive(i) {
			answer += i
		}
	}

	return strconv.Itoa(answer)
}
