package main

import (
	"fmt"
	"math"
	"time"
)

var totalPrimeNumbers int32 = 0
var MAX_INT = 100000000

func checkPrime(x int) {
	if x&1 == 0 {
		return
	}
	for i := 3; i < int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return
		}
	}
	totalPrimeNumbers++
}
func main() {
	start := time.Now()
	fmt.Println("program started")
	for i := 3; i < MAX_INT; i++ {
		checkPrime(i)
	}
	fmt.Println("checking till ", MAX_INT, "found :", totalPrimeNumbers, "prime numbers. took", time.Since(start))

}
