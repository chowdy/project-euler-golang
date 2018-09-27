package main

import "strconv"

/*
 * 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
 * What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 */

func Solution005() string {

	answer := 0
	MAX := 20

	Loop:
	for n := MAX; true; n++ {
		for i := 2; i <= MAX; i++ {
			if n % i != 0 {
				break
			} else if i == MAX {
				answer = n
				break Loop
			}
		}
	}

	return strconv.Itoa(answer)
}