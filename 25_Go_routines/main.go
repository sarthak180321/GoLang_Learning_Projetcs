package main
import(
	"fmt"
	"time"
)


func Greetings(){
	fmt.Println("hello from Go routine")
}
func main(){
	fmt.Println("hello from main")
	go Greetings()  //threading in Go lang is done through Go routines
	time.Sleep(3*time.Second)
	fmt.Println("i guess greeting is done")
}
//write Go infront of calls to to make it  a seprate Go routine
