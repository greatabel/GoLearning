package main

import (
    "fmt"

    "01Go_In_Action/ch5TypeSystem/i9entities_combine/entities"
)


// main is the entry point for the application.
func main() {
    // Create a value of type User from the entities package.
    a := entities.Admin{
        Rights: 10,
    }
    //设置未公开的内部类型的公共字段
    a.Name = "Bill"
    a.Email = "bill@test.com"

    // ./example71.go:16: unknown entities.User field 'email' in
    //                    struct literal

    fmt.Printf("User: %v\n", a)
    fmt.Println("即便内部类型是未公开的，内部类型里声明的字段依旧是公开的。既然 内部类型的标识符提升到了外部类型，" +
        "这些公开的字段也可以通过外部类型的字段的值来访问")
}