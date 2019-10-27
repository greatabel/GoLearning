package main

import (
    "fmt"
)

// main is the entry point for the application.
func main() {

    // var array [4][2]int

    array := [4][2]int{{10,11}, {20,21}, {30, 31}, {40, 41}}
    fmt.Printf("There are %d %d \n",array[0][0], array[0][1])

    // 不需要指定[]的值
    slice0 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
    //创建一个整形切片,长度和容量都是3个元素
    slice1 := []int{10, 20, 30}
    newslice := slice1[1:2]

    //使用空字符串初始化第100个元素
    slice2 := []string{99: ""}
    fmt.Printf("There are %s %d %d %s\n",slice0[0], slice1[0], newslice[0], slice2[0])


}
