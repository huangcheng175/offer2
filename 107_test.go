package offer2

import (
	"fmt"
	"testing"
)

func TestUpdateMatrix(t *testing.T) {
	fmt.Println(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}))
}

func updateMatrix(mat [][]int) [][]int {

	ret := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		ret[i] = make([]int, len(mat[i]))
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				ret[i][j] = getMaxDistance(mat, i, j)
			}
		}
	}

	return ret
}

func getMaxDistance(mat [][]int, i, j int) int {

	if mat[i][j] == 0 {
		return 0
	}

	q1 := make([]int, 0)
	q2 := make([]int, 0)

	q1 = append(q1, i*len(mat[0])+j)

	visited := make(map[int]bool)
	depth := 0

	for len(q1) > 0 {
		q2 = make([]int, 0)
		for _, v := range q1 {
			x, y := v/len(mat[0]), v%len(mat[0])
			if mat[x][y] == 0 {
				return depth
			}

			if visited[v] {
				continue
			}

			visited[v] = true
			q2 = append(q2, getNeighbor(mat, x, y)...)
		}
		q1 = q2
		depth++
	}

	return depth

}

func getNeighbor(mat [][]int, i, j int) []int {
	ret := make([]int, 0)
	if i-1 >= 0 {
		ret = append(ret, (i-1)*len(mat[0])+j)
	}
	if i+1 < len(mat) {
		ret = append(ret, (i+1)*len(mat[0])+j)
	}
	if j-1 >= 0 {
		ret = append(ret, i*len(mat[0])+j-1)
	}
	if j+1 < len(mat[0]) {
		ret = append(ret, i*len(mat[0])+j+1)
	}

	return ret
}
