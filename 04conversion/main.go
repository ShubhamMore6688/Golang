package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Give the rating to our app between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for giving a rating, ",input)

	numReader, err := strconv.ParseFloat(strings.TrimSpace(input),64)
	
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("Adding 1 in rating: ", numReader+1)
	}
}