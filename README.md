# golang_partition_labels

You are given a string `s`. We want to partition the string into as many parts as possible so that each letter appears in at most one part.

Note that the partition is done so that after concatenating all the parts in order, the resultant string should be `s`.

Return *a list of integers representing the size of these parts*.

## Examples

**Example 1:**

```
Input: s = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
This is a partition so that each letter appears in at most one part.
A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits s into less parts.

```

**Example 2:**

```
Input: s = "eccbbbbdec"
Output: [10]

```

**Constraints:**

- `1 <= s.length <= 500`
- `s` consists of lowercase English letters.

## 解析

給定一個字串 s, 

希望透過一個字串切割的方式，儘可能把所有相同字元切割在同一個子字串，然後儘可能把不同字元切割到不同子字串

要求寫一個演算法來計算透過上述的切割方式來找到切割後每個子字串的長度

對於每個字元因為希望儘可能切割到同一個子字串

所以需要紀錄每個字元出現的最後位置 

這邊以 HashMap lastPos 紀錄每個字元的最後位置

這樣每次透過 loop 每個字元都可以透過 lastPos 查到該字元同子字串最大延伸位置

![](https://i.imgur.com/iLLHBux.png)
![](https://i.imgur.com/ewVHe7D.jpg)

透過每次 greedy 採用最遠延長距離

start, end 初始化為 0

result 初始化為空陣列

每次比較 lastPos[s[idx]] 是否大於 end

如果是更新 end = lastPos[s[idx]]

當 idx == end 時 代表已做完這次分割

把 end - start +1 放入 result 中

更新 start = end + 1

這樣當遍歷完所有 s 的字元後

result 及為所求

時間複雜度是 O(N)

因為需要先用 hashMap 去紀錄每個字元的最後位置

所以空間複雜度是 O(N)

## 程式碼
```go
package sol

func partitionLabels(s string) []int {
	result := []int{}
	sLen := len(s)
	lastPos := make(map[byte]int)
	for idx := 0; idx < sLen; idx++ {
		lastPos[s[idx]] = idx
	}
	start, end := 0, 0
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for idx := 0; idx < sLen; idx++ {
		end = max(end, lastPos[s[idx]])
		if idx == end {
			// update result and start
			result = append(result, end-start+1)
			start = end + 1
		}
	}
	return result
}
```
## 困難點

1. 需要看出每次要切割都必須要找當下字元所屬子字串的最遠延伸
2. 理解更新 起始位置的時機點

## Solve Point

- [x]  建立一個 hashMap lastPos 來存儲每個字元的最後出現位置
- [x]  遍歷字串 s 所有位置的字元更新 lastPos
- [x]  初始化 start = 0, end = 0, result = 空陣列
- [x]  遍歷字串 s 所有位置的字元 更新 end = max(尋當下字元最後位置, end)
- [x]  如果 idx == end , 把 end - start +1 加入 result, 更新 start = idx + 1
- [x]  重複上述步驟直到 idx= len(s) -1
- [x]  回傳 result