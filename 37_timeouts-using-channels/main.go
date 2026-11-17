package main 
import(
	"fmt"
	"time"

)
func main(){
	fmt.Println("Hello, World!")
	ch:=make(chan string)
	go func(){
		time.Sleep(5*time.Second)
		ch<-"Data from channel"
	}()
	//time.after() is a function which returns a channel which will send the current time after the specified duration
	select {
	case data:= <-ch:
		fmt.Println("Data received from channel:",data)
	case <-time.After(3*time.Second):
		fmt.Println("Timeout occurred")
	}


}