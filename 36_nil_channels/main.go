package main
import(
	"fmt"
)
func main(){
	fmt.Println("Hello, World!")
	// We declared a channel but didn't initialize it using make().

// Its value is nil.

// Now look at these examples.
	var ch chan int 
	fmt.Println(ch)
	//ch<-10 // this will cause a deadlock because channel is nil and no one is reading from it
	//fmt.Println(<-ch) // this will also cause a deadlock because channel is nil and no one is writing to it
	//close(ch) // this will cause a panic because channel is nil and you can't close a nil channel
	

}