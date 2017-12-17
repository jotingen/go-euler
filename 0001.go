package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if os.Args[1] == "" {
		fmt.Println("Number required")
		os.Exit(1)
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Number required")
		os.Exit(1)
	}

	sum := 0
	for i := 1; i < num; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
