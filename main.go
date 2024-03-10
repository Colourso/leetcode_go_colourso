package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("hello")

	rand.Seed(time.Now().Unix())

	dp := make([][]int, 2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, rand.Intn(10)+1)
		fmt.Println(len(dp[i]))
	}

	dp[1][0] = 1
	fmt.Println(dp)

	beibao01 := Beibao01()
	fmt.Println(beibao01)
}

func Beibao01() int {
	M, N := 0, 0
	fmt.Scanln(&M, &N)
	weight := make([]int, M)
	value := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Scan(&weight[i])
	}
	for i := 0; i < M; i++ {
		fmt.Scan(&value[i])
	}

	fmt.Println(M)
	fmt.Println(N)
	fmt.Println(weight)
	fmt.Println(value)

	dp := make([][]int, M)
	for i := range dp {
		dp[i] = make([]int, N+1)
	}
	for i := 0; i <= N; i++ {
		if weight[0] <= i {
			dp[0][i] = value[0]
		}
	}
	fmt.Println(dp)

	// 先物品、后背包
	for i := 1; i < M; i++ {
		for j := 1; j <= N; j++ {
			if weight[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = getMax(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
		fmt.Println(dp)
	}

	return dp[M-1][N]
}

func getMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
