package offer2

import (
	"fmt"
	"testing"
)

func TestFindOrder(t *testing.T) {
	//fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}
func findOrder(numCourses int, prerequisites [][]int) []int {
	g := newGraph()

	for _, p := range prerequisites {
		g.addEdge(p[1], p[0])
	}

	order := []int{}
	queue := []int{}

	for i := 0; i < numCourses; i++ {
		if g.inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		u := queue[0]
		order = append(order, u)

		queue = queue[1:]

		for _, v := range g.edges[u] {
			g.inDegree[v]--
			if g.inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	for i := 0; i < numCourses; i++ {
		if g.inDegree[i] != 0 {
			return []int{}
		}
	}

	return order
}

type graph struct {
	edges    map[int][]int
	inDegree map[int]int
}

func newGraph() graph {
	return graph{
		edges:    make(map[int][]int),
		inDegree: make(map[int]int),
	}
}

func (g graph) addEdge(u, v int) {
	if g.edges == nil {
		g.edges = make(map[int][]int)
	}
	g.edges[u] = append(g.edges[u], v)
	g.inDegree[v]++
}
