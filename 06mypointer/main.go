package main

import "fmt"

func main() {
	fmt.Println("Welcome to lecture on pointer")

	myNum := 10
	var ptr = &myNum

	fmt.Println("value of add of mynum", ptr)
	fmt.Println("value of mynum", *ptr)
	*ptr = *ptr *2
	fmt.Println(myNum)

}