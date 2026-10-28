package main 
import(
	"fmt"
	"errors"
)


func Max[T int|float64](a T, b T) (T , error){
	if(a <0 || b<0) {
		return 0,errors.New("this is not possible")
	} else if(a>b){
		return a,nil
	} else if(a==b){
		return a,nil
	} else{
		return b,nil
	}
}


type taker interface{
	int | float32 
}

func less[T taker](nums []T) ([]T){
	return nums
}
func main(){
	fmt.Println("hello Generics")
	fmt.Println(less(make([]int,3,5)))
}
