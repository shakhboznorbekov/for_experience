// package main

// import (
// 	"fmt"
// 	"sort"
// )
// func main(){
// 	a := []int{0,1}
// 	fmt.Println(missingNumber(a))
// }
// func missingNumber1(nums []int) int {
// 	var a int
// 	sort.Ints(nums)
// 	for i:=0; i<len(nums)-1; i++{
// 		if (nums[i]+1)!=nums[i+1]{
// 			a = nums[i]+1
// 		}
// 	}
// 	return a
// }
// func missingNumber(nums []int) (ans int) {
// 	n := len(nums)
// 	ans = n
// 	for i, v := range nums {
// 		ans ^= (i ^ v)
// 	}
// 	return
// }