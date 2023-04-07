package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study in golang")
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime.Format("01-02-2006 Monday 15:04:05"))
	createDate := time.Date(2003, time.August, 06,5,5,6,6, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday 15:04:05"))
}