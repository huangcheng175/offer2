package offer2

import (
	"fmt"
	"testing"
)

func TestAllPathsSourceTarget(t *testing.T) {

	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
}

func allPathsSourceTarget(graph [][]int) [][]int {
	ret := make([][]int, 0)

	q1 := make([][]int, 0)

	q1 = append(q1, []int{0})

	for len(q1) > 0 {
		q2 := make([][]int, 0)
		for _, v := range q1 {
			if v[len(v)-1] == len(graph)-1 {
				ret = append(ret, v)
				continue
			}

			for _, vv := range graph[v[len(v)-1]] {
				//if vv > len(graph)-1 {
				//	continue
				//}
				q2 = append(q2, append(append([]int{}, v...), vv))
			}
		}
		q1 = q2
	}

	return ret

}
