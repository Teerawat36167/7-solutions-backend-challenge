package main

import (
	"fmt"
	"strconv"
)

func main() {
	var encoded string
	fmt.Print("Enter encoded string: ")
	fmt.Scanln(&encoded)

	result := findMinSumDecoding(encoded)
	fmt.Println("Decoded output:", result)
}

func findMinSumDecoding(encoded string) string {
	length := len(encoded) + 1
	minSum := -1
	var bestSum []int

	var backtrack func(pos int, current []int, sum int)
	backtrack = func(pos int, current []int, sum int) {
		if pos == length {
			valid := true
			for i := 0; i < len(encoded); i++ {
				left := current[i]
				right := current[i+1]

				if encoded[i] == 'L' && left <= right {
					valid = false
					break
				} else if encoded[i] == 'R' && left >= right {
					valid = false
					break
				} else if encoded[i] == '=' && left != right {
					valid = false
					break
				}
			}

			if valid && (minSum == -1 || sum < minSum) {
				minSum = sum
				bestSum = make([]int, length)
				copy(bestSum, current)
			}
			return
		}

		for digit := 0; digit <= 9; digit++ {
			current[pos] = digit
			valid := true

			if pos > 0 {
				left := current[pos-1]
				right := digit

				if encoded[pos-1] == 'L' && left <= right {
					valid = false
				} else if encoded[pos-1] == 'R' && left >= right {
					valid = false
				} else if encoded[pos-1] == '=' && left != right {
					valid = false
				}
			}

			if valid {
				backtrack(pos+1, current, sum+digit)
			}
		}
	}

	backtrack(0, make([]int, length), 0)

	result := ""
	for _, digit := range bestSum {
		result += strconv.Itoa(digit)
	}

	return result
}
