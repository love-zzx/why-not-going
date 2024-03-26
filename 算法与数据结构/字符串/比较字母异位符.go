package main

import "fmt"

// 这段代码首先判断两个单词的长度是否相同，如果不同则直接返回 false。然后通过遍历两个单词，统计每个字母出现的次数，并将统计结果存储在两个长度为 26 的数组中。最后比较两个数组中对应位置的元素是否相同，从而确定两个单词是否是字母异位词。

// 判断异位词,指的是字母顺序不一样;通过判断,如果相同条件返回true,否则返回false
// 实现: 字母小写26字符,设定2个数组,通过数组的小标对应累积相应下标的值,最后判断下每个数组的下标是否相等,如果不相等返回false
func main() {
	word1 := "hello"
	word2 := "hlleo"
	fmt.Println(isNewAnagram(word1, word2))
}

func isNewAnagram(word1, word2 string) bool {
	//比较长度是否相等
	if len(word1) != len(word2) {
		return false
	}

	//定义两个数组,用于存储数组下标用,通过它来比对值是否一致
	arr1 := [26]int{}
	arr2 := [26]int{}

	//数组打印之前
	fmt.Println(arr1)
	fmt.Println(arr2)

	//统计单词们在数组各自出现的次数
	for _, char := range word1 {
		arr1[char-'a']++
	}

	for _, char := range word2 {
		arr2[char-'a']++
	}

	//数组打印之后
	fmt.Println(arr1)
	fmt.Println(arr2)

	//检查单词出现的次数是否一致
	for i := 0; i < 26; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
