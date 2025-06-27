/*只出现一次的数字*/
func func_1(nums []int) int {
	var map_1 = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		map_1[nums[i]]++
	}

	for j, val := range map_1 {
		if val == 1 {
			return j
		}
	}
	return -1
}

func func_2(nums []int) int {
	var res = 0
	for _, val := range nums {
		res ^= val
	}
	return res
}

func singleNumber(nums []int) int {
	var res = 0
	//res = func_1(nums);
	res = func_2(nums)
	return res
}

/*回文数*/
func isPalindrome(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}
	res := 0
	for res < x/10 {
		res = res*10 + x%10
		x /= 10
	}

	return res == x || res == x/10
}

/*有效的括号 */
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	res := []rune{}
	for _, ch := range s {
		switch ch {
		case '(':
			res = append(res, ')')
		case '[':
			res = append(res, ']')
		case '{':
			res = append(res, '}')
		default:
			if len(res) == 0 || res[len(res)-1] != ch {
				return false
			}
			res = res[:len(res)-1]
		}
	}
	return len(res) == 0
}

/*最长公共前缀*/
func longestCommonPrefix(strs []string) string {
	res := strs[0]
	for i, ch := range res {
		for _, s := range strs {
			if i == len(s) || s[i] != byte(ch) {
				return res[:i]
			}
		}
	}
	return res
}

/*加一*/
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}

	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}

/*删除有序数组中的重复项*/
func removeDuplicates(nums []int) int {
	//return len(slices.Compact(nums))
	var index = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

/*合并区间*/
func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(left, right []int) int { return left[0] - right[0] })
	var res [][]int
	for _, p := range intervals {
		len_res := len(res)
		if len_res > 0 && p[0] <= res[len_res-1][1] {
			res[len_res-1][1] = max(res[len_res-1][1], p[1]) //更新右端点最大值
		} else {
			res = append(res, p)
		}
	}
	return res
}

/*两数之和 */
func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
