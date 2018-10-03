package main

import (
	"fmt"
	"./utils"
)

/*
 * Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
 * If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable
 * numbers.
 *
 * For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284.
 * The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.
 *
 * Evaluate the sum of all the amicable numbers under 10000.
 */
func Solution021() string {

	// Calculate d for every number below 10000
	dMap := map[int]int{}
	paMd := map[int]int{}
	for i := 2; i < 10000; i++ {
		d := 0
		for _, properDiv := range utils.GetProperDivisors(i) {
			d += properDiv
		}
		dMap[i] = d
		paMd[d] = i
	}

	// Find all amicable numbers
	var amicableNums []int
	for num, d := range dMap {
		println("d(", num, ") = ", d)

		//friend, friendFound := dMap[d]
		//if friendFound && friend != num {
		//	amicableNums = append(amicableNums, num)
		//	println(num, " (", dMap[num], ") and ", friend, " (", dMap[friend], ")")
		//}
	}

	// Sum all of the amicable numbers found
	answer := 0
	for _, num := range amicableNums {
		answer += num
	}

	return fmt.Sprint(answer)
}
