/*
	.		匹配除换行符以外的任意字符
	\w		匹配字母或数字或下划线或汉字
	\s		匹配任意的空白符
	\d		匹配数字
	\b		匹配单词的开始或结束
	^		匹配字符串的开始
	$		匹配字符串的结束
	*		重复零次或多次
	+		重复一次或更多次
	{n}		重复n次
	{n,}	重复n次或更多次
	{n,m}	重复n到m次
*/

package main

import "regexp"

import "fmt"

func main() {
	reg := regexp.MustCompile("a\\w+") // 正则表达式中的“\”需要转义

	result := reg.FindAllString("zhangsanl", -1)
	fmt.Println(result) // [angsanl]

	reg1 := regexp.MustCompile("(a\\w{1})\\w+(a\\w{1})")
	result1 := reg1.FindAllStringSubmatch("zhangsanl", -1)
	fmt.Println(result1) // [[angsan an an]]
}
