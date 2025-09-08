package main_test

func minDistance(word1 string, word2 string) int {
	pre := make([]int, len(word2)+1)
	cur := make([]int, len(word2)+1)
	for i := range len(pre) {
		pre[i] = i
	}

	for i := 1; i <= len(word1); i++ {
		cur[0] = i
		for j := 1; j < len(pre); j++ {
			if word1[i-1] != word2[j-1] {
				cur[j] = min(cur[j-1], pre[j-1], pre[j]) + 1
			} else {
				cur[j] = pre[j-1]
			}
		}
		pre, cur = cur, pre
	}
	return pre[len(word2)]
}
