package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeFile(filename string) {
    file, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close() // 将 "file.Close()" 压入 defer 栈中

    writer := bufio.NewWriter(file)
    defer writer.Flush() // 将 "writer.Flush()" 压入 defer 栈中

    fmt.Fprintln(writer, "123")
    // 当函数执行结束时，从 defer 栈中执行语句 - 后进先出，先 "writer.Flush()"，再 "file.Close()"
}

func main() {
    writeFile("defer.txt")
}
