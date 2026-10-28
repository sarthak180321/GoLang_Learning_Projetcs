package main

import (
	"fmt"
)

// *--> de refernce operator
// &-->address or reference operator

func update(score *int){
	*score=100
}
func main(){
	fmt.Println("hello")
	score:=10
	update(&score)
	fmt.Println(score)
}