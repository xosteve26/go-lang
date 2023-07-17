package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time functions")

	presentTime := time.Now()

	fmt.Println(presentTime.Format(" Monday Jan 2 15:04 2006 MST"))

	createdDate := time.Date(2020, time.April, 1, 3, 3, 21, 0, time.Local)

	fmt.Println(createdDate.Format("Monday Jan 02 15:04 2006 MST"))

}
