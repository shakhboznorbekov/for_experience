// package main

// import "fmt"

// //leetcode 1 answer
// func main(){
// 	var sonlar []int = []int{2, 7, 11, 13}
// 	var target int = 9
// 	fmt.Println(twoSum(sonlar,target))
// }

// func twoSum(nums []int, target int) []int {
// 	for i, j := range nums {
// 		for k := i + 1; k < len(nums); k++ {
// 			if nums[k]+j == target {
// 				return []int{i, k}
// 			}
// 		}
// 	}
// 	return []int{}
// }