package main

import "fmt"

func main() {
	fmt.Println("Hello in Array")
	
	var myList [4]string
	myList[0] = "Deepak"
	myList[1] ="SM"
	myList[3] = "Son"

	fmt.Println("myList value", myList)
	fmt.Println("Length of myList", len(myList))

	var VegList = [5]string{"D","M","K"}
	fmt.Println("VegList", VegList);
	



}
