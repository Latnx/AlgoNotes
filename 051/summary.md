### 二分答案法
1. 估计 最终答案的可能范围
2. 分析 **问题的答案**和**给定条件** 的 单调性
3. 建立函数, 计算已知答案的情况下, 判断是否符合条件
4. 在可能的答案上二分, 找到最终的答案
```go
l, r := L, R
ans := some_default
for l <= r {
    mid := (l + r) / 2
    if check(mid) {   // mid 可行
        ans = mid     // 保存答案
        r = mid - 1   // 左移，找更小的可行值
    } else {
        l = mid + 1   // mid 不可行，丢弃左区间
    }
}
return ans

```