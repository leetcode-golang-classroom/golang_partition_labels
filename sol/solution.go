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
