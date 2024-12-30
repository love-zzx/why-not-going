package codding

import (
	"fmt"
	"testing"
)

//移除元素
/*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。
假设 nums 中不等于 val 的元素数量为 left，要通过此题，您需要执行以下操作：
更改 nums 数组，使 nums 的前 left 个元素包含不等于 val 的元素。nums 的其余元素和 nums 的大小并不重要。
返回 left。

示例 1：
输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2,_,_]
解释：你的函数函数应该返回 left = 2, 并且 nums 中的前两个元素均为 2。
你在返回的 left 个元素之外留下了什么并不重要（因此它们并不计入评测）。

示例 2：
输入：nums = [0,1,2,2,3,0,4,2], val = 2
输出：5, nums = [0,1,4,0,3,_,_,_]
解释：你的函数应该返回 left = 5，并且 nums 中的前五个元素为 0,0,1,3,4。
注意这五个元素可以任意顺序返回。
你在返回的 left 个元素之外留下了什么并不重要（因此它们并不计入评测）。


提示：
0 <= nums.length <= 100
0 <= nums[i] <= 50
0 <= val <= 100
*/

//解析：给定一个数组nums和val,你需要输出不等val值的个数left,以及去掉数组中与val相等的元素,然后数组需要按照left个元素前进行排列,之后的元素置为空
//这种方法比较难理解,其实就是左指针相等的累加,
//然后在遍历的过程中,如果不等于val的则累加左指针,并且把当前右指针的值赋值给左指针。
//想要直接在脑海中算出结果比较困难,需要单独列出来即可理解：核心目的就是不断的交换当前指针给累加的左指针,达到
//达到重新赋值的效果，题目不知道写的啥玩意,不知所云，浪费时间。
//直接说移除掉重复的元素,然后返回未重复的元素,可以任意顺序返回即可

func removeElement(nums []int, val int) int {
	//遍历数组,元素是否相等,如果相等的话,放入到新的切片中,否则元素置为空
	//然后重写排序下数组,返回新数组、和left
	left := 0
	for _, v := range nums {
		if v != val { //如果当前值,不等于3,累加的左指针等于当前值
			nums[left] = v //当前右指针等于左累计指针的值
			left++
		}
	}
	return left
	//sort.Slice(nums, func(i, j int) bool {
	//	return i >= j
	//})

}

func TestRemoveElement(t *testing.T) {
	// k = 1
	nums := []int{3, 2, 2, 3}
	//nums := []int{3, 2, 2, 3} 判断是否等于3 => 等于   k
	//nums := []int{3, 2, 2, 3} 判断是否等于3 => 不等于 k0 = 2
	//nums := []int{3, 2, 2, 3} 判断是否等于3 => 不等于 k1 = 2
	//nums := []int{3, 2, 2, 3} 判断是否等于3 => 等于 结束

	//nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	//nums := []int{0, 1, 2, 2, 3, 0, 4, 2} 判断是否等于2 => 不等于 k0 = 0 累加1
	//nums := []int{0, 1, 2, 2, 3, 0, 4, 2} 判断是否等于2 => 不等于 k1 = 1 累加2
	//nums := []int{0, 1, 2, 2, 3, 0, 4, 2} 判断是否等于2 => 等于
	//nums := []int{0, 1, 2, 2, 3, 0, 4, 2} 判断是否等于2 => 等于
	//nums := []int{0, 1, 2, 2, 3, 0, 4, 2} 判断是否等于2 => 不等于 k2 = 3 累加3
	//nums := []int{0, 1, 2, 2, 3, 0, 4, 2} 判断是否等于2 => 不等于 k3 = 0 累加4
	//nums := []int{0, 1, 2, 2, 3, 0, 4, 2} 判断是否等于2 => 不等于 k4 = 4 累加5
	//nums := []int{0, 1, 2, 2, 3, 0, 4, 2} 判断是否等于
	//{0, 1, 3, 0, 4, 0, 4, 2}
	val := 3
	left := removeElement(nums, val)
	fmt.Println(left)
	fmt.Println(nums)
	//var mao := make(map[string]string)
	//var aa map[string]string
}

func removeElement2(nums []int, val int) int {
	//题目是：移除掉等于val的元素,然后剩余的元素按照前len(s)-k 元素排列，顺序不用保证，只要排到前面即可
	//思路：使用双指针：左指针是val元素，右指针是从尾部倒数；
	//左右指针长度判断遍历,如果val等于元素值,则左指针被右指针覆盖,右指针累减
	//否则,左指针累加。最终左指针不等右指针结束
	left, right := 0, len(nums)
	for left < right {
		if val == nums[left] { //左指针元素交换右指针元素
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}

func TestRemoveElement2(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	//nums := []int{3, 2, 2, 3} left = 0,right = 4,交换3和3 =>[]int{3, 2, 2, 3} => left = 0,right = 3
	//nums := []int{3, 2, 2, 3} left = 0,right = 3,交换3和2 =>[]int{2, 2, 2, 3} => left = 0,right = 2
	//nums := []int{2, 2, 2, 3} left = 0,right = 2,不相等 => left = 1
	//nums := []int{2, 2, 2, 3} left = 0,right = 2,不相等 => left = 2
	// left = 2,right = 2 退出循环{2, 2, 2, 3}
	val := 3
	left := removeElement2(nums, val)
	fmt.Println(left)
	fmt.Println(nums)
}
