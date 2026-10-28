package main
 
// when you run a Go file Go looks for main() function and package main
// package main tells that this is executabe file but not resuable	

// file having package main should have main() function 
// package main file are compiled as binary executable like .exe instead of .a archive 
// these are executable not resuable resources 

import "fmt"
 
func main(){
	fmt.Println("hell Go")

}