package main

import (
	"fmt"
	"strings"
	"xingzuo/shengxiao"
)

func main() {
	istr := "吉\t凶\t吉\t吉\t凶\t凶\t吉\t吉\t凶\t吉\t凶\t凶"
	ar := strings.Split(istr, "\t")
	for i := 0;i<len(shengxiao.ShengxiaoAr);i++{
		fmt.Print(shengxiao.ShengxiaoAr[i])
		fmt.Print(" ")
	}
	fmt.Println()
	for i := 0; i < len(ar); i++ {
		fmt.Print(ar[i]+" ")
	}

}
