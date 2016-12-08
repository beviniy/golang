/*
@Author: bevin

@Created on: 2016-12-06 14:15:23
*/

package main

import "fmt"
import "math"
import "reflect"
import "strconv"

func main() {
	a := 1
	b := 1.3
	fmt.Println(math.Sqrt(b))
	//fmt.Println(a * b)
	fmt.Println(a < int(b))
	c := "123"
	d := make([]rune, 2, 3)
	fmt.Println(len(c))
	//fmt.Println("%p", &d)
	fmt.Printf("%p,%v\n", d, d)
	d = append(d, rune(48))
	fmt.Printf("%p,%v\n", d, d)
	d = append(d, rune(48))
	fmt.Printf("%p,%v\n", d, d)
	e := rune(49)
	f := rune(1)
	g := e - f
	fmt.Println(g)
	for _, i := range "abcde" {
		fmt.Println(i - int32(1))
		fmt.Println(reflect.TypeOf(i))
	}
	h := make(map[int]int)
	fmt.Println(h[1])
	// i := 1
	// j := 1.0
	fmt.Println(100 / 3)

	ss := "08 10"
	fmt.Println(strconv.Atoi(ss))
}
