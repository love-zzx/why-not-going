package baseCase

import (
	"fmt"
	"testing"
)

// 定义结构体
type myStr struct {
	name string
	age  int
}

type Person struct {
	name string
	age  int
}

type Student struct {
	//第一种嵌套方式
	//p     Person
	//score float32

	//第二种嵌套方式-匿名嵌套
	Person
	score2 float32
	name   string
}

func TestChangeStruct(t *testing.T) {
	s := Student{
		Person{"名称1", 18},
		9.656,
		"名称2",
	}
	//fmt.Println(s.p.age)
	//fmt.Println(s.age)
	fmt.Println(s)
}

func TestStruct(t *testing.T) {
	aa := myStr{"hh", 18}
	//查询结构体
	fmt.Println(aa.name)
	fmt.Println(aa.age)
	//修改结构体
	aa.name = "去哪里"
	//删除结构体
	aa.name = ""
	fmt.Println(aa)

	var newStr []myStr
	//写入结构体
	newStr = append(newStr, myStr{
		name: "happy",
		age:  10,
	}, myStr{
		name: "happy",
		age:  10,
	})
	//修改结构体
	newStr[0].name = "kk"
	newStr[1].age = 12
	//删除结构体
	newStr[0] = myStr{}
	fmt.Println(newStr)
}
