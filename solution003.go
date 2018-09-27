package main

import (
	"fmt"
	"sort"
	"./utils"
)

/*
 * The prime factors of 13195 are 5, 7, 13 and 29.
 * What is the largest prime factor of the number 600851475143 ?
 */

func Solution003() string {

	pfs := utils.PrimeFactors(600851475143)
	sort.Ints(pfs)
	answer := pfs[len(pfs)-1]
	return fmt.Sprint(answer)

}
