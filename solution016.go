package main

import (
	"fmt"
	"math/big"
)

/*
 * 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
 * What is the sum of the digits of the number 2^1000?
 */

func Solution016() string {

	// Calculate 2^1000, big numbers style
	n := big.NewInt(2)
	big2 := big.NewInt(2)
	for i := 1; i < 1000; i++ {
		n.Mul(n, big2)
	}

	// Sum the digits
	answer := 0
	for _, c := range n.String() {
		answer += int(c - '0')
	}

	return fmt.Sprint(answer)
}
