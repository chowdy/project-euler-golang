package main

import (
	"fmt"
	"math/big"
)

/*
 * n! means n × (n − 1) × ... × 3 × 2 × 1
 *
 * For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
 * and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
 *
 * Find the sum of the digits in the number 100!
 */
func Solution020() string {

	// Calculate 100!❗️
	FACTORIAL := int64(100)
	fact := big.NewInt(1)
	big1 := big.NewInt(1)
	for n := big.NewInt(FACTORIAL); n.Cmp(big1) == 1; n.Sub(n, big1) {
		fact.Mul(fact, n)
	}

	// Sum the digits
	answer := 0
	for _, c := range fact.String() {
		answer += int(c - '0')
	}

	return fmt.Sprint(answer)
}
