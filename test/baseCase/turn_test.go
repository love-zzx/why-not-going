package baseCase

import (
	"fmt"
	"strconv"
	"testing"
)

func TestTurn(t *testing.T) {
	//int 转string
	var a = 0
	b := strconv.Itoa(a)
	fmt.Printf("b = %s,type of b = %T\n", b, b)

	//string 转int
	var c = "test"
	d, _ := strconv.Atoi(c)
	fmt.Printf("d = %d,type of d = %T", d, d)
}

func deferReturn() (ret int) {
	defer func() {
		ret++
	}()
	return 10
}

func TestDefer(t *testing.T) {
	ret := deferReturn()
	fmt.Println("ret = ", ret)
}
