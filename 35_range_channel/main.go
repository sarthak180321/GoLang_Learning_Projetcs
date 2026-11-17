package main
import(
	"fmt"
)
func main(){
	fmt.Println("Hello, World!")
	ch:=make(chan int,3)
	fmt.Print(ch)
	ch<-10
	ch<-20
	ch<-30
	close(ch)
	//you dont have to do like this 
	//value:=<-ch
	//value1:=<-ch
	//value2:=<-ch
	//you can do like this
	for value := range ch{
		fmt.Println(value)
	}

}