package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
)

func main() {
	var factorial bool
	flag.BoolVar(&factorial, "f", false, "Number is factorial")
	flag.Parse()

	if flag.Arg(0) == "" {
		fmt.Println("Number required")
		os.Exit(1)
	}

	factorsOfNum := []*big.Int{big.NewInt(1)}
	if factorial {
		for num, _ := new(big.Int).SetString(flag.Arg(0), 10); num.Cmp(big.NewInt(0)) == 1; num.Sub(num, big.NewInt(1)) {
			factorsOfNum = append(factorsOfNum, factors(new(big.Int).Set(num))...)
		}
	} else {
		num, _ := new(big.Int).SetString(flag.Arg(0), 10)
		factorsOfNum = append(factorsOfNum, factors(new(big.Int).Set(num))...)
	}

	fmt.Println(factorsOfNum)
	fmt.Println(factorsOfNum[len(factorsOfNum)-1])

}

func factors(num *big.Int) []*big.Int {

	//primes := []int{1}
	primes := []*big.Int{big.NewInt(1)}
	//factorsOfNum := []int{1}
	factorsOfNum := []*big.Int{}

	//for i := 2; i <= num/primes[len(primes)-1] && num > 1; i++ {
	for i := big.NewInt(2); i.Cmp(new(big.Int).Quo(num, primes[len(primes)-1])) <= 0; i.Add(i, big.NewInt(1)) {
		isPrime := true
		for _, prime := range primes {
			//if i%prime == 0 && prime > 1 {
			if new(big.Int).Mod(i, prime).Cmp(big.NewInt(0)) == 0 && prime.Cmp(big.NewInt(1)) == 1 {
				isPrime = false
			}
		}
		if isPrime {
			fmt.Printf("%d\r", i)
			//primes = append(primes, i)
			primes = append(primes, new(big.Int).Set(i))
			//for num%i == 0 {
			for new(big.Int).Mod(num, i).Cmp(big.NewInt(0)) == 0 {
				//factorsOfNum = append(factorsOfNum, i)
				factorsOfNum = append(factorsOfNum, new(big.Int).Set(i))
				//num = num / i
				num = new(big.Int).Div(num, i)
				//fmt.Println(i, num)
			}
		}
	}
	//if num > 1 {
	if num.Cmp(big.NewInt(1)) == 1 {
		//factorsOfNum = append(factorsOfNum, num)
		factorsOfNum = append(factorsOfNum, new(big.Int).Set(num))
	}

	return factorsOfNum
}
