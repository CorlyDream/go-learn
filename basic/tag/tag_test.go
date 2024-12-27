package tag

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_TagParse(t *testing.T) {
	TagParse()
}

func Test_pointerEqual(t *testing.T) {
	a := []string{"apple", "banana", "cherry"}
	b := a // b 是 a 的引用
    fmt.Printf("Address of a: %p\n", &a)
    fmt.Printf("Address of b: %p\n", &b) 
	fmt.Printf("Address of a's underlying array: %p\n", a)
    fmt.Printf("Address of b's underlying array: %p\n", b)
	fmt.Printf("reflect pointer: %v\n", reflect.ValueOf(a).Pointer())
	fmt.Println(reflect.ValueOf(a).Pointer())
}
