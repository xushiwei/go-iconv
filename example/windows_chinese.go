package main

import (
	"fmt"
	"os"
	"xushiwei.com/iconv"
)

func main() {
	cd, err := iconv.Open("gbk", "utf-8")
	if err != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd.Close()
	
	w := iconv.NewWriter(cd, os.Stdout, 0, true)
	
	fmt.Fprintln(w, "你好，世界！")
}

