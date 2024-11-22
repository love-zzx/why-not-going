package baseCase

import (
	"fmt"
	"testing"
)

var NewGlobal string

// 定义变量
func TestVariable(t *testing.T) {
	var a int
	var b int = 10
	c := 100
	NewGlobal = "It's fire war,now!"
	//打印变量
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(NewGlobal)
}

// 定义常量
func TestConst(t *testing.T) {
	const name = 1
	const (
		addOne = iota * 10
		addTwo
		addThree
		addFour
		addFive
	)
	fmt.Println(name)
	fmt.Println(addOne)
	fmt.Println(addTwo)
	fmt.Println(addThree)
	fmt.Println(addFour)
	fmt.Println(addFive)
}

// 数组
func TestNewArray(t *testing.T) {
	//定义数组
	var myArray1 = [2]string{}
	var book = [3]string{"mike", "The World!", "createTime"}
	fmt.Println(myArray1)
	fmt.Println(book)
	//查询数组
	fmt.Println(book[0])
	//修改数组
	book[0] = "micael"
	//删除数组【因为数组本身是连续的结构空间,所以无法直接删除,通过新数组的方式剔除掉数组元素】
	newArray := [2]string{}
	var j = 0
	for i, v := range book {
		if i != 2 || v != "createTime" {
			newArray[i] = v
			j++
		}
	}
	fmt.Println(newArray)
}

// 定义切片
func TestSlice(t *testing.T) {
	//var slice1 = []string{}
	var slice2 = make([]string, 3)
	//修改数组
	slice2[0] = "hh"
	slice2[1] = "ok"
	slice2[2] = "aa"
	fmt.Println(len(slice2)) //3
	fmt.Println(slice2)      //[hh ok aa]
	//删除数组【删除元素ok】
	//s1 := fmt.Sprintf("%s", slice2[0:1]) //左开右闭 => hh
	//s2 := fmt.Sprintf("%s", slice2[2:3]) //左开右闭 => aa
	s1 := slice2[0:1] //左开右闭 => hh
	s2 := slice2[2:3] //左开右闭 => aa
	var newSlice []string
	//newSlice = append(newSlice, s1...)
	//newSlice = append(newSlice, s2...) //写法2
	newSlice = append(newSlice, s1[0], s2[0]) //写法1
	fmt.Println(newSlice)
	//写入数组
	slice2 = append(slice2, "ok", "abc", "efg") //注意元素不够会发送扩容
}

// 定义队列
func TestLine(t *testing.T) {

}
