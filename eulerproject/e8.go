/*
@author: bevin

@Created on: 2016-12-06 20:16:14
*/

package main

import "fmt"
import "os"
import "strings"

func main() {
	fi, err := os.Open("e8_data.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	buf := make([]byte, 1038)
	fi.Read(buf)
	ss := string(buf)
	ss = strings.Replace(ss, "\r\n", "", -1) //windows '\r\n' other '\n'
	//fmt.Println(len(ss))

	max := 0
	for i := 0; i < len(ss)-13; i++ {
		tmp := 1
		for j := 0; j < 13; j++ {
			tmp *= int(ss[i+j] - 48)
		}
		if tmp > max {
			max = tmp
		}
	}
	fmt.Println(max)
}
