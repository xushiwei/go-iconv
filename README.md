go-iconv: libiconv for go

# Summary

go-iconv is a libiconv wrapper for go. libiconv Convert string to requested character encoding.

go-iconv project's homepage is: https://github.com/xushiwei/go-iconv


# Install

```
git clone git@github.com:xushiwei/go-iconv.git
cd go-iconv
make install
```

# Example

## Convert string

```go
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

	gbk := cd.ConvString("你好，世界！")

	fmt.Println(gbk)
}
```

## Output to io.Writer

```go
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

	output := ... // eg. output := os.Stdout || ouput, err := os.Create(file)
	autoSync := false // buffered or not
	bufSize := 0 // default if zero
	w := iconv.NewWriter(cd, output, bufSize, autoSync)

	fmt.Fprintln(w, "你好，世界！")

	w.Sync() // if autoSync = false, you need call Sync() by yourself
}
```

## Input from io.Reader

```go
import (
	"fmt"
	"io"
	"os"
	"xushiwei.com/iconv"
)

func main() {

	cd, err := iconv.Open("utf-8", "gbk") // gbk => utf8
	if err != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd.Close()
	
	input := ... // eg. input := os.Stdin || input, err := os.Open(file)
	r := iconv.NewReader(cd, , 0)
	
	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		fmt.Println("\nio.Copy failed:", err)
		return
	}
}
```

