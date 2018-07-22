package main

import "fmt"

// equal 字典比较
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

// appendInt 追加元素
func appendInt(x []int, y int) []int {
	var z []int

	zLen := len(x) + 1
	if zLen <= cap(x) {
		z = x[:zLen]
	} else {
		zCap := zLen
		if zCap < 2*len(x) {
			zCap = 2 * len(x)
		}

		z = make([]int, zLen, zCap)
		copy(z, x)
	}

	z[len(x)] = y
	return z
}
func main() {
	var x, y []int

	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap = %d\t%v\n", i, cap(y), y)
		x = y
	}
}
