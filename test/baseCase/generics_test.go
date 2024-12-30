package baseCase

import (
	"fmt"
	"testing"
)

// 泛型
func TestGenerics(t *testing.T) {
	var a = []int{1, 2, 3}
	//insert 0 at index 1 => 1 0 2 3
	//a = append(a[:1], append([]int{0}, a[1:]...)...)
	//a = IntSliceInsert(a, 1, 0)
	////insert -1 at index 2
	//a = IntSliceInsert(a, 2, -1)
	//fmt.Println(a)
	//
	//a2 := []string{"1", "2", "3"}
	////insert "0" at index 1 => 1 0 2 3
	//a2 = StringSliceInsert(a2, 1, "0")
	//a2 = StringSliceInsert(a2, 1, "0")
	////insert "-1" at index 2
	//a2 = StringSliceInsert(a2, 2, "-1")
	//a2 = StringSliceInsert(a2, 2, "-1")
	//fmt.Println(a2)

	a = SliceInsert(a, 1, 0)
	//insert -1 at index 2
	a = SliceInsert(a, 2, -1)
	fmt.Println(a)

	a2 := []string{"1", "2", "3"}
	//insert "0" at index 1 => 1 0 2 3
	a2 = SliceInsert(a2, 1, "0")
	a2 = SliceInsert(a2, 1, "0")
	//insert "-1" at index 2
	a2 = SliceInsert(a2, 2, "-1")
	a2 = SliceInsert(a2, 2, "-1")
	fmt.Println(a2)

	//INSERT {0,0} at 1
	a3 := []Pt{{1, 2}, {3, 4}}
	a3 = SliceInsert(a3, 1, Pt{0, 0})
	///insert {-1,-1} at 2
	a3 = SliceInsert(a3, 2, Pt{-1, -1})
	fmt.Println(a3)
}

// --------------------------观察,每增加一个类型,都需要增加一个类型方法，这样代码重复冗余

func IntSliceInsert(a []int, i, v int) []int {
	return append(a[:i], append([]int{v}, a[i:]...)...)
}

func StringSliceInsert(a []string, i int, v string) []string {
	return append(a[:i], append([]string{v}, a[i:]...)...)
}

func AnySliceInsert(a any /*interface{}*/, i any, v any) []any {
	// 1.使用reflect: 速度变慢,太重了
	// 2.使用了any:不就成了动态语言了,go 本身就是静态类语言的优势就没有了,在编译之前会做一些类型检查
	// 3.代码生成,go generate命令:太重了,解决了慢和不安全问题
	panic("no implemented")
}

// 使用泛型T替换掉数据类型
func SliceInsert[T any](a []T, i int, v T) []T {
	return append(a[:i], append([]T{v}, a[i:]...)...)
}

type Pt struct {
	X, Y int
}

//用一个函数解决函数类型的复用=>泛型,定义[T any],有泛型函数、泛型类型、泛型结构体使用】，主要的目的是泛型是针对那种优化得无法再优化的情况下，如果出现了代码一致，
//但是类型不一样的这种情况，可以针对类型复用，减少代码的重复这种情况才使用，不然会增加代码的复杂度
