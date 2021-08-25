package main

import "fmt"

type Age uint64

func (age Age) String() {
	fmt.Println("the age is", age)
}
func (age *Age) Modify() {
	*age = Age(30)
}
func main() {
	age := Age(25)
	age.String()
	(&age).Modify()
	//age.String()
}
