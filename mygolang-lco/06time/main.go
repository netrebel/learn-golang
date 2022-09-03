package main

import (
	"fmt"
	"time"
)

func main() {

	presentTime := time.Now()
	fmt.Println("Time: ", presentTime)
	fmt.Println("Time: ", presentTime.Format("01-02-2006 15:04:05 Monday "))

	// create Date
	createdTime := time.Date(2002, time.May,3,3,3,3,3, time.UTC)
	fmt.Println("Created time: ", createdTime)
}
