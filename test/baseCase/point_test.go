package baseCase

import (
	"fmt"
	"testing"
)

var global *int
var y int

func TestPoint(t *testing.T) {
	fmt.Println("y = ", y)           //0
	fmt.Println("global = ", global) //nil

	var x int
	x = 1
	global = &x //global地址指针x => 1
	g()

	fmt.Println("x = ", x)           //1
	fmt.Println("y = ", y)           //0
	fmt.Println("global = ", global) // 0x123123
}

func TestType(t *testing.T) {
	type MyInt = int
	var aa MyInt
	type KKint = MyInt
	var bb KKint = 100
	println(aa)
	println(bb)
}

func g() {
	var y = new(int)
	*y = 1
}
