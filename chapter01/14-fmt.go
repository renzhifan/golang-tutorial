package main

import "fmt"
import "os"

type point struct {
	x, y int
}

func main() {
	//Print系列函数会将内容输出到系统的标准输出，区别在于Print函数直接输出内容，Printf函数支持格式化输出字符串，Println函数会在输出内容的结尾添加一个换行符。
	// 打印结构体
	p := point{1, 2}
	fmt.Printf("%v\n", p) //输出结果为 {1 2}

	// 如果值是一个结构体，%+v 变体将包括结构体的字段名。
	fmt.Printf("%+v\n", p) //输出结果为 {x:1 y:2}

	// %#v 变体打印值的 Go 语法表示，即将生成该值的源代码片段。
	fmt.Printf("%#v\n", p) //输出结果为 main.point{x:1, y:2}

	// 打印类型
	fmt.Printf("%T\n", p) //输出结果为 main.point

	// 打印布尔值
	fmt.Printf("%t\n", true) //输出结果为 true

	// 打印整数。
	fmt.Printf("%d\n", 123) //输出结果为 123

	// 打印二进制
	fmt.Printf("%b\n", 14) //输出结果为 1110

	// 打印字符
	fmt.Printf("%c\n", 33) //输出结果为 ！

	// 打印 16 进制
	fmt.Printf("%x\n", 456) //输出结果为 1c8

	// 打印浮点数
	fmt.Printf("%f\n", 78.9) //输出结果为 78.900000
	// 指数型
	fmt.Printf("%e\n", 123400000.0) //输出结果为 1.234000e+08
	fmt.Printf("%E\n", 123400000.0) //输出结果为 1.234000E+08

	// 字符串
	fmt.Printf("%s\n", "\"string\"") //输出结果为 "string"
	// 双引号字符串.
	fmt.Printf("%q\n", "\"string\"") //输出结果为 "\"string\""

	// 每个字符用两位 16 进制表示。
	fmt.Printf("%x\n", "hex this") //输出结果为 6865782074686973

	// 打印指针`.
	fmt.Printf("%p\n", &p) //输出结果为 0xc00001c08

	// 字符宽度
	fmt.Printf("|%6d|%6d|\n", 12, 345) //输出结果为 |    12|   345|

	//字符精度
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45) //输出结果为 |  1.20|  3.45|

	// 左对齐
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) //输出结果为 |1.20  |3.45  |

	//同样可以控制字符的宽度
	fmt.Printf("|%6s|%6s|\n", "foo", "b") //输出结果为 |   foo|     b|

	// 同样字符左对齐.
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") //输出结果为 |foo   |b     |

	// 合并
	s := fmt.Sprintf("a %s", "string") //输出结果为 a string
	fmt.Println(s)
	// 同样的效果
	fmt.Fprintf(os.Stderr, "an %s\n", "error") //输出结果为 an error
}
