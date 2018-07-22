package main

import (
	"fmt"
)

func ToString(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return x
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))

	}
}

func main() {
	var data = struct {
		a int
		b uint
		c bool
		d string
		e float32
	}{
		a: 1,
		b: 2,
		c: true,
		d: "ok",
		e: 3.14,
	}

	fmt.Println(data.a)
	fmt.Println(data.b)
	fmt.Println(data.c)
	fmt.Println(data.d)
	//fmt.Println(data.a)
}
