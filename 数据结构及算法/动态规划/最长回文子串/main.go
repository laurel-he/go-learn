package main

import "strings"

func longestPalindrome(s string) string {
	return ""
}

type nonRepeatQueue struct {
	items []string
}

func (q *nonRepeatQueue) EnQueue(s string) string {
	q.items = append(q.items, s)
	return strings.Join(q.items, "")
}

func (q *nonRepeatQueue) DeQueue() string {
	if len(q.items) < 0 {
		return ""
	}
	if len(q.items) == 0 {
		q.items = []string{}
		return q.items[0]
	}
	q.items = q.items[1:]
	return strings.Join(q.items, "")
}

func (q *nonRepeatQueue) longestPalindrome(s string) string {
	allStrArr := strings.Split(s, "")
	for _, strItem := range allStrArr {
		lastItem := q.EnQueue(strItem)
		if checkIsPalindrome(lastItem) == true {

		}
	}
}

func checkIsPalindrome(str string) bool {
	if str == reverseString(str) {
		return true
	}
	return false
}

func reverseString(s string) string {
	r := []rune(s) // 将字符串转换为 rune 切片，以处理 Unicode 字符
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i] // 交换字符
	}
	return string(r) // 转回字符串
}

func main() {
	str := "babad"
	longestPalindrome(str)
}
