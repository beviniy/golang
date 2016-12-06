/*
@author: bevin

@Created on: 2016-12-06 10:49:12
*/

package main

import "fmt"

func main() {
	sum := 0
	i1 := 0
	i2 := 1
	for {
		current := i1 + i2
		if current > 4000000 {
			break
		}
		if current%2 == 0 {
			sum += current
		}
		i1 = i2
		i2 = current
	}
	fmt.Println(sum)
}
