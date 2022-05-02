package main

import (
	"fmt"
	"strings"
)

func numberThousandToWords(num int) string {
	if num >= 1000 {
		panic("Invalid number")
	}

	digitToWords := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	special := map[int]string{2: "twen", 3: "thir", 4: "for", 5: "fif", 8: "eigh"}
	specialHun := map[int]string{10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen", 18: "eighteen"}

	digits := [3]int{}
	count := 0
	for num > 0 {
		digits[count] = num % 10
		num /= 10
		count += 1
	}

	words := ""
	if digits[2] > 0 {
		words += digitToWords[digits[2]] + " " + "hundred"
	}

	if digits[1] > 0 {
		if digits[1] == 1 {
			if v, ok := specialHun[10*digits[1]+digits[0]]; ok {
				words += " " + v
			} else {
				words += " " + digitToWords[digits[0]] + "teen"
			}
			return strings.Trim(words, " ")
		}
		if v, ok := special[digits[1]]; ok {
			words += " " + v + "ty"
		} else {
			words += " " + digitToWords[digits[1]] + "ty"
		}
	}

	if digits[0] > 0 {
		words += " " + digitToWords[digits[0]]
	}

	return strings.Trim(words, " ")
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	blocks := [4]int{}
	count := 0
	for num > 0 {
		blocks[count] = num % 1000
		num /= 1000
		count += 1
	}
	words := ""
	if blocks[3] > 0 {
		words += " " + numberThousandToWords(blocks[3]) + " " + "billion"
	}
	if blocks[2] > 0 {
		words += " " + numberThousandToWords(blocks[2]) + " " + "million"
	}
	if blocks[1] > 0 {
		words += " " + numberThousandToWords(blocks[1]) + " " + "thousand"
	}
	words += " " + numberThousandToWords(blocks[0])
	return strings.Title(strings.Trim(words, " "))
}

func main() {
	fmt.Println(numberToWords(1010))
}
