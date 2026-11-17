package main

import(
	"fmt"
	"time"
	"sync"
)
func consume(ch chan int,wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("consuming from channel")
	fmt.Println("Hault started")
	fmt.Println(<-ch)
	fmt.Println("Hault ended")
}
func produce(ch chan int,wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("producing to channel")
	ch<-10
}
func main(){

	fmt.Println("Hello, World!")
	//what is a buffered channel?
	// when You enter data in channel and channel capacity is exceeded 
	//then Go will hault that routine until some other routine reads from that channel
	//similary when you read from a channel and channel is empty then GO will hault until data is found in the channel
	//lets see this in action
	var wg sync.WaitGroup
	wg.Add(2)
	ch:=make(chan int)
	fmt.Println(ch)
	go consume(ch,&wg)
	time.Sleep(3*time.Second)
	go produce(ch,&wg)
	wg.Wait()
}