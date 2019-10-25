package main

import "fmt"

func main() {
    /*
这个程序要输出一个浮点数 3.14，但是在格式化字符串里并没有对应的格式化参数。如果对 这段代码执行go vet，会得到如下消息:
    go vet main.go
main.go:6: no formatting directive in Printf call
    */
    fmt.Printf("There quick brown fox jumped over lazy dogs", 3.14)
}