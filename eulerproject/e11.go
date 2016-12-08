/*
@Author: bevin

@Created on: 2016-12-08 10:04:58
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers := readfile("e11_data.txt")
	max := 0
	maxfunc := func(nums []int) int {
		max := 0
		for _, each := range nums {
			if max < each {
				max = each
			}
		}
		return max
	}

	for i := 0; i < 16; i++ { //four directions
		for j := 0; j < 20; j++ {
			tmp1, tmp2, tmp3, tmp4 := 1, 1, 1, 1
			for k := 0; k < 4; k++ {
				tmp1 *= numbers[i+k][j]
				tmp2 *= numbers[j][i+k]
				if j < 16 {
					tmp3 *= numbers[i+k][j+k]
					tmp4 *= numbers[i+3-k][j+k]
				}
			}
			var nums = []int{tmp1, tmp2, tmp3, tmp4, max}
			max = maxfunc(nums)
		}
	}
	fmt.Println(max)
}

func readfile(filename string) (numbers [20][20]int) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)
	count := 0
	for {
		line, err := bfRd.ReadBytes('\n')
		for i, ns := range strings.Split(strings.Replace(string(line), "\r\n", "", -1), " ") {
			numbers[count][i], _ = strconv.Atoi(ns)
		}
		count++
		if err != nil {
			break
		}
	}
	return
}
