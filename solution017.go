package main

import "fmt"

/*
 * If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19
 * letters used in total.
 *
 * If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?
 *
 * NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115
 * (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with
 * British usage.
 */

var onesWords = map[int]string {
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var teens = map[int]string {
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

var tensWords = map[int]string {
	1: "ten",
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

func numInWords(n int) string {

	wordy := ""

	if n < 10 {
		wordy = onesWords[n]
	} else if n > 10 && n < 20 {
		wordy = teens[n]
	} else if n < 100 {
		tens := n / 10
		ones := n - tens*10

		tensWord, _ := tensWords[tens]
		onesWord, _ := onesWords[ones]

		wordy = tensWord + onesWord
	} else if n < 1000 {
		hundreds := n / 100
		tens := (n - hundreds*100) / 10
		ones := n - hundreds*100 - tens*10

		hundredsWord, _ := onesWords[hundreds]
		wordy += hundredsWord + "hundred"

		if tens == 1 && n - hundreds * 100 < 20 {
			wordy += "and"
			wordy += teens[n - hundreds * 100]
		} else {
			wordy += "and"
			wordy += tensWords[tens]
			wordy += onesWords[ones]
		}
	} else if n == 1000 {
		return "onethousand"
	}

	return wordy
}

func Solution017() string {
	answer := 0
	for i := 1; i <= 1000; i++ {
		answer += len(numInWords(i))
		println(i, " : ", numInWords(i))
	}
	return fmt.Sprint(answer)
}
