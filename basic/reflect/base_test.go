package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_typeof(t *testing.T) {
    var x float64 = 3.4
    fmt.Println("type:", reflect.TypeOf(x))
}

func Test_valueof(t *testing.T) {
    var x float64 = 3.4
    v := reflect.ValueOf(x)
    fmt.Println("value:", v)

	// modify value
	pv := reflect.ValueOf(&x)
	pv.Elem().SetFloat(7.1)
	fmt.Println("value:", x)
}
