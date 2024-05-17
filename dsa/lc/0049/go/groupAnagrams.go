package leetcode

import (
	"slices"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)
	for _, s := range strs {
		k := [26]int{}
		for _, v := range s {
			k[v-'a']++
		}
		groups[k] = append(groups[k], s)
	}

	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}
	return result
}

type sortRunes []rune

func (sr sortRunes) Len() int {
	return len(sr)
}

func (sr sortRunes) Less(i, j int) bool {
	return sr[i] < sr[j]
}

func (sr sortRunes) Swap(i, j int) {
	sr[i], sr[j] = sr[j], sr[i]
}

func groupAnagramsRunes(strs []string) [][]string {
	groups := map[string][]string{}
	for _, str := range strs {
		runes := []rune(str)
		sort.Sort(sortRunes(runes))
		sortedString := string(runes)
		groups[sortedString] = append(groups[sortedString], str)
	}

	res := make([][]string, 0, len(groups))
	for _, v := range groups {
		res = append(res, v)
	}
	return res
}

func groupAnagramsRunesV2(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, s := range strs {
		runes := []rune(s)
		slices.SortFunc(runes, func(a, b rune) int {
			if a < b {
				return -1
			} else if a > b {
				return 1
			}
			return 0
		})
		sortedString := string(runes)
		groups[sortedString] = append(groups[sortedString], s)
	}

	res := make([][]string, 0, len(groups))
	for _, v := range groups {
		res = append(res, v)
	}
	return res
}
