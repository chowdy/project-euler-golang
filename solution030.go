package main

import (
	"./utils"
	"fmt"
	)

/*
 * Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:
 *
 * 1634 = 14 + 64 + 34 + 44
 * 8208 = 84 + 24 + 04 + 84
 * 9474 = 94 + 44 + 74 + 44
 * As 1 = 14 is not a sum it is not included.
 *
 * The sum of these numbers is 1634 + 8208 + 9474 = 19316.
 *
 * Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
 */

func Solution030() string {
	answer := 0

	for i := 10; i < 9*9*9*9*9 * 5; i++ {

		sumOfFifthPowers := 0
		for _, digit := range utils.DigitsOfInt(i) {
			sumOfFifthPowers += digit*digit*digit*digit*digit
		}

		if sumOfFifthPowers == i {
			answer += i
		}
	}
	return fmt.Sprint(answer)
}
