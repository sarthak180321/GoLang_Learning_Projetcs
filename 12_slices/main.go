package main  
// inside The slices hsa Three components 
//--> pointer to first element + size + capacity
//-->make() initialize arry with (type,size,capacity) 
import (
	"fmt"
	"slices"
)
func main(){
	var nums []int

	nums=append(nums, 10,20)
	nums=append(nums, 40)
	fmt.Println(nums)

	myarray:=[]string{"string","sarthak"} // initialization
	fmt.Println(myarray)

	nums2:=make([]int,3,5);
	fmt.Println(nums2)
	nums=slices.Delete(nums,0,1)
	fmt.Println(nums)


	nums3:=nums[1:2]
	fmt.Println(nums3)
	copy(nums3,nums[1:2])
	fmt.Println(nums3)

}