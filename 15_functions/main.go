package main 

import (
	"fmt"
)
func add(a int,b int) (int , int){
	return a+b,a*b
}
func main(){
	fmt.Println(add(10,20))
	a,b:=add(1,2)
	fmt.Println(a,b)
}