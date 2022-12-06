package main

import (
	"fmt"
	"strings"
)

func practiceFunc(name string) (int, string) {
	// defer can execute the code that function finished !!
	defer fmt.Println("I'm done!")
	return len(name), strings.ToUpper(name)
}
// This is "naked" return
func nakedReturn(name string) (length int, upper string) {
	length = len(name)
	upper = strings.ToLower(name)
	return
}

func main() {
	A := "george"
	fmt.Println(practiceFunc(A))
	fmt.Println(nakedReturn(A))
}
