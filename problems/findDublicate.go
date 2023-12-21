// package main

// import "fmt"

// func main(){
// 	nums := []int{1,3,4,2,2}
// 	for i:=0; i<len(nums); i++{
// 		for j:=i+1; j<len(nums); j++{
// 			if nums[i]== nums[j] {
// 				fmt.Println(nums[i])
// 			}
// 		}
// 	}
// }

// func findDuplicate(nums []int) int {
// 	left, right := 1, len(nums)-1
// 	for left < right {
// 		mid := (left + right) >> 1
// 		cnt := 0
// 		for _, v := range nums {
// 			if v <= mid {
// 				cnt++
// 			}
// 		}
// 		if cnt > mid {
// 			right = mid
// 		} else {
// 			left = mid + 1
// 		}
// 	}
// 	return left
// }