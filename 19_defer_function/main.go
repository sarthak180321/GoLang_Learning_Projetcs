package main

import (
	"fmt"
)
func closeFile(){
	fmt.Println("closing file")
}
func main(){
	fmt.Println("hello")
	defer closeFile()
	defer fmt.Println("run in last in first out order")
	fmt.Println("openeing file ")
	fmt.Println("wrote into file")
	fmt.Println("saved the file")
	fmt.Println("task done")
}