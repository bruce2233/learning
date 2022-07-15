// package main
package mem
import (
	"fmt"
	"unsafe"
)

type big struct {
	name string
	age  int
}

func main() {

	var s []big //结构体切片
	var pOld uintptr
	var capOld, capNow int
	var reallocTimes, recapTimes = 0, 0
	for i := 0; i < 10000; i++ {
		//必须在添加元素前，保存前一个元素的地址
		//一旦自动扩容后，前一个元素的地址也会改变,无法检测底层数组是否重写
		if i > 0 {
			pOld = uintptr(unsafe.Pointer(&s[i-1]))
		}
		capOld = cap(s) //添加元素前容量
		sNewEle := big{
			name: "zhangye gogogogogogogogogogogogogo",
			age:  22,
		} //添加的切片新元素
		s = append(s, sNewEle)

		fmt.Print("i:", i, " ")
		fmt.Printf("&s[i]: %p ", &s[i])              //当前元素地址
		fmt.Print("&s[i].name: ", &(s[i].name), " ") //结构体字段的地址
		fmt.Print("cap(s): ", cap(s), " \n")         //当前容量

		if i > 0 {
			pNow := uintptr(unsafe.Pointer(&s[i]))
			if pNow-pOld != unsafe.Sizeof(sNewEle) { //底层数组重写
				fmt.Print("s底层数组重写,发生拷贝", i)
				reallocTimes++
			}
			capNow = cap(s)
			if capNow != capOld { //扩容
				recapTimes++
			}
		}
	}
	fmt.Print("underlying realloc times: ", reallocTimes, "\n") //底层数组重写次数
	fmt.Print("recap times: ", recapTimes, "\n")                //扩容次数
}
