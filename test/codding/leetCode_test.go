package codding

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
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	//merge1(nums1, m, nums2, n)
	//merge2(nums1, m, nums2, n)
	merge3(nums1, m, nums2, n)
}

// 合并切片+排序 时间复杂度：O((m+n)*log(m+n))
func merge1(nums1 []int, m int, nums2 []int, n int) {
	//先合并再排序
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
	fmt.Println(nums1)
}

// 双指针 时间复杂度: O(n+m) 空间复杂度O(n)
// 题解：把num1和num2本身当前已经排序好的队列，然后从2个队列中取值放入到新sorted切片中
func merge2(nums1 []int, m int, nums2 []int, n int) {
	//定义切片sorted作为中间件,然后合并到nums1
	var sorted []int
	p1, p2 := 0, 0
	for {
		//如果数组1最后元素,把数组2元素都插入
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		} else if p2 == n { //如果数组2最后元素,把数组1元素都插入
			sorted = append(sorted, nums2[p1:]...)
			break
		}
		if nums1[p1] > nums2[p2] { //如果数组2小于数组1,选择数组2元素,否则相反
			sorted = append(sorted, nums2[p2])
			p2++
		} else {
			sorted = append(sorted, nums1[p1])
			p1++
		}
	}
	copy(nums1, sorted)
	fmt.Println(nums1)
}

// 逆向双指针[从后往前] 时间复杂度: O(n+m) 空间复杂度O(1)
func merge3(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		cur := 0
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
	fmt.Println(nums1)
}
