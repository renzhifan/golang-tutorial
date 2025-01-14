package main

import (
	"fmt"
	"os"
	"unsafe"
)

func main() {
	name := "Hello Golang"
	nameP := &name
	fmt.Println("name变量的值为:", name)
	fmt.Println("name变量的内存地址为:", nameP)
	fmt.Println("name变量的值为:", *nameP)
	os.Exit(1)
	ii := 10
	iii := 11
	pp := &ii
	ppp := &iii
	ppp, pp = pp, ppp
	fmt.Println(pp)
	fmt.Println(*pp)
	//panic("dd");
	a := 100
	var ptr *int // 声明指针类型
	ptr = &a     // 初始化指针类型值为变量 a
	fmt.Println(ptr)
	fmt.Println(*ptr)

	// 通过指针传值
	b := 200
	swap(&a, &b)
	fmt.Println(a, b)

	// 指针类型转化
	i := 10
	var p *int = &i

	var fp *float32 = (*float32)(unsafe.Pointer(p))
	*fp = *fp * 10
	fmt.Println(i) // 100

	// 指针运算
	arr := []int{1, 2, 3}
	ap := &arr

	sp := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ap)) + unsafe.Sizeof(arr[0])))
	*sp += 3
	fmt.Println(arr)
}

func swap(a, b *int) {
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}
