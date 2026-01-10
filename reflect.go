package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Sample struct {
// 	Name string `required:"true" max:"10"`
// }

// type Person struct {
// 	Name string `required:"true" max:"10"`
// 	Age int `required:"true"`
// }

// func readField(value any) {
// 	var valueType reflect.Type = reflect.TypeOf(value)

// 	fmt.Println("Type Name", valueType.Name())

// 	for i := 0; i < valueType.NumField(); i++ {
// 		var valueField reflect.StructField = valueType.Field(i)
// 		fmt.Println(valueField.Name, "with type", valueField.Type)

// 		// struct tag
// 		fmt.Println("struct tag required:", valueField.Tag.Get("required"))
// 		fmt.Println("struct tag max:", valueField.Tag.Get("max"))
// 	}
// }

// func isValid(value any) bool {
// 	t := reflect.TypeOf(value)

// 	for i := 0; i < t.NumField(); i++ {
// 		f := t.Field(i)
// 		if f.Tag.Get("required") == "true" {
// 			data := reflect.ValueOf(value).Field(i).Interface()
// 			valid := data != ""

// 			if !valid {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }

// func main() {
// 	// readField(Sample{"Yudha"})
// 	// readField(Person{"Dede", 24})

// 	person := Person{
// 		Name: "Yudha",
// 		Age: 8,
// 	}

// 	fmt.Println(isValid(person))
// }