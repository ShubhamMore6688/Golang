package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string
	fruitList[0] = "mango"
	fruitList[1] = "Banana"
	fruitList[2] = "Papaya"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegyList = [4]string{"potato","beans","veg"}
	fmt.Println(vegyList)

}