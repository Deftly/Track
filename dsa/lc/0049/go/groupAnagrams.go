package leetcode

import "sort"

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
	record, res := map[string][]string{}, [][]string{}
	for _, str := range strs {
		runes := []rune(str)
		sort.Sort(sortRunes(runes))
		sortedString := string(runes)
		record[sortedString] = append(record[sortedString], str)
	}

	for _, v := range record {
		res = append(res, v)
	}
	return res
}
