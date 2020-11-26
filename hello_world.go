package main

import fm "fmt"

func main() {
	var i = Image{"http:xxx", 1, 2, 3}
	i.size = 1024
	_, arg2 := Method2(1, 1)
	fm.Println("arg2 =", arg2)
	fm.Println("Hello World")
	fm.Println(sum(1, 2))
	fm.Println(reduce(1, 2))
	fm.Println(i.Method1(3, 2))
	fm.Println(i.size)
	fm.Println(i.height)
	fm.Println(i.width)
	fm.Println(i.url)

	var a = 1.91
	fm.Println(a)
	var b = int(a)

	fm.Println(b)
}

func sum(a, b int) int {
	return a + b
}

func reduce(a, b int) int {
	return a - b
}

const c = "C"

var v int = 5

// Image type
type Image struct {
	url    string
	size   int
	width  int
	height int
}

// Method1 xx
func (t Image) Method1(a, b int) int {
	return a % b
}

// Method2 xx
func Method2(a, b int) (int, int) {
	a++
	b++
	return a, b
}

const (
	_ = iota
	// KB 1左移10位
	KB = 1 < (10 * iota)
	// MB 1左移20位
	MB
	// GB 1左移30位
	GB
	// TB 1左移40位
	TB
	// PB 1左移50位
	PB
)
