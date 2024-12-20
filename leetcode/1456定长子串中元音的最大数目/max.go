package main

import "fmt"

//func maxVowels(s string, k int) int {
//	strSplit := strings.Split(s, "")
//	if k > len(strSplit) {
//		return 0
//	}
//	for i := 0; i < len(strSplit)-k; i++ {
//		strRes := getStrRes(strSplit, i, k)
//		getVowelsNum(strRes)
//	}
//
//}

// getStrRes 获取
func getStrRes(strRes []string, begin int, k int) (res []string) {
	if begin > len(strRes) {
		return []string{}
	}
	for i := begin; i < len(strRes); i++ {
		if i < (begin + k) {
			res = append(res, strRes[i])
		}
	}
	return res
}

// getVowelsNum 获取一个字符串中元音的数目
func getVowelsNum(strRes []string) int {
	vowels := map[string]int{
		"a": 1,
		"e": 2,
		"i": 3,
		"o": 4,
		"u": 5,
	}
	var count int
	for i := 0; i < len(strRes); i++ {
		if _, ok := vowels[strRes[i]]; ok {
			count++
			continue
		}
	}
	return count
}

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

type Set map[string]struct{}

func main() {
	var a uint = 1
	var b uint = 2
	fmt.Println(a - b)
}
