package main

import (
	"fmt"
)
func main(){
	a := []int{2,2,1,1,1,2,2}
	fmt.Println(majorityElement(a))
}
func majorityElement(nums []int) int {
	cnt, m := 0, 0
	for _, v := range nums {
		if cnt == 0 {
			m, cnt = v, 1
		} else {
			if m == v {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return m
}