package main 
import(
	"fmt"
)

func main(){
	fmt.Println("Hello, World!")
	// suppose you have a N number of channels and You want to wait on them collectievly
	//like if any channel gets data you process and move ahead 
	// this makes Your code non Blocking and more efficient
	ch:=make(chan int)
	ch1:=make(chan int)
	ch<-10
	select {
	case data:=<-ch:
		fmt.Println("Data received from channel 1:",data)
	case data1:=<-ch1:
		fmt.Println("Data received from channel 2:",data1)
	default:
		fmt.Println("No data received from any channel")
	}
}
