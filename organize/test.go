package organize

import (
	"fmt"
	"reflect"
)

type Book struct {
	Id      int
	Title   string
	Price   float32
	Authors []string
}

func main() {
	book := Book{}
	e := reflect.ValueOf(&book).Elem()

	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		fmt.Printf("%v %v %v\n", varName, varType, varValue)
	}
}
