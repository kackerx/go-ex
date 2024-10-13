package dp

import (
	"fmt"
	"testing"
)

func Test62(t *testing.T) {
	fmt.Println(uniquePaths(3, 3))
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m+1) // m X n的有dp[m][n]条路径到达
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= n; i++ {
		dp[1][i] = 1
	}

	for i := 0; i <= m; i++ {
		dp[i][1] = 1
	}

	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m][n]
}

func Test63(t *testing.T) {
	fmt.Println(uniquePathsWithObstacles(
		[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
	))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = 1
		} else {
			break
		}
	}

	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 0 {
			dp[0][i] = 1
		} else {
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else {
				dp[i][j] = 0
			}
		}
	}

	return dp[m-1][n-1]
}
