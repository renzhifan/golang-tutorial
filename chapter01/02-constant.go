package main

//常量
import "fmt"

const ( // iota 被重置为 0
	c0 = iota // c0 = 0
	c1        // c1 = 1
	c2        // c2 = 2
)

const (
	u = iota * 2 // u = 0
	v            // v = 2
	w            // w = 4
)

const x = iota // x = 0
const y = iota // y = 0

//枚举
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays
)

func main() {
	fmt.Printf("x:%v\n", x)
	fmt.Printf("y:%v\n", y)

	fmt.Printf("Sunday: %v\n", Sunday)
	fmt.Printf("Monday: %v\n", Monday)
	fmt.Printf("Tuesday: %v\n", Tuesday)
	fmt.Printf("Wednesday: %v\n", Wednesday)
	fmt.Printf("Thursday: %v\n", Thursday)
	fmt.Printf("Friday: %v\n", Friday)
	fmt.Printf("Saturday: %v\n", Saturday)
	fmt.Printf("numberOfDays: %v\n", numberOfDays)
}
