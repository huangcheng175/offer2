package offer2

import (
	"fmt"
	"testing"
)

/*
[[0,1,2,3,4,5,6,7,8,9],[19,18,17,16,15,14,13,12,11,10],[20,21,22,23,24,25,26,27,28,29],[39,38,37,36,35,34,33,32,31,30],[40,41,42,43,44,45,46,47,48,49],[59,58,57,56,55,54,53,52,51,50],[60,61,62,63,64,65,66,67,68,69],[79,78,77,76,75,74,73,72,71,70],[80,81,82,83,84,85,86,87,88,89],[99,98,97,96,95,94,93,92,91,90],[100,101,102,103,104,105,106,107,108,109],[119,118,117,116,115,114,113,112,111,110],[120,121,122,123,124,125,126,127,128,129],[139,138,137,136,135,134,133,132,131,130],[0,0,0,0,0,0,0,0,0,0]]
*/

func TestLongest(t *testing.T) {
	fmt.Println(longestIncreasingPath([][]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{19, 18, 17, 16, 15, 14, 13, 12, 11, 10},
		{20, 21, 22, 23, 24, 25, 26, 27, 28, 29},
		{39, 38, 37, 36, 35, 34, 33, 32, 31, 30},
		{40, 41, 42, 43, 44, 45, 46, 47, 48, 49},
		{59, 58, 57, 56, 55, 54, 53, 52, 51, 50},
		{60, 61, 62, 63, 64, 65, 66, 67, 68, 69},
		{79, 78, 77, 76, 75, 74, 73, 72, 71, 70},
		{80, 81, 82, 83, 84, 85, 86, 87, 88, 89},
		{99, 98, 97, 96, 95, 94, 93, 92, 91, 90},
		{100, 101, 102, 103, 104, 105, 106, 107, 108, 109},
		{119, 118, 117, 116, 115, 114, 113, 112, 111, 110},
		{120, 121, 122, 123, 124, 125, 126, 127, 128, 129},
		{139, 138, 137, 136, 135, 134, 133, 132, 131, 130},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}))
	//longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}})
	//longestIncreasingPath([][]int{{1}})
	//longestIncreasingPath([][]int{{1, 2}})
	//longestIncreasingPath([][]int{{1}, {2}})
	//longestIncreasingPath([][]int{{1, 2}, {3, 4}})
	//longestIncreasingPath([][]int{{1, 2, 3, 4, 5}, {16, 17, 24, 23, 6}, {15, 18, 25, 22, 7}, {14, 19, 20, 21, 8}, {13, 12, 11, 10, 9}})
	//longestIncreasingPath([][]int{{1, 2, 3, 4, 5}, {16, 17, 24, 23, 6}, {15, 18, 25, 22, 7}, {14, 19, 20, 21, 8}, {13, 12, 11, 10, 9}})
	//longestIncreasingPath([][]int{{1, 2, 3, 4, 5}, {16, 17, 24, 23, 6}, {15, 18, 25, 22, 7}, {14, 19, 20, 21, 8}, {13, 12, 11, 10, 9}})
	//longestIncreasingPath([][]int{{1, 2, 3, 4, 5}, {16, 17, 24, 23, 6}, {15, 18, 25, 22, 7}, {14, 19, 20, 21, 8}, {13, 12, 11, 10, 9}})
}

var (
	dirs          = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}
	rows, columns int
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, columns = len(matrix), len(matrix[0])
	memo := make([][]int, rows)
	for i := 0; i < rows; i++ {
		memo[i] = make([]int, columns)
	}
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			ans = max(ans, dfs(matrix, i, j, memo))
		}
	}
	return ans
}

func dfs(matrix [][]int, row, column int, memo [][]int) int {
	if memo[row][column] != 0 {
		return memo[row][column]
	}
	memo[row][column]++
	for _, dir := range dirs {
		newRow, newColumn := row+dir[0], column+dir[1]
		if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] > matrix[row][column] {
			memo[row][column] = max(memo[row][column], dfs(matrix, newRow, newColumn, memo)+1)
		}
	}
	return memo[row][column]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//
//func longestIncreasingPath(matrix [][]int) int {
//	longest := -1
//	var dfs func(longest int, i, j int)
//	dfs = func(length int, i, j int) {
//
//		//	fmt.Println(length, i, j, matrix[i][j])
//		if i >= 1 && matrix[i-1][j] > matrix[i][j] {
//			//fmt.Println("up..", length, i, j, matrix[i][j])
//			dfs(length+1, i-1, j)
//		}
//
//		if i < len(matrix)-1 && matrix[i+1][j] > matrix[i][j] {
//			//fmt.Println("down..", length, i, j, matrix[i][j])
//			dfs(length+1, i+1, j)
//		}
//
//		if j >= 1 && matrix[i][j-1] > matrix[i][j] {
//			//fmt.Println("left..", length, i, j, matrix[i][j])
//			dfs(length+1, i, j-1)
//		}
//
//		if j < len(matrix[i])-1 && matrix[i][j+1] > matrix[i][j] {
//			//fmt.Println("right..", length, i, j, matrix[i][j])
//			dfs(length+1, i, j+1)
//		}
//
//		longest = max(longest, length)
//
//	}
//
//	for i := 0; i < len(matrix); i++ {
//		for j := 0; j < len(matrix[i]); j++ {
//			fmt.Println("i, j", i, j)
//			dfs(1, i, j)
//		}
//	}
//
//	return longest
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//
//	return b
//}
