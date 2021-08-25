package main

import "fmt"

type Integer int

type Math interface {
	Add(i Integer) Integer
	Multiply(i Integer) Integer
}

func (a Integer) Equal(b Integer) bool {
	return a == b
}

func (a Integer) Add(b Integer) Integer {
	return a + b
}

func (a Integer) Multiply(b Integer) Integer {
	return a * b
}

type person struct {
	name string
	age  uint
	addr address
}

type address struct {
	province string
	city     string
}

func (p person) String() string {

	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
}

type Stringer interface {
	String() string
}

func main() {
	p := person{
		age:  5,
		name: "Golang",
		addr: address{
			province: "北京",
			city:     "北京",
		},
	}
	defer func() {
		fmt.Println(p.addr.province)
	}()
	var x Integer
	var y Integer
	x, y = 10, 15
	fmt.Println(x.Equal(y))
	//
	//var a Integer = 1
	//var m Math = &a
	//fmt.Println(m.Add(1))
}
