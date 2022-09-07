package main

import (
	"crypto/rand"
	"math/big"

	// "math/rand"

	"fmt"
)

func main() {

	// Seed for "math/rand"
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(523232323))
	fmt.Println("myRandomNum", myRandomNum)
}
