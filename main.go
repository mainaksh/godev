package main

import "fmt"

func main() {
	fmt.Println("hello go")

	foo()
	for i:=0; i<100; i++{
		if i%2==0{
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("i m in foo")
}

func bar(){
	fmt.Println("in bar exiting from main")
}