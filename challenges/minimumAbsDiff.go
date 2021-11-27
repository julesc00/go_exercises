package main

import (
	"fmt"
	"sort"
)

// var arr1 = []int32{-1, 5, 10, 20, 28, 3}
// var arr2 = []int32{26, 134, 135, 15, 17}

func minAbsDiff(arr []int32) int32 {
	var f foo
	f = arr
	temp := Abs(f[0] - f[1])

	sort.Sort(f)
	fmt.Println(f)

	for i := 0; i < len(f) - 1; i++ {
		if Abs(f[i] - f[i+1]) < temp {
			temp = Abs(f[i] - f[i+1])
		}
	}
	return int32(temp)
}

// Abs returns the absolute value of x
func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

//type foo that implements Sort for int32 type
type foo []int32

func (f foo) Len() int {
	return len(f)
}

func (f foo) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f foo) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
