package main

import "fmt"

func main() {
	fmt.Println("Hello")
	myNumber := 23
	var myPointer = &myNumber
	fmt.Println("My pointer", myPointer)
	fmt.Println("My pointer", *myPointer)

	*myPointer = *myPointer + 10
	fmt.Println("New Value of my pointer: ", *myPointer);

}
