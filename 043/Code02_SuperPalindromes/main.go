package main

import "strconv"

// 906. 超级回文数
// https://leetcode.cn/problems/super-palindromes/
// 1-18位数字 O(n)必然超时，所以需要找种子
// 使用根号后的数字作为种子，数量级变为n^9, 仍然需要优化
// 使用前一半作为生成，数量级变为n^5
func isPalindromes(num int) bool {
	offset := 1
	for num/offset >= 10 {
		offset *= 10
	}
	for num != 0 {
		if num/offset != num%10 {
			return false
		}
		num = (num % offset) / 10
		offset /= 100
	}
	return true
}
func makePalindromes(num int) (int, int) {
	num1, num2 := num, num
	num1 *= 10
	num1 += num % 10
	num /= 10
	for ; num != 0; num /= 10 {
		num1 *= 10
		num2 *= 10
		num1 += num % 10
		num2 += num % 10
	}
	return num1 * num1, num2 * num2
}

func superpalindromesInRange(left string, right string) int {
	lNum, _ := strconv.Atoi(left)
	rNum, _ := strconv.Atoi(right)
	res := 0
	for i := 0; ; i++ {
		palNum1, palNum2 := makePalindromes(i)
		if palNum2 > rNum {
			return res
		}
		if palNum1 >= lNum && palNum1 <= rNum && isPalindromes(palNum1) {
			res++
		}
		if palNum2 >= lNum && isPalindromes(palNum2) {
			res++
		}
	}
}

func main() {
	superpalindromesInRange("4", "1000")
}
