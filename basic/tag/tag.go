package tag

import (
	"fmt"
	"reflect"
)

type Person struct {
    Name    string `json:"name" validate:"required"`
    Age     int    `json:"age"`
}

func TagParse() {
	person := Person{Name: "John", Age: 30}
	t := reflect.TypeOf(person)
	
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		validateTag := field.Tag.Get("validate")
		fmt.Printf("Field: %s, JSON Tag: %s, Validate Tag: %s\n", field.Name, jsonTag, validateTag)
	}
	
}