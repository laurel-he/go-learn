package main

import "fmt"

// dealArray 测试捕获数组越界错误
func dealArray() {
	defer func() {
		fmt.Println("____________err_____")
		if r := recover(); r != nil {
			fmt.Println("++++++++++++err++++++", r)
		}
	}()
	arrA := [3]int{1, 2, 3}
	fmt.Println("arrA:", arrA)
	for i := 0; i < 10; i++ {
		fmt.Println("try to print:", arrA[i])
	}
}

func testArrLen(arr [5]int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("___recover res is: __", r)
		}
	}()
	fmt.Println("_______arr len is__:", len(arr))
	fmt.Println("_____arr cap res", cap(arr))
	for i := 0; i < 10; i++ {
		fmt.Printf("____test__arr__res__: i:%v, res:%v", i, arr[i])
	}
}

func main() {
	go dealArray()
	go testArrLen([...]int{3, 2, 1, 10, 3})
	fmt.Println("_____done_____")
}
