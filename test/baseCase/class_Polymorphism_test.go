package baseCase

import (
	"fmt"
	"testing"
)

// 多态(Polymorphism) 定义：子类实现了接口方法,通过指针的方式指向不同的实现子类,从而实现了【一种数据结构展示多种数据结果】
//定义一个接口：猫、狗都是动物，能跑，能跳

type NAnimal interface {
	Jump()
	Run()
}

type NCat struct {
	name string
	age  int
}

func (c *NCat) Jump() {
	fmt.Println(c.name, "开始跳")
}

func (c *NCat) Run() {
	fmt.Println(c.name, "开始跑")
}

type NDog struct {
	name string
	age  int
}

func (d *NDog) Jump() {
	fmt.Println(d.name, "开始跳")
}

func (d *NDog) Run() {
	fmt.Println(d.name, "开始跳")
}

// 根据接口类调用子类方法
func callNAnimal(a NAnimal) {
	a.Jump()
	a.Run()
}

func TestPolymorphism(t *testing.T) {
	//基本上和接口interface使用的例子一致,唯一的区别是不用通过定义父类a 来接收子类的实现 而实通过以下:
	//核心点：不必使用a来进行接收,直接传入子类的实例的指针即可: callAnimal(&cat)
	//调用动物-猫
	cat := Cat{
		name: "小猫",
		age:  1,
	}
	//调用动物-狗
	dog := Dog{
		name: "小狗",
		age:  2,
	}
	callAnimal(&cat)
	fmt.Println("-------------------------")
	callAnimal(&dog)
}

//=== RUN   TestPolymorphism
//小猫 开始跳
//小猫 开始跑
//-------------------------
//小狗 开始跳
//小狗 开始跳
//--- PASS: TestPolymorphism (0.00s)
//PASS
