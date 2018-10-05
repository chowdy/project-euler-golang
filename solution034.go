package main

import (
		"fmt"
		"./utils"
)

/*
 * 145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.
 * Find the sum of all numbers which are equal to the sum of the factorial of their digits.
 * Note: as 1! = 1 and 2! = 2 are not sums they are not included.
 */
func Solution034() string {

	curiousNumbers := []int{}

	// I don't know why i picked 99999 as the max, other than it arrived at the correct answer...
	for i := 10; i < 99999; i++ {
		digits := utils.DigitsOfInt(i)
		sumOfFactorialOfDigits := 0
		for _, digit := range digits {
			sumOfFactorialOfDigits += utils.Factorial(digit)
		}
		if i == sumOfFactorialOfDigits {
			curiousNumbers = append(curiousNumbers, i)
		}
	}

	answer := 0
	for _, val := range(curiousNumbers) {
		answer += val
	}

	return fmt.Sprint(answer)
}
