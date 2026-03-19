package web3_homework

import (
	"fmt"
	"sort"
	"strconv"
)

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	var count int
	for _, n1 := range nums {
		count = 0
		for _, n2 := range nums {
			if n1 == n2 {
				count++
				if count > 1 {
					break
				}
			}
		}
		if count == 1 {
			fmt.Printf(" 数组: %v 只出现一次的元素是 %d ", nums, n1)
			return n1
		}
	}
	return 0
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	str := strconv.Itoa(x)
	l := len(str)
	if l == 1 {
		return true
	}
	for i := 0; i < l/2; i++ {
		if str[i] != str[l-i-1] {
			return false
		}
	}
	return true
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	if len(s)%2 != 0 {
		fmt.Printf(" 字符串: %v 无效", s)
		return false
	}

	checkMap := make([]rune, 0)

	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			checkMap = append(checkMap, ch)
		case ')':
			if len(checkMap) != 0 && checkMap[len(checkMap)-1] == '(' {
				checkMap = checkMap[:len(checkMap)-1]
			}
		case '}':
			if len(checkMap) != 0 && checkMap[len(checkMap)-1] == '{' {
				checkMap = checkMap[:len(checkMap)-1]
			}
		case ']':
			if len(checkMap) != 0 && checkMap[len(checkMap)-1] == '[' {
				checkMap = checkMap[:len(checkMap)-1]
			}
		}
	}
	if len(checkMap) == 0 {
		fmt.Printf(" 字符串: %v 有效\n", s)
		return true
	}
	fmt.Printf(" 字符串: %v 无效\n", s)
	return false
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		fmt.Printf(" 无公共前缀是 \n")
		return ""
	}
	r := ""
	defer func() {
		if r == "" {
			fmt.Printf("无公共前缀\n")
		} else {
			fmt.Printf("公共前缀是 %v \n", r)
		}
	}()
	// 以第一个字符串为基准
	for i := 0; i < len(strs[0]); i++ {
		// 获取第一个字符串的第i个字符
		char := strs[0][i]

		// 遍历其他字符串
		for j := 1; j < len(strs); j++ {
			// 如果当前字符串长度不够，或者字符不匹配
			if i >= len(strs[j]) || strs[j][i] != char {
				r = strs[0][:i]
				return r
			}
		}
	}
	r = strs[0]
	return r
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	original := digits
	result := make([]int, len(digits))
	copy(result, digits)

	defer func() {
		fmt.Printf("输入: %v, 输出: %v\n", original, result)
	}()

	for i := len(result) - 1; i >= 0; i-- {
		if result[i] < 9 {
			result[i]++
			return result
		}
		result[i] = 0
	}
	result = append([]int{1}, result...)
	return result
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0 // 慢指针，指向当前不重复序列的最后一个位置
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	// 数组为空返回
	if len(intervals) == 0 {
		return [][]int{}
	}
	// 按照区间起点升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 结果集
	res := make([][]int, 0)
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		cur := intervals[i]

		if cur[0] <= last[1] {
			// [3,5] [4,6]
			if cur[1] > last[1] {
				last[1] = cur[1]
			}
		} else {
			res = append(res, cur) // 不重叠，直接添加
		}
	}
	return res
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	// key值，val索引
	m := make(map[int]int)

	for i, num := range nums {
		tarNum := target - num
		if j, ok := m[tarNum]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}
