package offer2

import (
	"fmt"
	"testing"
)

func TestOpenLock(t *testing.T) {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))

}
func openLock(deadends []string, target string) int {
	q1 := make([]string, 0)
	q2 := make([]string, 0)

	v1 := make(map[string]bool)
	for _, v := range deadends {
		v1[v] = true
	}

	if v1["0000"] {
		return -1
	}
	if v1[target] {
		return -1
	}

	q1 = append(q1, "0000")

	visited := make(map[string]bool)
	depth := 0

	for len(q1) > 0 {
		q2 = make([]string, 0)
		for _, v := range q1 {
			if v == target && !v1[v] {
				return depth
			}

			if visited[v] {
				continue
			}

			if v1[v] {
				continue
			}

			visited[v] = true
			q2 = append(q2, getNeighbor(v)...)
		}
		q1 = q2
		depth++

	}

	return -1
}

func getNeighbor(str string) []string {
	ret := make([]string, 0)
	for i := 0; i < len(str); i++ {
		v := str[i]
		if v == '0' {
			ret = append(ret, str[:i]+"1"+str[i+1:])
			ret = append(ret, str[:i]+"9"+str[i+1:])
		} else if v == '9' {
			ret = append(ret, str[:i]+"8"+str[i+1:])
			ret = append(ret, str[:i]+"0"+str[i+1:])
		} else {
			ret = append(ret, str[:i]+string(v+1)+str[i+1:])
			ret = append(ret, str[:i]+string(v-1)+str[i+1:])
		}
	}

	//fmt.Println("getNeighbor...", str, ret)
	return ret
}
