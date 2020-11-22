package main

import (
	"fmt"
	"strings"
)

/*
判断字符串中字符是否全都不同
问题描述

请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。 给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。 保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。
 */

func main() {
	fmt.Println(StringAllEqual("aaaaaaaaaaaaaaaa"))
	fmt.Println(StringAllEqual("aaaaaaaaaa1aaaaaa"))
	fmt.Println(StringAllEqual("aaaaaaaaaAaaaaaaa"))
	fmt.Println(StringAllEqual(""))
}

func StringAllEqual(str string) bool {
	if len(str) == 0 {
		return true
	}
	return len(str) == strings.Count(str, string(str[0]))
	//ch := str[0]
	//for i := 1; i < len(str); i++ {
	//	if str[i] != ch {
	//		return false
	//	}
	//}
	//return true
}