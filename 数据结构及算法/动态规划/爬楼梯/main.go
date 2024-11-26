package main

import "fmt"

//// climbStairs 通过递归实现
//func climbStairs(n int) int {
//	if n == 0 {
//		return 0
//	}
//	if n == 1 {
//		return 1
//	}
//	if n == 2 {
//		return 2
//	}
//	return climbStairs(n-1) + climbStairs(n-2)
//}

// climbStairs 通过动态规划实现
func climbStairs(n int) int {
	res := make([]int, n+1)
	res[0] = 1
	res[1] = 1
	res[2] = 2
	for i := 2; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return res[n]
}

func main() {
	fmt.Println(climbStairs(4))
}
