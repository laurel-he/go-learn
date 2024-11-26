package main

import "fmt"

func selectionSort(intArr []int) (res []int) {
	minKey := 0
	length := len(intArr)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if intArr[j] < intArr[minKey] {
				tmp := intArr[j]
				intArr[j] = intArr[minKey]
				intArr[minKey] = tmp
			}
		}
	}
	return intArr
}

func main() {
	test := []int{2, 5, 1, 0, 8, 6, 7, 2}
	fmt.Println(selectionSort(test))
}
