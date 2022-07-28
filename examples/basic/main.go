package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_sensitive"
)

func main() {
	filter := zdpgo_sensitive.New()

	// 加载本地字典
	filter.LoadWordDict("dict/dict.txt")

	// 添加敏感词
	filter.AddWord("垃圾")

	// 把词语中的字符替换成指定的字符，这里的字符指的是rune字符，比如*就是'*'。
	result := filter.Replace("这篇文章真的好垃圾", '*')
	fmt.Println(result)

	// 直接移除词语
	result = filter.Filter("这篇文章真的好垃圾啊")
	fmt.Println(result)

	// 查找并返回第一个敏感词，如果没有则返回false
	isOk, result := filter.FindIn("这篇文章真的好垃圾")
	fmt.Println(isOk, result)
	// output => true, 垃圾

	// 验证内容是否ok，如果含有敏感词，则返回false和第一个敏感词。
	isOk, result = filter.Validate("这篇文章真的好垃圾")
	fmt.Println(isOk, result)
	// output => false, 垃圾

	// 查找内容中的全部敏感词，以数组返回。
	results := filter.FindAll("这篇文章真的好垃圾")
	fmt.Println(results)
	// output => [垃圾]

	// 加载网络词库。
	filter.LoadNetWordDict("https://raw.githubusercontent.com/importcjj/sensitive/master/dict/dict.txt")

	// 设置噪音模式，排除噪音字符。
	isOk, result = filter.FindIn("这篇文章真的好垃x圾")
	fmt.Println(isOk, result)

	// 去除噪音
	filter.UpdateNoisePattern(`x`)
	isOk, result = filter.FindIn("这篇文章真的好垃x圾") // true, 垃圾
	fmt.Println(isOk, result)

	isOk, result = filter.Validate("这篇文章真的好垃x圾") // False, 垃圾
	fmt.Println(isOk, result)
}
