package offer2

import "sort"

//func groupAnagrams(strs []string) [][]string {
//	res := make([][]string, 0)
//	countsMap := make(map[[26]int][]string)
//
//	for _, str := range strs {
//		counts := countStr(str)
//		countsMap[counts] = append(countsMap[counts], str)
//	}
//
//	for _, v := range countsMap {
//		res = append(res, v)
//	}
//
//	return res
//}
//
//func countStr(str string) [26]int {
//	res := [26]int{}
//	for i := 0; i < len(str); i++ {
//		res[str[i]-'a']++
//	}
//
//	return res
//}

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)

	resMap := make(map[string][]string)
	for _, str := range strs {
		vb := []byte(str)
		sort.Slice(vb, func(i, j int) bool {
			return vb[i] < vb[j]
		})
		newStr := string(vb)
		resMap[newStr] = append(resMap[newStr], str)
	}

	for _, v := range resMap {
		res = append(res, v)
	}

	return res
}
