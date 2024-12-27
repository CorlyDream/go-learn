package slice

import "fmt"

func f(s []int) {
	// 这样不能改变s中元素的值
	//for _, v := range s {
	//	v=10
	//}
	
        // 这样可以
	for i := range s {
		s[i] = 10
	}
}

func modifySlice(s []int) {
    s = append(s, 4) // 触发扩容，底层数组改变
    s[0] = 99       // 修改的是新数组
    fmt.Println("Inside function:", s)
}