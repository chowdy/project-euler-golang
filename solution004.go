package main

import (
	"fmt"
	"./utils"
)

/*
 * A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers
 * is 9009 = 91 × 99.

 * Find the largest palindrome made from the product of two 3-digit numbers.
 */

func Solution004() string {

	answer := 0

	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			prod := i*j
			if utils.IsPallindromeInt(prod) && prod > answer {
				answer = prod
			}
		}
	}

	return fmt.Sprint(answer)

}
