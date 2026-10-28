package main   

import (
	"fmt"
)

func add(a int,b int) int{
	return a+b
}
func multi(a int, b int) int {
	return a*b
}
func myfunc(a int,b int,operation func(a int,b int) int) int{
	return operation(a,b)
}
func main(){
	fmt.Println("this is callback and functions as paramter fot another functions")
	fmt.Println(myfunc(10,20,add))
}