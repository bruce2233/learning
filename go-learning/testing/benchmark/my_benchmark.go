package benchmark

import (
	"fmt"
	"reflect"

	// "reflect"
	"unsafe"
)

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return fib(n-2) + fib(n-1)
	}
}

func a() {
	// b := reflect.StringHeader{
	// }
	oriStr := "qwer"
	newStr := oriStr + "a"
	fmt.Printf("oriStr p: %p\n", &oriStr)
	fmt.Printf("unsafe(&oriStr) p: %p\n", unsafe.Pointer(&oriStr))
	b := (*reflect.StringHeader)(unsafe.Pointer(&oriStr))
	// fmt.Printf("b p: %p\n", b)
	fmt.Printf("oriStr v: %v\n", *b)
	fmt.Printf("newStr p: %p\n", &newStr)
	b = (*reflect.StringHeader)(unsafe.Pointer(&newStr))
	// fmt.Printf("b p: %p\n", b)
	fmt.Printf("newStr v: %v\n", *b)
}
