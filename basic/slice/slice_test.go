package slice

import (
	"fmt"
	"testing"
)

func Test_f(t *testing.T) {
	s := []int{0, 0}
	f(s)
	fmt.Println(s) // [10,10,10]
}

func Test_modifySlice(t *testing.T) {
    slice := []int{1, 2, 3}    // 长度 3，容量 3
    fmt.Println("Before:", slice)

    modifySlice(slice)          // 扩容后，函数内和函数外的底层数组不一致
    fmt.Println("After:", slice)
}