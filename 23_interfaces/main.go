package main

import (
	"fmt"
)
// to declare a interfce type interfaceName interface{}


type speaker interface{
	Speak() string
}

type sony struct{}
type boat struct{}
func (u sony) Speak() string{
	fmt.Println("this sony speaker speaking")
	return "sony"
}
func (u boat) Speak() string{
	fmt.Println("this is boat speaker speakinh")
	return "boat"
}

func Bark(s speaker) {
	fmt.Println(s.Speak())
}
func main(){
	fmt.Println("hello")
	Bark(sony{})
	Bark(boat{})
}
 