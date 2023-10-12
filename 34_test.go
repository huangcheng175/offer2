package offer2

func isAlienSorted(words []string, order string) bool {

	if len(words) <= 1 {
		return true
	}
	sortOrders := sortOrder(order)
	p1 := 0
	p2 := 1
	for p2 < len(words) {
		if !compareAlien(words[p1], words[p2], sortOrders) {
			return false
		}
		p1++
		p2++
	}

	return true

}

func sortOrder(order string) (orders [26]uint8) {
	for i := 0; i < len(order); i++ {
		orders[order[i]-'a'] = uint8(i) + 1
	}

	return
}

func compareAlien(s1, s2 string, orders [26]uint8) bool {
	p1 := 0
	var (
		v1, v2 uint8
	)
	for p1 < len(s1) && p1 < len(s2) {
		v1 = orders[s1[p1]-'a']
		v2 = orders[s2[p1]-'a']

		if v1 < v2 {
			return true
		} else if v1 > v2 {
			return false
		} else {
			p1++
		}
	}

	return len(s1) <= len(s2)
}
