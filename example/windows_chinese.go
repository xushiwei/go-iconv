package main

import (
	"fmt"
	"os"
	"xushiwei.com/iconv"
)

func main() {
	cd, err := iconv.OpenWith("gbk", "utf-8", os.Stdout, 0)
	if err != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd.Close()
	fmt.Fprintln(cd, "你好，世界！")
	cd.Sync()
}

