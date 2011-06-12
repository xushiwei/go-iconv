package main

import (
	"fmt"
	"os"
	"xushiwei.com/iconv"
)

func main() {

	cd1, err1 := iconv.Open("utf-8", "gbk") // gbk => utf8
	if err1 != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd1.Close()

	w1 := iconv.NewWriter(cd1, os.Stdout, 0, true)
	
	cd2, err2 := iconv.Open("gbk", "utf-8") // utf8 => gbk
	if err2 != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd2.Close()
	
	w2 := iconv.NewWriter(cd2, w1, 0, true)
	
	fmt.Fprintln(w2, "你好，世界！")
}

