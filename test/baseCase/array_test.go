package baseCase

import (
	"fmt"
	"sync"
	"testing"
)

func TestArray(t *testing.T) {
	var myArray1 [10]int
	myArray2 := [10]int{1, 2, 3, 4, 5}
	myArray3 := [4]int{1, 2, 3, 4}
	fmt.Println(myArray1)
	fmt.Println(myArray2)
	fmt.Println(myArray3)

	fmt.Println("------------------------")
	printArray(myArray3)
	fmt.Println(myArray3)
}

func printArray(myArray [4]int) {
	fmt.Println(myArray)
	myArray[1] = 100
	fmt.Println(myArray)
}

func TestDic(t *testing.T) {
	// 创建一个长度为3的固定长度数组并传入函数
	arrayToPass := []string{"apple", "banana", "cherry"}
	driArray(arrayToPass, 1, 2, 3)
}

func driArray(myArray []string, myInt ...int) {
	for i, s := range myArray {
		fmt.Println("i = ", i, " s = ", s)
	}
	fmt.Println(myInt)
}

func TestArrayChange(t *testing.T) {
	var arr [3]int
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	arr[2] = 100
	fmt.Println(arr)
}

func TestMultiArray(t *testing.T) {
	var info [3][4]string
	fmt.Println(info)
	fmt.Println("---------------------")
	info[0] = [4]string{"a", "b", "c", "d"}
	info[1] = [4]string{"e", "f", "g", "h"}
	info[2] = [4]string{"i", "j", "k", "l"}
	fmt.Println(info)
	fmt.Println("---------------------")
	info[0][1] = "hello"
	info[0][2] = "world!"
	fmt.Println(info)
}

var mapMutex sync.Mutex

func syncMap(t testing.T) {
	//1. 线程安全的map，无非就是多进程的情况下，读取是安全的，这里肯定是需要用到锁的，可以结合syny.mutext进行处理,然后锁也分为读锁和写锁，读锁性能稍微好点，需要分开来处理即可
	//2. 分析map： 创建、增加、删除、查看 这3个方法中，通过加锁的方式实现线程安全。

	//创建map【写锁】
	//添加【写锁】
	//删除【写锁】
	//查看【读锁】
}

//func createMap(key string, MyMap map[any]any) map[any]any {
//	mapMutex.Lock()
//	//MyMap
//	mapMutex.Unlock()
//
//}
