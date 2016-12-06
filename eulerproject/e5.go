/*
@Author: bevin

@Created on: 2016-12-06 16:14:24
*/

package main

import "fmt"

func main() {
	//fmt.Println(getYS(20))
	res_map := make(map[int]int)
	for i := 1; i <= 20; i++ {
		nmap := getYS(i)
		update_map(&nmap, &res_map)
	}
	res := 1
	//fmt.Println(res_map)
	for k, v := range res_map {
		for i := 0; i < v; i++ {
			res *= k
		}
	}
	fmt.Println(res)
}

func getYS(number int) map[int]int {
	res_map := make(map[int]int)
	for number != 1 {
		for i := 2; i <= number; i++ {
			if number%i == 0 {
				res_map[i]++
				number /= i
				break
			}
		}
	}
	return res_map
}

func update_map(nmap *map[int]int, cur *map[int]int) {
	for k, v := range *nmap {
		if (*cur)[k] < v {
			(*cur)[k] = v
		}
	}
}
