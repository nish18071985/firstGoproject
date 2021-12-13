package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Hello Go!!!")

	firstNum := strconv.Itoa(10)
	secondNum := "15"

	fmt.Println(firstNum + secondNum)
	fmt.Println(math.Max(1, 2))
}
