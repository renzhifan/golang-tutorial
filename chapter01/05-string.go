package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 转义字符
	results := "Search results for \"Golang\":\n" +
		"- Go\n" +
		"- Golang\n" +
		"- Golang Programming\n"
	fmt.Printf("%s", results)

	// 字符串切片
	str := "0123456789"
	str1 := str[:5]  // 获取索引5（不含）之前的子串
	str2 := str[7:]  // 获取索引7（含）之后的子串
	str3 := str[0:5] // 获取从索引0（含）到索引5（不含）之间的子串
	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)
	fmt.Println("str3:", str3)

	str4 := str[:]
	fmt.Println("str4:", str4)

	// 字符编码
	msg := "Hello, 世界"
	n := len(msg) //  获取到的是 msg 的字节长度
	for i := 0; i < n; i++ {
		ch := msg[i] // 依据下标取字符串中的字符，ch 类型为 byte
		fmt.Println(i, ch, reflect.TypeOf(ch))
	}

	for i, ch := range msg {
		fmt.Println(i, string(ch), reflect.TypeOf(ch)) // 通过 range 遍历，ch 类型是 rune
	}

	//字符串和其他基本类型之间的转化
	//将整型转化为字符串
	v1 := 65
	v2 := string(v1) // v2 = A
	fmt.Println("v2:", v2)
	v3 := 30028
	v4 := string(v3) // v4 = 界
	fmt.Println("v4:", v4)
	//将 byte 数组或者 rune 数组转化为字符串
	v5 := []byte{'h', 'e', 'l', 'l', 'o'}
	v6 := string(v5)
	fmt.Println("v6:", v6)
	v7 := []rune{0x5b66, 0x9662, 0x541b}
	v8 := string(v7)
	fmt.Println("v8:", v8)

}
