package main

import (
	"fmt"
)
type Employe struct{
	name string
	age  int
	Salary float64
	IsRemote bool
}

//conmcetion functions to structs 
func (e Employe) updateName(s string){
	e.name=s
}
// by value called 

func (e* Employe) update(s string){
	e.name=s
}
// by reference called Will update and reflect changes

func main(){
	fmt.Println("hello")
	e:=Employe{"sarthak",23,40000,true};
	fmt.Println(e)
	e.updateName("sarthak arora")
	fmt.Println(e)
	e.update("ayush")
	fmt.Println(e)
}