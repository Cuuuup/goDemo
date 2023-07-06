package main

import "fmt"
import "unsafe"

// 全局变量声明
var (
	n1 = 1
	n2 = "2"
)

func main() {
	//输出文本
	fmt.Println("Hello, World!")
	//声明空变量
	var name, email string
	fmt.Println(name, email)
	//赋值声明变量
	no, phone := 001, 10086
	fmt.Println(no, phone)
	//输出全局变量
	fmt.Println(n1, n2)

	//整数类型 int
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

	//浮点型 float64
	var price32 float32 //4字节 范围不一样 精度缺失
	var price64 float64 //8字节 64精度比32的准确
	price32, price64 = 1.25, 1.43
	fmt.Println(price32, price64)

	//字符类型 char go语言统一使用utf-8不会有乱码问题
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

	//布尔类型 bool 只有 true false 其他替代不允许使用 占用字节1
	var _t = true
	_t = false
	var _f = false
	_f = true
	fmt.Println(unsafe.Sizeof(_t), unsafe.Sizeof(_f))

	//字符串类型 string  go语言的字符串的字节使用uft8编码表示Unicode文本
	//TODO go中的字符串是不可以改变的 比如拿address中的 address[0] = '法' 是不被允许的
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

}
