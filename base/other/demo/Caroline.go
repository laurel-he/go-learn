package main

import (
	"fmt"
)

/*
1 kj = 0.239卡
100g 2091
65g
324.7卡
*/
func getRightCaroline(kj float64, kilo float64, minRes float64) (res float64) {
	if minRes != 0 {
		return kj / 100 * kilo * 0.239 * minRes
	} else {
		return kj / 100 * kilo * 0.239
	}
}

func testChan() {
	testCh := make(chan int, 20)
	for i := 0; i < 10; i++ {
		testCh <- i
	}
	select {
	case tmp := <-testCh:
		fmt.Println("tmp:", tmp)
	default:
		fmt.Println("timeout...")
		break
	}
}
func main() {
	ks := 2091
	kilo := 65
	minres := 0.8
	testChan()
	fmt.Println("_____res______", getRightCaroline(float64(ks), float64(kilo), minres))
}
