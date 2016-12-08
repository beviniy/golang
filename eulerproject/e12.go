/*
@Author: bevin

@Created on: 2016-12-08 11:52:38
*/
package main

import (
	"fmt"
)

func main() {
	cur := 0
	tmp := 1
	for {
		cur += tmp
		tmp++
		res := getNumofDividors(cur)
		if res > 500 {
			break
		}
	}
	fmt.Println(cur)

}

func getNumofDividors(num int) int {
	count := 1
	ysmap := make(map[int]int)
	for num != 1 {
		for i := 2; i <= num; i++ {
			if num%i == 0 {
				ysmap[i]++
				num /= i
				break
			}
		}
	}
	for _, v := range ysmap {
		count *= (v + 1)
	}
	return count
}
