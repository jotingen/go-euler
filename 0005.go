package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if os.Args[1] == "" {
		fmt.Println("Numbers required")
		os.Exit(1)
	}
	divisors, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Numbers required")
		os.Exit(1)
	}

	num := divisors //Has to start at first divisor
	for  {
		evenlyDivisible := true
		for divisor := 1; divisor <= divisors; divisor++ {
			if num%divisor != 0 {
				evenlyDivisible = false
				break
			}
		}
		fmt.Printf("%d\r",num)
		if evenlyDivisible {
			break
		}
		num = num + divisors //Has to be incrmented by divisors 
	}

		fmt.Printf("%d\n",num)

}
