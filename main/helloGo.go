package main

import (
	"fmt"
	"strconv"
)
import "unsafe"

// TODO 全局变量声明
var (
	n1 = 1
	n2 = "2"
)

func main() {
	//TODO 输出文本
	fmt.Println("Hello, World!")
	//声明空变量
	var name, email string
	fmt.Println(name, email)
	//赋值声明变量
	no, phone := 001, 10086
	fmt.Println(no, phone)
	//输出全局变量
	fmt.Println(n1, n2)

	//TODO 整数类型 int
	var _int int     //最大
	var _int8 int8   //范围 -128~127
	var _int16 int16 //大一些 正常计算够用 后面类推
	var _int32 int32
	var _int64 int64
	//赋值调用后如果变量超出范围会报错
	_int = 1
	_int8 = 8
	_int16 = 16
	_int32 = 32
	_int64 = 64
	fmt.Println(_int, _int8, _int16, _int32, _int64)
	//打印变量占用字节数 unsafe.Sizeof
	fmt.Println(unsafe.Sizeof(_int))

	//TODO 浮点型 float64
	var price32 float32 //4字节 范围不一样 精度缺失
	var price64 float64 //8字节 64精度比32的准确
	price32, price64 = 1.25, 1.43
	fmt.Println(price32, price64)

	//TODO 字符类型 char go语言统一使用utf-8不会有乱码问题
	//流程
	//字符->对应码值->二进制->存储
	//二进制->对应码值->字符->读取
	var bt byte
	bt = 'a' //byte类型 那几个字节 很小 存储单个字母
	//这里直接输出的是字符对应的码值 是ascii码对应的码值数字
	fmt.Println(bt)          //97
	var btString = "bt=%c\n" //bt变量按字符串格式输出
	fmt.Printf(btString, bt)
	//overflow溢出 导致的 字符串的数字格式输出 字符保存以数字类型保存的方式
	var stringByInt int = '北'
	var _stringByInt = "stringByInt=%d"
	fmt.Println(stringByInt, _stringByInt)
	var n1 = 10 + 'a' //10 + 97(a的ascii码是97) = 107
	fmt.Println(n1)

	//TODO 布尔类型 bool 只有 true false 其他替代不允许使用 占用字节1
	var _t = true
	_t = false
	var _f = false
	_f = true
	fmt.Println(unsafe.Sizeof(_t), unsafe.Sizeof(_f))

	//TODO 字符串类型 string  go语言的字符串的字节使用uft8编码表示Unicode文本
	//go中的字符串是不可以改变的 比如拿address中的 address[0] = '法' 是不被允许的
	var address = "中国移动 10086 hello !"
	fmt.Println(address)
	email = "邮箱:731005184@qq.com"
	fmt.Println(email)
	//反引号``字符串原生格式输出 原汁原味 可以防止攻击 以及 输出源码等效果
	supperString := `//TODO go中的字符串是不可以改变的 比如拿address中的 address[0] = '法' 是不被允许的
	var address = "中国移动 10086 hello !"`
	fmt.Println(supperString)
	//字符串拼接 使用 + 号
	stringA, stringB := "hello ", "word!"
	stringC := stringA + stringB + //拼接过长可以分行写 把+号留在末尾
		"haha~" +
		">>??~~!!!"
	fmt.Println(stringC)

	//TODO 变量数据转换 基本数据类型的互相转换不能自动转换 必须显示转换
	var i32 int = 100
	var f64 float64 = float64(i32)
	fmt.Printf("f64=%v \n", f64)

	//TODO 数据类型转换string类型
	//第一种方式
	var num1 = 99
	var num2 = 23.456
	var b = true
	b = false
	var h byte = 'h'
	var str string
	str = fmt.Sprintf("%d \n", num1)
	fmt.Printf("str type : %T str = %v \n", str, str)
	str = fmt.Sprintf("%f \n", num2)
	fmt.Printf("str type : %T str = %v \n", str, str)
	str = fmt.Sprintf("%t \n", b)
	fmt.Printf("str type : %T str = %q \n", str, str)
	str = fmt.Sprintf("%c \n", h)
	fmt.Printf("str type : %T str = %q \n", str, str)
	//第二种方式
	var num3 = 99
	num3 += 55
	var num4 = 23.456
	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str = %q \n\n", str, str)
	//f代表格式 10表示小数保留10位 64是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	str = strconv.FormatBool(b)

	//TODO string类型转为基本数据类型 数据必须是有效转换 不然会出问题
	var num5 = 4567
	str = strconv.Itoa(num5) //数字转字符串 但是变量本身还是数字类型
	fmt.Printf("str type %T value= %v \n", str, str)
	var stringBool1 = "true"
	var bs bool
	bs, _ = strconv.ParseBool(stringBool1)
	fmt.Printf("bs type %T value = %v \n", bs, bs)
	var stringInt1 = "123"
	var ints int64
	ints, _ = strconv.ParseInt(stringInt1, 10, 64)
	fmt.Printf("ints type %T value = %v \n", ints, ints)
	var f1 = "123.444"
	var ff float64
	ff, _ = strconv.ParseFloat(f1, 64)
	fmt.Printf("ff type %T value = %v \n", ff, ff)

	//TODO 指针 指针变量指向的是变量的一个地址 地址指向的空间值才是值
	var i = 10
	fmt.Println("i的地址是:", &i) //变量地址使用&符+变量名 变量在内存的地址
	//var ptr *int = &i
	//ptr 是变量
	//ptr 的类型是 *int
	//ptr 本身的值是 &i
	//指针也有类型 必须同类型才能赋值
	//所有的值类型 都有对应的指针类型 指针用&表示 指针值用*表示
	var ptr = &i //变量ptr要存储变量指针 所以变量ptr本身也有一个指针 如果在找个变量存ptr的指针可无线套娃
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr指针的值=%v\n", *ptr)    //使用*号+变量名 取出指针指向值
	*ptr = 9                            //*ptr直接拿指针ptr·指向的值 进行赋值 相当于给i赋值成了9
	fmt.Printf("修改后ptr指针的值=%v\n", *ptr) //使用*号+变量名 取出指针指向值

	//值类型 int float bool string [] struct 变量直接存储值 内存通常在栈中分配
	//引用类型 指针 slice 切片 map 管道 chan interface

	//TODO 键盘输入语句 需要接收用户输入的数据 就可以使用键盘输入语句
	var names string
	var age byte
	var sal float64
	var isPass bool
	//接收文本 第一种方式
	fmt.Println("请输入姓名:")
	fmt.Scanln(&names)
	fmt.Println("年龄:")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水:")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过考试:")
	fmt.Scanln(&isPass)
	fmt.Printf("名字是:%v \n 年龄是:%v \n 薪水是:%v \n 是否通过考试:%v \n", names, age, sal, isPass)
	//接收文本 第二种方式 可以按指定格式输入
	fmt.Println("请输入你的姓名 年龄 薪水 是否通过考试使用空格隔开")
	fmt.Scanf("%s %d %f %t", &names, &age, &sal, &isPass)
	fmt.Printf("名字是:%v \n 年龄是:%v \n 薪水是:%v \n 是否通过考试:%v \n", names, age, sal, isPass)

	//TODO 原码 反码 补码
	//对于有符号的而言
	//1. 二进制的最高位是符号位 0表示正数 1表示负数
	//2. 负数的反码=他的原码符号位不变，其他位置取反 比如
	//1
	//原码0000 0001 反码是 0000 0001 补码 0000 0001
	//-1
	//原码1000 0001 反码是 1111 1110 补码 1111 1111
	//故此得出结论
	//3. 负数的反码 = 它的符号位以外 全部相反
	//4. 负数的补码 = 他的反码+1
	//5. 0的反码 补码都是0
	//6. 计算机运算时 都是以补码的方式来运算的
	// 1+1  1-1 = 1 + (-1)

	//TODO 位运算
	//与运算符 都是1 取1 否则0
	//分析 2 = 0000 0010
	//     3 = 0000 0011
	//与计算 =  0000 0010
	//结果 0000 0010 = 2
	fmt.Printf(strconv.Itoa(2 & 3))
	//或运算符 有一个1 取1 否则0
	//分析 2 = 0000 0010
	//     3 = 0000 0011
	//或计算 =  0000 0011
	//结果 0000 0011 = 3
	fmt.Printf(strconv.Itoa(2 | 3))
	//异或运算符 一个1一个0取1 相同取0
	//分析 2 =   0000 0010
	//     3 =   0000 0011
	//异或计算 =  0000 0001
	//结果 0000 0001 = 1
	fmt.Printf(strconv.Itoa(2 ^ 3))
	//异或运算符 一个1一个0取1 相同取0
	//分析 -2 =  1000 0011(原)
	//           1111 1100(反)
	//		     1111 1101(补)
	//     3  =  0000 0011
	//异或计算 =  1111 1110(补)
	//           1111 1101(反)
	//           1000 0010(原)
	//结果 1000 0010 = -2
	fmt.Printf(strconv.Itoa(-2 ^ 3))

}
