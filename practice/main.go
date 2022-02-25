package main

import "fmt"

var bString string = "Another Go variable"

const aConst int = 63

func main() {
	aString := "This is Go!"
	fmt.Println(aString)
	fmt.Printf("The variable type is %T\n", bString)
	fmt.Println(bString)

	// Assigning variables
	anInt := 43
	fmt.Println(anInt)

	var defaultInt int
	fmt.Println(defaultInt)

	anotherString := "This is another string!"
	fmt.Println(anotherString)
	fmt.Printf("Variable type is %T\n", anotherString)

	var anotherInt = 53
	fmt.Println(anotherInt)
	fmt.Printf("Variable type %T\n", anotherInt)

	fmt.Printf("Variable type %T\n", aConst)
	fmt.Println(aConst)
}
