package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if os.Args[1] == "" {
		fmt.Println("Number of digits required")
		os.Exit(1)
	}
	digits, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Number of digits required")
		os.Exit(1)
	}

	for i := int(math.Pow10(digits) - 1); i > 0; i-- {
		for j := int(math.Pow10(digits) - 1); j >= i; j-- {
			product := i * j
			if isPalindrome(product) {
				fmt.Printf("%d %d  %d\n", i, j, product)
				os.Exit(0)
			}
		}
	}

}

func isPalindrome(num int) bool {
	runes := []rune(strconv.Itoa(num))
	palindrome := true
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			palindrome = false
		}
	}
	return palindrome
}
