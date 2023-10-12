package offer2

func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	v := [26]int{}
	for i := 0; i < len(s1); i++ {
		v[s1[i]-'a']++
	}

	for i := 0; i < len(s2); i++ {
		v[s2[i]-'a']--
	}

	for _, val := range v {
		if val != 0 {
			return false
		}
	}

	return true

}
