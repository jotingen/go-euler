package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"sort"
)

type BigInts []*big.Int
func (s BigInts) Len() int      { return len(s) }
func (s BigInts) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type BySize struct{ BigInts }
func (s BySize) Less(i, j int) bool { return s.BigInts[i].Cmp(s.BigInts[j])==-1 }

func main() {
	var factorial bool
	flag.BoolVar(&factorial, "f", false, "Number is factorial")
	flag.Parse()

	if flag.Arg(0) == "" {
		fmt.Println("Number required")
		os.Exit(1)
	}
	if flag.Arg(1) == "" {
		fmt.Println("k required")
		os.Exit(1)
	}

	//k, _ := new(big.Int).SetString(flag.Arg(1), 10)

	factorsOfNum := []*big.Int{big.NewInt(1)}
	if factorial {
		for num, _ := new(big.Int).SetString(flag.Arg(0), 10); num.Cmp(big.NewInt(0)) == 1; num.Sub(num, big.NewInt(1)) {
			factorsOfNum = append(factorsOfNum, factors(new(big.Int).Set(num))...)
		}
		sort.Sort(BySize{factorsOfNum})
	} else {
		num, _ := new(big.Int).SetString(flag.Arg(0), 10)
		factorsOfNum = append(factorsOfNum, factors(new(big.Int).Set(num))...)
	}
	fmt.Println(factorsOfNum)



	numberOfFactors := len(factorsOfNum)
	fmt.Println(numberOfFactors)

	//TODO Unifinished, need to figure a way to combine primes into unique combinations for k buckets
		
	//  //Equation 
	//  //  (factors+k-1)!/(factors!(k-1)!)
	//  //Equation with 1 required in each bucket 
	//  //  (factors-1)!/((factors-k)!(k-1)!)
	//  factors_m1_factorial := big.NewInt(1) 
	//  for i := numberOfFactors-1; i > 0; i-- {
	//         factors_m1_factorial = new(big.Int).Mul(factors_m1_factorial,big.NewInt(int64(i))) 
	//  }
	//  fmt.Println(factors_m1_factorial)

	//  factors_mk_factorial := big.NewInt(1) 
	//  for i := new(big.Int).Sub(big.NewInt(int64(numberOfFactors)),k); i.Cmp(big.NewInt(0))==1; i.Sub(i, big.NewInt(1)) {
	//         factors_mk_factorial = new(big.Int).Mul(factors_mk_factorial,i) 
	//  }              
	//  fmt.Println(factors_mk_factorial)

	//  k_m1_factorial := big.NewInt(1) 
	//  for i := new(big.Int).Sub(k,big.NewInt(1)); i.Cmp(big.NewInt(0))==1; i.Sub(i, big.NewInt(1)) {
	//         k_m1_factorial = new(big.Int).Mul(k_m1_factorial,i) 
	//  }            
	//  fmt.Println(k_m1_factorial)

	//  ways := new(big.Int).Div(factors_m1_factorial,new(big.Int).Mul(factors_mk_factorial,k_m1_factorial))

	//  fmt.Println(ways)

}

//Reused from 003
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
