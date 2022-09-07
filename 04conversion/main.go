package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter rating:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating:", input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Something went wrong", err)
	} else {
		fmt.Println("Adding 1 to rating:", numRating+1)
	}
}
