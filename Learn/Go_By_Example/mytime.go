package main

import (
	"fmt"
	"time"
)

func mainI() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(time.Now().Year(), time.July, 10, 23, 23,0,0, time.UTC)
	
	fmt.Println(createdDate.Format("01/02/2006 Monday"))

}
