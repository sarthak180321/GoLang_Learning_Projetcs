package main 
import (
	"fmt"
	"sync"
)
func work(id int,wg *sync.WaitGroup){
	defer wg.Done() //-1 counter when function is about to finish
	fmt.Println("worker",id)
}
func main(){
	fmt.Println("hello from waitGroups")
	var wg sync.WaitGroup
	for i:=1; i<3; i++ {
		wg.Add(1) // increases wg value
		go work(i,&wg) //call with counter+1 and pass same wg Groyp reference
	}
	wg.Wait() //waits till counter is Zero

}