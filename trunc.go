package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Please enter a floating point number:\n")
	fmt.Scanln(&num)
	fmt.Print("Truncated integer:\n", num)
}

//Output:-
/*
go run trunc.go
Please enter a floating point number:
20.3
Truncated integer:
20
*/
