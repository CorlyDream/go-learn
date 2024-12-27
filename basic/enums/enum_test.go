package enums

import (
	"fmt"
	"testing"
)

type Color int

const (
    Red Color = iota
    Green
    Blue
)

func (c Color) String() string {
    return [...]string{"Red", "Green", "Blue"}[c]
}

func Test_enum(t *testing.T) {
    var c Color = Green
    fmt.Println(c) // 输出: Green
}
