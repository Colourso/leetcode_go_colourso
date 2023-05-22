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
}
