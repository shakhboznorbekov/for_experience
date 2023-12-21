// package main

// import (
// 	"fmt"
// )
// func main(){
// 	a := []int{1,2,3,4,5,5,7}
// 	fmt.Println(findErrorNums(a))
// }
// func findErrorNums(nums []int) []int {
// 	n := len(nums)
// 	for i := 0; i < n; i++ {
// 		for nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
// 			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
// 		}
// 	}
// 	for i := 0; i < n; i++ {
// 		if nums[i] != i+1 {
// 			return []int{nums[i], i + 1}
// 		}
// 	}
// 	return []int{-1, -1}
// }