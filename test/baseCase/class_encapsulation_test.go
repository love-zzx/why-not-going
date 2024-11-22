package baseCase

import (
	"fmt"
	"testing"
)

// 通过结构体来定义类的封装encapsulation
// 类名、属性名、方法名首字母大写表示对外可以访问
type Hero struct {
	Name  string
	Ad    int
	Level int
}

//// 封装Show方法
//func (h *Hero) Show() {
//	fmt.Println("Name = ", h.Name)
//	fmt.Println("Ad = ", h.Ad)
//	fmt.Println("Level = ", h.Level)
//}
//
//func (h *Hero) GetName() string {
//	return h.Name
//}
//
//func (h *Hero) SetName(str string) {
//	h.Name = str
//}

// 封装Show方法
func (h Hero) Show() {
	fmt.Println("Name = ", h.Name)
	fmt.Println("Ad = ", h.Ad)
	fmt.Println("Level = ", h.Level)
}

func (h Hero) GetName() string {
	return h.Name
}

func (h Hero) SetName(str string) {
	h.Name = str
}

// 封装
func TestEncapsulation(t *testing.T) {
	h := Hero{
		Name:  "张三的书",
		Ad:    16,
		Level: 1,
	}
	h.Show()
	fmt.Println(h.GetName())
	h.SetName("李四的笔")
	fmt.Println(h.GetName())
}

//====>> 使用引用的情况：结果1*
//=== RUN   TestEncapsulation
//Name =  张三的书
//Ad =  16
//Level =  1
//张三的书
//李四的笔 //局部变量修改,对外产生了影响 => 引用传递
//--- PASS: TestEncapsulation (0.00s)
//PASS

//====>> 未使用引用的情况：结果2
//=== RUN   TestEncapsulation
//Name =  张三的书
//Ad =  16
//Level =  1
//张三的书
//张三的书 //局部变量修改,对外无影响=> 值拷贝
//--- PASS: TestEncapsulation (0.00s)
//PASS
