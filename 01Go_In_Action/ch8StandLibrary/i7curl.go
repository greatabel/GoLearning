package main

import (
    "io"
    "log"
    "net/http"
    "os"
)

func main() {

    r, err := http.Get(os.Args[1])
    if err != nil {
        log.Fatalln(err)
    }
    //创建文件保存相应内容
    file, err := os.Create(os.Args[2])
    if err != nil {
        log.Fatalln(err)
    }
    defer file.Close()

    //使用MultiWriter可以同时向文件和标准输出设备写
    dest := io.MultiWriter(os.Stdout, file)

    // Read the response and write to both destinations.
    io.Copy(dest, r.Body)
    if err := r.Body.Close(); err != nil {
        log.Println(err)
    }
}