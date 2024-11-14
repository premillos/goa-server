package main

import (
	"fmt"

	"com.goa/pkg/utils"
)

func main() {
	fmt.Println("hello world")

	res, _ := utils.GenerateRandomInteger(10)

	fmt.Println("res", res)
}
