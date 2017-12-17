package main

import (
	"fmt"
)

func main() {

	sum := 0
	i := 1
	j := 1
	for i < 4000000 {
		//fmt.Println(i)
		if i%2 == 0 {
			sum += i
		}
		k := i + j
		i = j
		j = k
	}
	fmt.Println(sum)
}
