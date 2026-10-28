package main 

import (
	"fmt"
)

func main(){
	score:=3

	switch score {
		case 1:
			fmt.Println(score)
		case 2:
			fmt.Println(score)
		case 3:
			fmt.Println(score)
		default:
			fmt.Println("score higher  then expected")
	}
	
}