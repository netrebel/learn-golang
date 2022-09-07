package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Welcome 👏🏻")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Pizza rating:")

	user, _ := reader.ReadString('\n')

	fmt.Println("Rating:", user)

}
