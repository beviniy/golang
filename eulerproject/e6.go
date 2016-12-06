/*
@Author: bevin

@Created on: 2016-12-06 16:47:27
*/
package main

import "fmt"
import "math"

func main() {
	myPow_2 := func(n int) int {
		return int(math.Pow(float64(n), 2))
	}
	number := 100
	sum1 := 0
	sum2 := 0
	for i := 1; i <= number; i++ {
		sum1 += myPow_2(i)
		sum2 += i
	}
	sum := myPow_2(sum2) - sum1
	fmt.Println(sum)
}
