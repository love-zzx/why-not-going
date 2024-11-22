package baseCase

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	var a int = 100
	var b int = 200

	fmt.Println("交换前a的值为 : ", a)
	fmt.Println("交换前b的值为 : ", b)

	swap2(a, b)

	fmt.Println("交换后a的值为 : ", a)
	fmt.Println("交换后b的值为 : ", b)
}

func TestQuotation(t *testing.T) {
	var a int = 100
	var b int = 200

	fmt.Println("交换前a的值为 : ", a)
	fmt.Println("交换前b的值为 : ", b)

	swap3(&a, &b)

	fmt.Println("交换后a的值为 : ", a)
	fmt.Println("交换后b的值为 : ", b)
}

func swap(x, y int) int {
	var temp int
	temp = x
	x = y
	y = temp
	return temp
}

// 作用域 值传递 引用传递 Scope, value passing, reference passing
func swap2(x, y int) int {
	x, y = y, x
	return x
}

func swap3(x, y *int) int {
	*x, *y = *y, *x
	return *x
}
