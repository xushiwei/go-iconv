package main

import (
	"fmt"
	"xushiwei.com/iconv"
)

func main() {
	cd, err := iconv.Open("gbk", "utf-8")
	if err != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd.Close()
	fmt.Fprintln(cd, "你好，世界！")
}

