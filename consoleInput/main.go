package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Catch from text input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Your entered: ", input)
}
