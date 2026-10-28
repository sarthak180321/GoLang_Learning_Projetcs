package main

import (
	"errors"
	"fmt"
)
func divide(a int,b int) (int,error){
	if b==0 {
		return 0,errors.New("you cant send 0 as divisor")
	} else {return a/b,nil}

}
func main(){
	result,err:=divide(10,0);
	if err==nil {
		fmt.Println(result)
	} else {
		fmt.Println(err.Error())
	}
	fmt.Println("hello")
}