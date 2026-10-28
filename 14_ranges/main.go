package main  
import (
	"fmt"
)
func main(){
	nums:=[]int{6,7,8,9}
	fmt.Println(nums)
	for index,value := range nums{
		fmt.Println(index,"this is my index")
		fmt.Println(value,"this is my value")
	}
}