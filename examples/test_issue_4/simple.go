package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_sensitive"
)

func main() {
	filter := zdpgo_sensitive.New()
	filter.LoadWordDict("dict/dict.txt")
	fmt.Println(filter.Replace("xC4x", '*'))
}
