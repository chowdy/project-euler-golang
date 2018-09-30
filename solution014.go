package main

import "fmt"

/*
 * The following iterative sequence is defined for the set of positive integers:
 *
 * n → n/2 (n is even)
 * n → 3n + 1 (n is odd)
 *
 * Using the rule above and starting with 13, we generate the following sequence:
 *
 * 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
 * It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
 * Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
 *
 * Which starting number, under one million, produces the longest chain?
 *
 * NOTE: Once the chain starts the terms are allowed to go above one million.
 */

var collatzChainLengthMemo = map[int]int {
	13: 10,
}

func collatzChainLength(n int) int {
	memo, ok := collatzChainLengthMemo[n]
	if ok { return memo }

	length := 1
	curr := n
	for curr != 1 {

		memo, ok := collatzChainLengthMemo[curr]
		if ok { return memo + length }

		if curr % 2 == 0 {
			curr = curr / 2
		} else {
			curr = curr * 3 + 1
		}
		length++
	}

	collatzChainLengthMemo[n] = length
	return length
}

func Solution014() string {

	answer := 0
	max := 0
	for i := 1; i < 1000000; i++ {
		length := collatzChainLength(i)
		if length > max {
			answer = i
			max = length
		}
	}

	return fmt.Sprint(answer)
}