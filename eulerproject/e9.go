/*
@author: bevin

@Created on: 2016-12-06 20:48:49
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	myPow_2 := func(n int) int {
		return int(math.Pow(float64(n), 2))
	}
	res := 0
label:
	for c := 1000 / 3; c < 500; c++ {
		b := 0
		for b < c {
			a := 1000 - b - c
			if myPow_2(a)+myPow_2(b) == myPow_2(c) {
				res = a * b * c
				fmt.Println(a, b, c)
				break label
			}
			b++
		}
	}
	fmt.Println(res)
}
