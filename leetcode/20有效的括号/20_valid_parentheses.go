package main

import "strings"

type strQueue struct {
	val  string
	next *strQueue
}

// EnQueue 入队
func (q *strQueue) EnQueue(str string) {
	if q.val == "" {
		q.val = str
	} else {
		q.next.EnQueue(str)
	}
}

// DeQueue 出队
func (q *strQueue) DeQueue(str string) {
	if q.val == "" {
		q.val = str
	} else {
		q.next.EnQueue(str)
	}
}

func isValid(s string) bool {
	strSplilt := strings.Split(s, "")

}

func main() {
	str := "()[]{}"
}
