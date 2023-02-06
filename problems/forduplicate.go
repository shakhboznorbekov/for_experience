package main 

import (
	"fmt"
)

func main() {
	var a int

	numbers :=[]int{1,2,3,4,5,6,7,7,4,8,8,32}

	for n, i :=range numbers {
		for j:=n; j<len(numbers)-1; j++ {
			if i==numbers[j+1]{
				a += 1
			}
		}
	}
	if a>=1 {
		fmt.Println("Duplicate array")
		return
	}
	fmt.Println("Not duplicate array")

}
	