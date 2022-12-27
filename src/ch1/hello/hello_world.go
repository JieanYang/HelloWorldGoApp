package main // 包，表明代码所在的模块（包）

import (
	"fmt" // 引用代码依赖
	"os"
)

// 功能实现
func main() {
	// fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Hello World", os.Args[1]) // 打印从命令行进入的参数
	}
	// fmt.Println("Hello World")
	os.Exit(0)
}

// **应用程序入口**
// 1. 必须是 main 包： package main
// 2. 必须是 main 方法：func main()
// 3. 文件名不一定是 main.go
