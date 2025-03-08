package kata

func lengthOfLongestSubstring(substring string) int {
	left := 0
	result := 0
	seen := map[rune]int{}
	for i, r := range substring {
		if lastSeen, exists := seen[r]; exists && lastSeen >= left {
			left = lastSeen + 1
		}
		seen[r] = i
		result = max(result, i-left+1)
	}
	return result
}

// {p:0}			left=0 res=1
// {p:0,w:1}			left=0 res=2
// {p:0,w:2}			left=2 res=2
// {p:0,w:2,k:3}			left=2 res=2
// {p:0,w:2,k:3,e:4}			left=2 res=3
