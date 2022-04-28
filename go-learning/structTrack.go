package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	name   string
	age    int
	gender bool
	height float64
	weight float32
}

func main() {
	p := Person{
		name:   "z啧啧啧做做做做做做",
		age:    22,
		gender: true,
		height: 1.80,
	}
	fmt.Println(unsafe.Sizeof(p))
	fmt.Println(unsafe.Sizeof(p.name))
	fmt.Println(unsafe.Sizeof(p.age))
	fmt.Println(unsafe.Sizeof(p.gender))
	fmt.Println(unsafe.Sizeof(p.height))
	fmt.Println(unsafe.Sizeof(p.weight))
	fmt.Println(unsafe.Sizeof(&p.gender))
	fmt.Println(unsafe.Sizeof(&p.age))
	var t uintptr
	fmt.Println(unsafe.Sizeof(t))
	var a Person
	fmt.Println(a)
	var b string
	fmt.Println(b)
	myrune := []rune(p.name)
	fmt.Println("rune")
	myrune[0] = '张'
	fmt.Printf("%c\n", myrune[0])
	myrune = append(myrune, '野')
	fmt.Println(len(myrune))
	fmt.Println(unsafe.Sizeof(myrune))
	fmt.Println("p")
	fmt.Println(&p.name)
	fmt.Println(&p.age)
	fmt.Println(&p.gender) //内存对齐
	fmt.Println(&p.height)
	fmt.Println(&p.weight)
	fmt.Println(unsafe.Sizeof("a"))

}
