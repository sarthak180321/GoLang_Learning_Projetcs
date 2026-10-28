package main 

//arrays are fixed size and cannot be changed 
import (
	"fmt"
)
//declration syntax
//var name [size]type
func main(){
	var nums [5]int;
	nums[1]=10;  //by default all zero
	fmt.Println(nums);

	//array literals
	arr:=[3]int{1,2,3}
	fmt.Println(arr)

	fmt.Println(len(nums))
	// lenght of the array

}