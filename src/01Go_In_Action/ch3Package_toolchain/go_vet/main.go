package main

import "fmt"

func main() {
	/*
		这个程序要输出一个浮点数 3.14，但是在格式化字符串里并没有对应的格式化参数。如果对 这段代码执行go vet，会得到如下消息:
		    go vet main.go
		main.go:6: no formatting directive in Printf call

	    fmt 是 Go 语言社区很喜欢的一个命令。fmt 工具会将开发人员的代码布局成和 Go 源代码 类似的风格，
	    不用再为了大括号是不是要放到行尾，
	    或者用 tab(制表符)还是空格来做缩进而 争论不休。使用 go fmt 后面跟文件名或者包名，就可以调用这个代码格式化工具
	*/
	fmt.Printf("There quick brown fox jumped over lazy dogs", 3.14)
}
