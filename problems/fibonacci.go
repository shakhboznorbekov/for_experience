package main 

import (
	"fmt"
)

// func main (){
// 	var a, b, c, limit int

// 	fmt.Println("Fibonacci number limit:")
// 	fmt.Scan(&limit)

// 	b = 1

// 	for i := 0; i < limit; i++ {

// 		fmt.Printf("%d ", a + b)
// 		c = a + b
// 		a = b
// 		b = c
// 	}
// }

// func fib(N int) int {
// if N < 2 {
// 	return N
// }

// return fib(N-1) + fib(N-2)
// }

SELECT id
SUM(IF(month = "Jan", revenue, NULL)) AS Jan_Revenue,
SUM(IF(month = "Feb", revenue, NULL)) AS Feb_Revenue,
SUM(IF(month = "Mar", revenue, NULL)) AS Mar_Revenue,
SUM(IF(month = "Apr", revenue, NULL)) AS Apr_Revenue,
SUM(IF(month = "May", revenue, NULL)) AS May_Revenue,
SUM(IF(month = "Jun", revenue, NULL)) AS Jun_Revenue,
SUM(IF(month = "Jul", revenue, NULL)) AS Jul_Revenue,
SUM(IF(month = "Aug", revenue, NULL)) AS Aug_Revenue,
SUM(IF(month = "Sep", revenue, NULL)) AS Sep_Revenue,
SUM(IF(month = "Oct", revenue, NULL)) AS Oct_Revenue,
SUM(IF(month = "Nov", revenue, NULL)) AS Nov_Revenue,
SUM(IF(month = "Dec", revenue, NULL)) AS Dec_Revenue
FROM Department 
ORDER BY id
