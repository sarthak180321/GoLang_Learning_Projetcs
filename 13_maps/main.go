package main

import (
	"fmt"
	"maps"
)

func main(){
	//literal mao
	colors:=map[string]string{
		"red":"#ff0000",
		"green":"#00ff00",
	}
	fmt.Println(colors)

	ages:=make(map[string]int);//--> initialize empty map
	fmt.Println(ages) 
	ages["12"]=14
	fmt.Println(ages)
	fmt.Println(ages["ayush"]) //-->returns zero if not present 
	fmt.Println(len(ages)) //-->size of map
	clear(ages) //--> clear map
	_,ok:=ages["alice"]
	fmt.Println(ok);
	ages2:=make(map[string]int)
	fmt.Println(maps.Equal(ages,ages2))
	for key,value := range colors{
		fmt.Println(key,value)
	}
}