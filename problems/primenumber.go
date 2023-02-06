package main

import (
	"fmt"
)

func main (){
	var a,b int
	fmt.Scanf("%d", &a)
	for i:=2; i<a; i++{
		if a%i==0{
			b++
		}
	}
	if b>=1 && a>1{
		fmt.Println("prime number")
	}else{
		fmt.Println("not prime numner")
	}
}