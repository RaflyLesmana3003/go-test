package main

import (
	"fmt"

	"main.go/helper"
)

func main() {
	fmt.Println("=================")
	fmt.Println("sorting =================")
	fmt.Println("=================")
	helper.InitializeData()

	fmt.Println("=================")
	fmt.Println("multiple sorting =================")
	fmt.Println("=================")
	helper.InitializeDataMulti()

	fmt.Println("=================")
	fmt.Println("palindrome =================")
	fmt.Println("=================")
	helper.InitializePalindrome()

}
