package main

func main() {
	if "bab" != longestPalindrome("babad") {
		panic("error")
	}
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	check := func(i, j int) bool {
		left := i
		right := j - 1

		for left < right {
			if s[left] != s[right] {
				return false
			}

			left++
			right--
		}

		return true
	}
	m := make(map[string][]int)
	for i := 0; i < len(s); i++ {
		ss := string(s[i])
		m[ss] = append(m[ss], i)
	}
	start := 0
	end := 1
	for _, idxs := range m {
		if len(idxs) < 2 {
			continue
		}
		for i := 0; i < len(idxs)-1; i++ {
			for j := i + 1; j < len(idxs); j++ {
				ss := s[idxs[i] : idxs[j]+1]
				if len(ss) <= end-start {
					continue
				}
				if check(idxs[i], idxs[j]+1) {
					start = idxs[i]
					end = idxs[j] + 1
				}
			}
		}
	}
	return s[start:end]
}
