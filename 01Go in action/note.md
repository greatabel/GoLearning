@
Go平衡了底层系统语言的能力，以及现代语言中所见到的高级特性。

C/C++提供了很快的执行速度、Ruby/Python擅长快速开发。Go架起2者桥梁，提供高性能，同时让开发更快捷

Go简化了解决依赖的算，提供了更快的编译速度。编译Go，编译器只关注直接被引用的库。

1.goroutine很像线程，但占用内存远少于线程。
通道channel是内置的数据结构，可以让用户在不同的goroutine之间发送具有类型的消息
2. 通道channel，帮用户避免其他语言常见的共享内存访问问题。


@ 类型系统
Go提供无继承的类型系统，使用组合composition设计模式
Go 内置值类型：bool, int, unit, float, string, complex, array
       引用类型：slice（序列数组）、map（映射）、chan（管道）








































































