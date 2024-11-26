package main

import (
	"fmt"
	"strings"
)

type nonRepeatQueue struct {
	items []string
}

func (q *nonRepeatQueue) EnQueue(s string) {
	q.items = append(q.items, s)
}

func (q *nonRepeatQueue) DeQueue() string {
	if len(q.items) < 0 {
		return ""
	}
	if len(q.items) == 0 {
		q.items = []string{}
		return q.items[0]
	}
	temp := q.items[0]
	q.items = q.items[1:]
	return temp
}

//
//func longestPalindrome(s string) string {
//	//maxPalindrome := make()
//}

// checkIsPalindrome 检测一个字符串是否是回文字符串
func checkIsPalindrome(s string) bool {
	stringSlice := strings.Split(s, "")
	strLen := len(stringSlice)
	for i := 0; i < strLen; i++ {
		for j := strLen - 1; j >= 0; i++ {
			if stringSlice[i] != stringSlice[j] {
				return false
			}
		}
	}
	return true
}

func GetNonRepeatArr(str string) int {
	stringSlice := strings.Split(str, "")
	var queue nonRepeatQueue
	var lenArr []int
	for _, strItem := range stringSlice {
		_, checkIn := checkInSlice(strItem, queue)
		queue.EnQueue(strItem)
		if checkIn == true {
			for {
				deQueueItem := queue.DeQueue()
				if deQueueItem == strItem {
					lenArr = append(lenArr, len(queue.items))
					break
				} else {
					continue
				}
			}
		} else {
			lenArr = append(lenArr, len(queue.items))
		}
	}
	return checkMax(lenArr)
}

func checkInSlice(str string, queue nonRepeatQueue) (int, bool) {
	for k, item := range queue.items {
		if item == str {
			return k, true
		}
	}
	return 0, false
}

func checkMax(lenArr []int) int {
	var max int
	for _, item := range lenArr {
		if item >= max {
			max = item
		}
	}
	return max
}
func main() {
	// abcabcbb 3
	//  bbbbb 1
	//pwwkew 3
	//aab 2
	//dvdf 3
	fmt.Println(GetNonRepeatArr("dvdf"))
}
