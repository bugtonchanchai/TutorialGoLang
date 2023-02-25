package pg

import (
	"fmt"
)

func Pg06(isActive bool) {
	if !isActive {
		return
	}
	fmt.Println("Welcome to playground 06 (empty interface)")

	var i interface{}
	describe2(i)

	i = 42
	describe2(i)

	i = "hello"
	describe2(i)
}

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
