package main

import (
	"fmt"
	"math"
)

func main() {
	var floatValue1 float64
	floatValue1 = 10.0
	floatValue2 := 10.0
	var boolValue bool
	boolValue = floatValue1 == floatValue2
	floatValue4 := 0.1
	floatValue5 := 0.7
	floatValue6 := floatValue4 + floatValue5

	fmt.Println(floatValue1)
	fmt.Println(boolValue)
	fmt.Println(floatValue6)

	//浮点数的比较
	p := 0.00001
	if math.Dim(floatValue1, floatValue2) < p {
		fmt.Println("floatValue1 和 floatValue2 相等")
	}
}
