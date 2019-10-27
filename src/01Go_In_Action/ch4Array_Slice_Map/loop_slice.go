package main

import (
    "fmt"
)

// main is the entry point for the application.
func main() {

    slice := []int{10, 20, 30, 40}
    for index, value := range slice {
        fmt.Printf("Index: %d Value: %d\n", index, value)
    }

    // 迭代每个元素，并显示值和地址
    for index, value := range slice {
    fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
    }

    //使用空白标识符 （下划线）忽略索引值
    for _, value := range slice {
        fmt.Printf("# Value: %d\n", value)
    }



}
