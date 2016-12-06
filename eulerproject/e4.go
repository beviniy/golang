/*
@Author: bevin

@Created on: 2016-12-06 14:57:40
*/

package main

import "fmt"
import "math"
import "strconv"

func main() {

	res := make([]int, 0, 10000)
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			if judgeNum(i * j) {
				res = append(res, i*j)
			}
		}
	}
	max := 0
	for _, number := range res {
		if max < number {
			max = number
		}
	}
	fmt.Println(max)
}

func judgeNum(number int) bool {
	tmp := 0
	for i, n := range strconv.Itoa(number) {
		tmp += (int(n) - 48) * int(math.Pow10(i))
	}
	return number == tmp
}
