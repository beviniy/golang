/*
@author: bevin

@Created on: 2016-12-06 10:49:12
*/

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 1000; i++ {
		if (i%3) == 0 || (i%5) == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
