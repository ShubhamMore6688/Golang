package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give the rating to our pizza :")

	//comma ok // error syntax (like try catch)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ",input)
	fmt.Printf("the type of input is %T ", input)
	
}