package main

import (
	"errors"
	"fmt"
)

func add(a, b *int) int {
	*a *= 2
	*b *= 3
	return *a + *b
}

func sum(a, b int) int {
	return a + b
}

//多值返回
func sum2(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能是负数")
	}
	return a + b, nil
}

//命名返回参数
func sum3(a, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能是负数")
	}
	sum = a + b
	err = nil
	return
}

//可变参数
func sumAll(params ...int) int {
	sum := 0
	for _, i := range params {
		sum += i
	}
	return sum

}

//闭包函数

func colsure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	x, y := 1, 2
	z := add(&x, &y)
	fmt.Printf("add(%d, %d) = %d\n", x, y, z)
	a := sum(1, 2)
	fmt.Println(a)
	result, err := sum2(-11, 22)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	result1, err := sum3(1, 3)

	fmt.Println(result1, err)

	fmt.Println(sumAll(1, 2, 3))

	//匿名函数
	sumNil := func(a, b int) int {
		return a + b
	}
	fmt.Println(sumNil(1, 2))
}
