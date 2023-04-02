package main

import "fmt"
const constVariable string = "hhadddsdfja"

func main() {
	var username string = "variables"
	fmt.Println(username)
	fmt.Printf("type of variable: %T \n", username)

	var number uint8 = 255
	fmt.Println(number)
	fmt.Printf("the type of variable is : %T \n",number)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("the type of variables is : %T \n",isLoggedIn)

	var numFloat float64 = 79.764528828789498798228947
	fmt.Println(numFloat)

	//implicit type
	var name = "this is my college"
	fmt.Println(name)


	//no var style
	rollno := 57575755757
	fmt.Println(rollno)

	//default values and some aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("the type of variable : %T \n",anothervariable)

	//constant variable
	fmt.Println(constVariable)


}