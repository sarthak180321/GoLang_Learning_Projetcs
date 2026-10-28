package main
import (
	"fmt"
)
//Questions aroused --> dynamic size channels?
func main(){
	fmt.Println("concurrency models and DEsign patterns in Go")
	subs:=[]chan string{
		make(chan string,10),
		make(chan string,10),
		make(chan string,10),
	}
	for i,c :=range subs{
		go func(id int,ch chan string){
			for msg:=range ch{
				fmt.Printf("consumer %d: %s\n", id, msg)
			}
		}(i,c)
	}
	msg := "hello"
for _, c := range subs {
    c <- msg
}
}