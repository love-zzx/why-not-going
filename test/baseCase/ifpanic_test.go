package baseCase

import (
	"fmt"
	"reflect"
	"testing"
)

func funcName(a interface{}) {
	fmt.Println("----------------------------------------------------------------")
	fmt.Println(a)

	//value, ok := a.(string) //类型断言
	// 通过反射获取值的类型信息
	value := reflect.ValueOf(a)

	kind := value.Kind()
	if kind == reflect.Struct {
		fmt.Println("The value is a struct.")
	} else if kind == reflect.Int {
		fmt.Println("The value is a Int.")
	} else if kind == reflect.String {
		fmt.Println("The value is a String.")
	} else if kind == reflect.Float32 {
		fmt.Println("The value is a struct.")
	} else {
		fmt.Println("understand kind = ", kind)
	}
	fmt.Println("The value is ", value)
}

type Book struct {
	title string
}

func TestAa(t *testing.T) {
	book := Book{"Golang"}
	funcName(book)
	funcName(10)
	funcName("abc")
	funcName(3.14)
}
