package main

import (
	"project-euler-golang/utils"
	"strconv"
)

/*
 * By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
 * What is the 10 001st prime number?
 */

func Solution007() string {

	MAX := 10001
	count := 1
	lastPrime := 2
	curr := 3

	for count < MAX {
		if utils.IsPrimeNaive(curr) {
			count++
			lastPrime = curr
		}
		curr++
	}

	return strconv.Itoa(lastPrime)
}
