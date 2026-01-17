package kata

import (
	"slices"
	"strings"
)

// ::KATA START::
func groupAnagrams(strs []string) [][]string {
	anagrams := map[string][]string{}
	for _, s := range strs {
		key := sortStr(s)
		anagrams[key] = append(anagrams[key], s)
	}

	answer := make([][]string, 0, len(anagrams))
	for _, group := range anagrams {
		answer = append(answer, group)
	}
	return answer
}

func sortStr(s string) string {
	runes := []rune(strings.ToLower(s))
	slices.Sort(runes)
	return string(runes)
}

// ::KATA END::
