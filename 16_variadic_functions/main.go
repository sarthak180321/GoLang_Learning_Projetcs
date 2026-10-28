package main

import "fmt"

// -->named functions
func addition(a int,b int) (result int){
	result=a+b
	return result
	// -->this helps in Good Code readability
}


func add(numbers ...int) int {
	sum:=0
	for _,number := range numbers{
		sum=sum+number
	}
	return sum
}



func main(){
	a:=addition(10,20)
	fmt.Println(a)
	fmt.Println(add(10,20,30,40))
	numbers:=[]int{10,20,30,40}
	fmt.Println(add(numbers...))

	myfunction:=func(a int,b int) int{
		return a*b
	}
	fmt.Println(myfunction(10,20))

}