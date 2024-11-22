package baseCase

import (
	"fmt"
	"testing"
)

// 继承(inheritance)
// 父类
type FatherHuman struct {
	name string
	sex  bool
}

func (f *FatherHuman) Eat() {
	fmt.Println("FatherHuman Eat()...")
}

func (f *FatherHuman) Walk() {
	fmt.Println("FatherHuman Walk()...")
}

// 子类【重写父类的: Eat和Walk】
type ChildHuman struct {
	FatherHuman //继承了父类的所有方法和属性
	level       int
}

func (c *ChildHuman) Eat() {
	fmt.Println("Children Eat()...")
}

func (c *ChildHuman) Walk() {
	fmt.Println("Children Walk()...")
}

func (c *ChildHuman) Smile() {
	fmt.Println("Children is laughing")
}

func TestInheritance(t *testing.T) {
	f := FatherHuman{
		name: "超人daddy",
		sex:  false,
	}
	fmt.Printf("name = %s,sex = %v\n", f.name, f.sex)
	f.Eat()
	f.Walk()
	fmt.Println("-------------------------------------继承")
	c := ChildHuman{
		FatherHuman: f, //这里也可以直接重置
		//FatherHuman: FatherHuman{},
		level: 0,
	}
	c.name = "超人宝宝"
	c.sex = true
	fmt.Printf("name = %s,sex = %v\n", c.name, c.sex)
	c.Eat()
	c.Walk()
	c.Smile()
}

//通过继承,重写了方法Eat()和Walk(),增加了Smale()
//=== RUN   TestInheritance
//name = 超人daddy,sex = false
//FatherHuman Eat()...
//FatherHuman Walk()...
//-------------------------------------继承
//name = 超人宝宝,sex = true
//Children Eat()...
//Children Walk()...
//Children is laughing
//--- PASS: TestInheritance (0.00s)
//PASS
