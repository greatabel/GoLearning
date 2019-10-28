package main

import (
    "fmt"

    "01Go_In_Action/ch5TypeSystem/i8entities/entities"
)


// main is the entry point for the application.
func main() {
    // Create a value of type User from the entities package.
    u := entities.User{
        Name:  "Bill",
        // 当一个标识符的名字以小写字母开头时，这个标识符就是未公开的，即包外的代码不可见
        Email: "bill@email.com",
    }

    // ./example71.go:16: unknown entities.User field 'email' in
    //                    struct literal

    fmt.Printf("User: %v\n", u)
}