package main 

import(
	"fmt"
);

func main(){
	items:=4;
	itemPrice:=49;
	if total:=itemPrice*items; total>=1000{
		fmt.Println("You declraed and also checked it wow this is amazing");
	}else{
		fmt.Println("still checked and ran wow");
	}
}