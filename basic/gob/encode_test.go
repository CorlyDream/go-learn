package gob

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

// 自定义类型
type Person struct {
    Name  string
    Age   int
}

func init() {
    // 注册类型
    gob.Register(Person{})
}

func Test_sendData(t *testing.T) {
	buf := new(bytes.Buffer)
    encoder := gob.NewEncoder(buf)
    person := Person{Name: "Alice", Age: 30}
    if err := encoder.Encode(person); err != nil {
        fmt.Println("Encode Error:", err)
    }
    fmt.Println(buf)
}
