/*
@Author: bevin

@Created on: 2016-12-06 13:58:11
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var N int = 600851475143
	for i := 2; i < N; i++ {
		if N%i == 0 {
			if isPrime(N / i) {
				fmt.Println(N / i)
				break
			}
		}
	}
}

func isPrime(number int) bool {
	for i := 2; i <= int(math.Sqrt(float64(number+1))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
