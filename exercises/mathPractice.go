package main

import (
	"fmt"
	"math"
)

func main() {
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("The sume is: ", intSum)

	f1, f2, f3 := 25.3, 77.20, 87.40
	floatSum := f1 + f2 + f3
	fmt.Println("The float sum is: ", floatSum)

	// for rounding to 2 decimals, though I think there should be another way
	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("Rounded float sum is: ", floatSum)

	// Getting a radius, and rounding to 2 decimals.
	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	circumference = math.Round(circumference*100) / 100
	fmt.Printf("Circumference: %.2f\n", circumference)
}
