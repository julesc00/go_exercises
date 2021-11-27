package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("struct1:  %v\n", p)  		// struct1:  {1 2}
	fmt.Printf("struct2:  %+v\n", p)  	// struct2:  {x:1 y:2}
	fmt.Printf("struct3:  %#v\n", p)  	// main.point{x:1, y:2}
	fmt.Printf("type:  %T\n", p)  		// type:  main.point
	fmt.Printf("bool:  %t\n", true)  	// bool:  true
	fmt.Printf("int:  %d\n", 123)  	// int:  123
	fmt.Printf("bin:  %b\n", 14)		// bin:  1110
	fmt.Printf("char:  %c\n", 33)		// char:  !
	fmt.Printf("hex:  %x\n", 456)		// hex:  1c8
	fmt.Printf("Float1: %f\n", 78.9)	// Float1: 78.900000
	// Float scientific notation
	fmt.Printf("float2:  %e\n", 123400000.0)  // float2:  1.234000e+08
	fmt.Printf("float3: %E\n", 123400000.0) 	// float3: 1.234000E+08
	fmt.Printf("str1:  %s\n", "\"string\"")  	// str1:  "string"
	fmt.Printf("str2: %q\n", "\"string\"") 	// str2: "\"string\""
	fmt.Printf("str3: %x\n", "hex this")		// str3: 6865782074686973
	fmt.Printf("pointer: %p\n", &p)		// pointer: 0xc0000b2010
	fmt.Printf("width1:  |%6d|%6d|\n", 12, 345)  // width1:  |    12|   345|
	fmt.Printf("width2:  |%6.2f|%6.2f|\n", 1.2, 3.45)  // width2:  |  1.20|  3.45|
	fmt.Printf("%.2f\n", 3.4444)		// 3.44
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
	fmt.Printf("width5:  |%-6s|%-6s|\n", "foo", "b")	 // width4:  |foo   |b     |
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)		// sprintf: a string

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")  // io: an error
}