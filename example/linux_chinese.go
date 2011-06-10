package main

import (
	"fmt"
	"os"
	"xushiwei.com/iconv"
)

func main() {

	cd1, err1 := iconv.OpenWith("utf-8", "gbk", os.Stdout, 0, false)
	if err1 != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd1.Close()
	
	cd2, err2 := iconv.OpenWith("gbk", "utf-8", cd1, 0, false)
	if err2 != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd2.Close()
	
	fmt.Fprintln(cd2, "你好，世界！")
}

