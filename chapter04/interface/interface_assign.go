package main

import (
	"fmt"
	"reflect"
)

type Number1 interface {
	Equal(i int) bool
	LessThan(i int) bool
	MoreThan(i int) bool
}
type Number2 interface {
	Equal(i int) bool
	MoreThan(i int) bool
	LessThan(i int) bool
	Add(i int)
}
type Number int

func (n Number) Equal(i int) bool {
	return int(n) == i
}

func (n Number) LessThan(i int) bool {
	return int(n) < i
}

func (n Number) MoreThan(i int) bool {
	return int(n) > i
}

func (n *Number) Add(i int) {
	*n = *n + Number(i)
}

type Num1 interface {
	Equal(i int) bool
	LessThan(i int) bool
	MoreThan(i int) bool
}
type Num2 interface {
	Equal(i int) bool
	LessThan(i int) bool
	MoreThan(i int) bool
}

type Num int

func (n Num) Equal(i int) bool {
	return int(n) == i
}
func (n Num) LessThan(i int) bool {
	return int(n) < i
}
func (n Num) MoreThan(i int) bool {
	return int(n) > i
}

func main() {
	var number1 Number = 1
	var number2 Number2 = &number1
	if number3, ok := number2.(Number1); ok {
		fmt.Println(number3.Equal(1))
		fmt.Println(reflect.TypeOf(number3))
	}
	var num Num = 1
	var num1 Num1 = num
	if num2, ok := num1.(Num2); ok {
		fmt.Println(num2.Equal(1))
	}
}
