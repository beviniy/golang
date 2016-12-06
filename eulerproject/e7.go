/*
@Author: bevin

@Created on: 2016-12-06 16:57:09
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	count := 0
	cur := 2
	for {
		if isPrime(cur) {
			count++
			if count == 10001 {
				break
			}
		}
		cur++
	}
	fmt.Println(cur)
}

func isPrime(number int) bool {
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
