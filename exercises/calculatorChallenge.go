package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// A simple adding calculator
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter two numbers to add\n")
	fmt.Print("Enter a number: ")
	input1, _ := reader.ReadString('\n')
	fmt.Print("Enter another number: ")
	input2, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	float2, err2 := strconv.ParseFloat(strings.TrimSpace(input2), 64)

	if err != nil || err2 != nil {
		fmt.Println(err)
		panic("Value must be a number!")
	} else {
		floatSum := float1 + float2
		floatSum = math.Round(floatSum*100) / 100
		fmt.Printf("Your sum of %v + %v = %v \n", float1, float2, floatSum)
	}
}
