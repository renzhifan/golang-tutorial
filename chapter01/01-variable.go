package main

//变量声明和命名规则

import "fmt"

func main() {
	var v1 int      // 整型
	var v2 string   // 字符串
	var v3 bool     // 布尔型
	var v4 [10]int  // 数组，数组元素类型为整型
	var v5 struct { // 结构体，成员变量 f 的类型为64位浮点型
		f float64
	}
	var v6 *int            // 指针，指向整型
	var v7 map[string]int  // map（字典），key为字符串类型，value为整型
	var v8 func(a int) int // 函数，参数类型为整型，返回值类型为整型

	v9 := 2
	v1, v9 = v9, v1
	// 打印上述变量值
	fmt.Printf("v1: %v\n", v1)
	fmt.Printf("v2: %v\n", v2)
	fmt.Printf("v3: %v\n", v3)
	fmt.Printf("v4: %v\n", v4)
	fmt.Printf("v5: %v\n", v5)
	fmt.Printf("v6: %v\n", v6)
	fmt.Printf("v7: %v\n", v7)
	fmt.Printf("v8: %v\n", v8)
	fmt.Printf("v9: %v\n", v9)

	//匿名变量
	_, nickName := GetName()
	fmt.Printf("nickName: %v\n", nickName)
	//变量赋值与多重赋值
	v11 := 11
	v12 := 12
	v12, v11 = v11, v12
	fmt.Printf("v11: %v\n", v11)
	fmt.Printf("v12: %v\n", v12)
}

func GetName() (userName, nickName string) {
	return "aaa", "bbb"
}
