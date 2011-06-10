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

## Output to stdout

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

    fmt.Fprintln(cd, "你好，世界！")
}
```

## Output to other output devices

```go
import (
    "fmt"
    "xushiwei.com/iconv"
)

func main() {

    output := ... // eg. output, err := os.Create(file)

    autoSync := false // buffered or not
    cd, err := iconv.OpenWith("gbk", "utf-8", output, 0, autoSync)
    if err != nil {
        fmt.Println("iconv.Open failed!")
        return
    }
    defer cd.Close()

    fmt.Fprintln(cd, "你好，世界！")

    cd.Sync() // if autoSync = false, you need call Sync() by yourself
}
```

