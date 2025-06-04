package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 获取最大值
 * @param a int整型
 * @param b int整型
 * @return int整型
 */

// 符号位
func sign(a int32) int {
	return filp(int32(uint32(a) >> 31))
}

// 取反
func filp(a int32) int {
	return int(1 ^ a)
}

// 防溢出
func getMax(a, b int) int {
	c := a - b
	sa := sign(int32(a))
	sb := sign(int32(b))
	sc := sign(int32(c))
	// 判断A和B符号
	diffAB := sa ^ sb
	sameAB := filp(int32(diffAB))
	// 返回A的条件
	// 1. AB异号(可能溢出), A非负, B就为负，A > B
	// 2. AB同号(不会溢出), C非负, A > B
	returnA := diffAB*sa + sameAB*sc
	returnB := filp(int32(returnA))
	return a*returnA + b*returnB
}
