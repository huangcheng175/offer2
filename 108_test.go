package offer2

import "testing"

func TestLadderLength(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	t.Log(ladderLength(beginWord, endWord, wordList))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {

	q1 := make([]string, 0)
	q2 := make([]string, 0)

	v1 := make(map[string]bool)
	for _, v := range wordList {
		v1[v] = true
	}

	if !v1[endWord] {
		return 0
	}

	q1 = append(q1, beginWord)
	visited := make(map[string]bool)
	depth := 0

	for len(q1) > 0 {
		q2 = make([]string, 0)
		for _, v := range q1 {
			if v == endWord {
				return depth + 1
			}

			if visited[v] {
				continue
			}

			visited[v] = true
			q2 = append(q2, getNeighborWords(v, wordList)...)
		}
		q1 = q2
		depth++

	}

	return 0

}

func getNeighborWords(word string, wordList []string) []string {

	ret := make([]string, 0)
	for _, v := range wordList {
		if v == word {
			continue
		}

		if isNeighbor(word, v) {
			ret = append(ret, v)
		}

	}

	return ret

}

func isNeighbor(a, b string) bool {
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff == 1
}
