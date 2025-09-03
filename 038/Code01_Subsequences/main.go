package main

// 字符串的全部子序列
// 子序列本身是可以有重复的，只是这个题目要求去重
// 测试链接 : https://www.nowcoder.com/practice/92e6247998294f2c933906fdedbc6e6a

// 其实就是一个遍历二叉树
// 找到对应的二叉树，如何遍历
// 如何记录结果，以及选择不同的分支
func f1(s string, path []byte, index int, set map[string]struct{}) {
	if index == len(s) {
		set[string(path)] = struct{}{}
		return
	}
	path = append(path, s[index])
	f1(s, path, index+1, set)
	path = path[:len(path)-1]
	f1(s, path, index+1, set)
}

func generatePermutation(s string) []string {
	m := map[string]struct{}{}
	res := []string{}
	f1(s, []byte{}, 0, m)
	for k := range m {
		res = append(res, k)
	}
	return res
}
