package main
import (
	"fmt"
)
func main(){
	fmt.Println("this  is channels")
	chchannela:=make(chan string)
	go func(){
		chchannela<-"go channel data pushed" //pushin data into pipe
	}()
	msg:=<-chchannela //getting data from pipe
	fmt.Println(msg)
}