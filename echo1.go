package main

import (
	"fmt"
	"os"
)

func main() {
	Echo1()
	echo2()
}

// Echo1 echo
func Echo1() {
	var s, blank string // 使用显示声明 来说明初始化变量的重要性。相反下面的方法中使用隐式声明来表明初始化变量不重要
	for i := 1; i < len(os.Args); i++ {
		s += blank + os.Args[i]
		blank = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, blank := "", "" // 推荐使用短变量的声明方式。
	for _, arg := range os.Args[1:] {
		s += blank + arg
		blank = " "
	}
	fmt.Println(s)
}
