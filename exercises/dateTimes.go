package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("Time now is: ", n)

	t := time.Date(2021, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Launched at: ", t)
	fmt.Println(t.Format(time.ANSIC))

	// _ to ignore the error that parse generates.
	parsedTime, _ := time.Parse(time.ANSIC, "Wed Nov 10 23:00:00 2021")
	fmt.Printf("Type of parsedTime is %T\n", parsedTime)
}
