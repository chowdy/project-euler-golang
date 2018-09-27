package main

import (
	"fmt"
	"sort"
)

/*
 * The prime factors of 13195 are 5, 7, 13 and 29.
 * What is the largest prime factor of the number 600851475143 ?
 */

func primeFactors(n int) []int {

	pfs := make([]int, 0)

	for i := 2; n > 1; i++ {
		if n % i == 0 {
			pfs = append(pfs, i)
			n /= i
		}
	}

	return pfs
}

func Solution003() string {

	pfs := primeFactors(600851475143)
	sort.Ints(pfs)
	answer := pfs[len(pfs)-1]
	return fmt.Sprint(answer)

}
