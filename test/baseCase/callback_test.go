package baseCase

import (
	"fmt"
	"testing"
)

var local int = 0

func autoAdd() int {
	local += 1
	return local
}

// 全局添加
func TestGlobalAutoAdd(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(autoAdd())
	}
	//=== RUN   TestGlobalAutoAdd
	//1
	//2
	//3
	//4
	//5
	//--- PASS: TestGlobalAutoAdd (0.00s)
	//PASS
}

func autoAdd1() int {
	local := 0
	return func() int {
		local += 1
		return local
	}()
}

// 局部添加
func TestLocalAutoAdd(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(autoAdd1())
	}
	//=== RUN   TestLocalAutoAdd
	//1
	//1
	//1
	//1
	//1
	//--- PASS: TestLocalAutoAdd (0.00s)
	//PASS
}

func autoAdd3() func() int {
	local := 0
	return func() int {
		local += 1
		return local
	}
}

// 闭包添加
func TestClosureAutoAdd(t *testing.T) {
	//初始化autoAdd3
	localAdd := autoAdd3()
	for i := 0; i < 5; i++ {
		fmt.Println(localAdd())
	}
}
