package baseCase

import (
	"fmt"
	"sort"
	"testing"
)

/*
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
*/
func TestMergeSortArray(t *testing.T) {
	//方法1: 合并+排序
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 3
	n := 3

	//merge1(nums1, m, nums2, n)
	//merge2(nums1, m, nums2, n)
	merge3(nums1, m, nums2, n)
}

// 合并切片+排序 时间复杂度：O((m+n)*log(m+n))
func merge1(nums1 []int, m int, nums2 []int, n int) {
	//思路：复制合并,自带函数排序
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
	fmt.Println(nums1)
}

// 双指针 时间复杂度: O(n+m) 空间复杂度O(n)
// 题解：把num1和num2本身当前已经排序好的队列，然后从2个队列中取值放入到新sorted切片中
func merge2(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		//如果数组1最后一个元素,则把p2剩余元素合并进行【因为p2本身是按照正序排序】
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		//如果数组2最后一个元素,则把p1剩余元素合并进行【因为p1本身是按照正序排序】
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		//如果小于
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
	fmt.Println(nums1)
}

// 逆向双指针[从后往前] 时间复杂度: O(n+m) 空间复杂度O(1)
func merge3(nums1 []int, m int, nums2 []int, n int) {
	//直接使用p1存储变量,因为它的空间足够大
	//p1和p2下标从最后开始倒数,总容量是：m+n-1倒数,tail 表示最后元素,递减 【遍历nums--】 只需遍历一遍Num1
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int   //当前值
		if p1 == -1 { //数组1是第一个元素,则数组2填充,递减
			cur = nums2[p2]
			p2--
		} else if p2 == -1 { //数组2是第一个元素,则数组1填充,递减
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] { //如果数组1值大于数组2值,取数组1
			cur = nums1[p1]
			p1--
		} else { //如果数组2值大于数组1值,取数组2
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
	fmt.Println(nums1)
}
