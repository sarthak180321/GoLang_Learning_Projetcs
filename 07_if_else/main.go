package main

import (
	"fmt"
)

func main(){
	score:=72

	if score>=90 {
		fmt.Println("A grade with score ",score)
	} else if score>=80{
		fmt.Println("B grade with score",score)
	} else{
		fmt.Println("not Good with Grade less than B and score",score)
	}
}