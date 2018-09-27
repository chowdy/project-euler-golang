package main

import "strconv"

/*
 * A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
 * a^2 + b^2 = c^2
 * For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
 *
 * There exists exactly one Pythagorean triplet for which a + b + c = 1000.
 * Find the product abc.
 */

func Solution009() string {

	answer := 0

	Loop:
		for a := 1; a < 1000; a++ {
			for b := a+1; b < 1000; b++ {
				for c := b+1; c < 1000; c++ {
					if a + b + c == 1000 && a*a + b*b == c*c {
						answer = a*b*c
						break Loop
					}
				}
			}
		}

	return strconv.Itoa(answer)
}