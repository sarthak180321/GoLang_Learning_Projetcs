package main

import (
	"fmt"
)
type Operation func(int,int) int

func process(a int,b int,operation Operation) int{
	result:=operation(4,3);
	return a*b*result
}
func main(){
	fmt.Println("hello")
	fmt.Println(process(10,20,func(a int,b int)int{
		return a+b
	}))
}