/*
@Author: bevin

@Created on: 2016-12-08 10:00:23
*/
package main

import "fmt"
import "math"

func main() {
	N := 2000000
	sum := 0
	for i := 2; i < N; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}

func isPrime(number int) bool {
	for i := 2; i <= int(math.Sqrt(float64(number+1))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
