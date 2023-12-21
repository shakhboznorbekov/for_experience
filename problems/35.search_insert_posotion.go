// package main

// import (
// 	"fmt"
// )
// func main(){
// 	a := []int{1,3,5,6}
// 	b := 5
// 	fmt.Println(searchInsert(a,b))
// }
// func searchInsert(nums []int, target int) int {
// 	left, right := 0, len(nums)
// 	for left < right {
// 		mid := (left + right) >> 1
// 		if nums[mid] >= target {
// 			right = mid
// 		} else {
// 			left = mid + 1
// 		}
// 	}
// 	return left
// }