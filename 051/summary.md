### 二分答案法
1. 估计 最终答案的可能范围
2. 分析 **问题的答案**和**给定条件** 的 单调性
3. 建立函数, 计算已知答案的情况下, 判断是否符合条件
4. 在可能的答案上二分, 找到最终的答案
```go
l, r := L, R        // 搜索区间 [l, r)
for l < r {
    mid := l + (r - l) / 2
    if check(mid) {    // mid 可行
        r = mid        // 保留 mid
    } else {
        l = mid + 1
    }
}
return l   // 或 r，二者相等

```
```go
func binFindLess(arr []int, target int, len int) int {
	l, r := 0, len-1
	ans := -1

	for l <= r {
		mid := (l + r) / 2
		if arr[mid] < target {
			ans = mid
			r = mid - 1 // 继续向左找
		} else {
			l = mid + 1
		}
	}

	return ans
}
```