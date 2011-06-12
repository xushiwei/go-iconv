package main

import (
	"fmt"
	"io"
	"os"
	"bytes"
	"xushiwei.com/iconv"
)

func main() {

	cd1, err1 := iconv.Open("utf-8", "gbk") // gbk => utf8
	if err1 != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd1.Close()

	b1 := bytes.NewBuffer(nil)
	w1 := iconv.NewWriter(cd1, b1, 16, true)
	fmt.Fprintln(w1,
`		你好，世界！你好，世界！你好，世界！你好，世界！
		你好，世界！你好，世界！你好，世界！你好，世界！`)

	cd2, err2 := iconv.Open("gbk", "utf-8") // utf8 => gbk
	if err2 != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd2.Close()
	
	r2 := iconv.NewReader(cd2, b1, 16)
	_, err := io.Copy(os.Stdout, r2)
	if err != nil {
		fmt.Println("\nio.Copy failed:", err)
		return
	}
}

