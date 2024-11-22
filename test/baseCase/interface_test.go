package baseCase

import (
	"fmt"
	"testing"
)

//定义接口：猫、狗都是动物，能跑，能跳
//通过:接口 = 结构体 【接口接收了结构体的实现】

type Animal interface {
	Jump()
	Run()
}

type Cat struct {
	name string
	age  int
}

func (c *Cat) Jump() {
	fmt.Println(c.name, "开始跳")
}

func (c *Cat) Run() {
	fmt.Println(c.name, "开始跑")
}

type Dog struct {
	name string
	age  int
}

func (d *Dog) Jump() {
	fmt.Println(d.name, "开始跳")
}

func (d *Dog) Run() {
	fmt.Println(d.name, "开始跳")
}

// 根据接口类调用子类方法
func callAnimal(a Animal) {
	a.Jump()
	a.Run()
}

func TestInterface(t *testing.T) {
	var a Animal
	//调用动物-猫
	cat := Cat{
		name: "小猫",
		age:  1,
	}
	a = &cat //关联接口方法与结构体方法【a接收了结构体的实现,相当于 implement】=>cat指针指向动物: cat实现了run和jump方法,相当于Animal接收了cat实现的方法
	//a.Jump()
	//a.Run()
	callAnimal(a)
	fmt.Println("-------------------------")
	//调用动物-狗
	dog := Dog{
		name: "小狗",
		age:  2,
	}
	a = &dog //关联接口方法与结构体方法【a接收了结构体的实现,相当于 implement】=>dog指针指向动物: dog实现了run和jump方法,相当于Animal接收了dog实现的方法
	callAnimal(a)
}
