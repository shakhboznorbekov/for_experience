// package main

// import "fmt"

// func main(){
// 	a := []int{2,4,3}
// 	b := []int{5,6,4}
// 	c,d,k:= 0,0,0

// 	for _,val := range a{
// 		c = 10*c+val
// 	}

// 	for _,val := range b{
// 		d = 10*d+val
// 	}

// 	var out int =c+d
// 	var new []int
// 	for out>1{
// 		k = out%10
// 		new = append(new, k)
// 		out = out/10
// 	}

// 	fmt.Println(new)
// }
// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
//  func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

// 	carry := 0
// 	head := new(ListNode)
// 	cur := head
// 	for l1 != nil || l2 != nil || carry != 0 {
// 		n1, n2 := 0, 0
// 		if l1 != nil {
// 			n1, l1 = l1.Val, l1.Next
// 		}
// 		if l2 != nil {
// 			n2, l2 = l2.Val, l2.Next
// 		}
// 		num := n1 + n2 + carry
// 		carry = num / 10
// 		cur.Next = &ListNode{num % 10, nil}
// 		cur = cur.Next
// 	}
// 	return head.Next
// }