package main

import (
	"fmt"
	"mygo/pg"
)

func main() {
	pg.Pg01(false)
	pg.Pg02(false)
	pg.Pg03(false)
	pg.Pg04(false)
	pg.Pg05(false)
	pg.Pg06(true)

	fmt.Println("exit")
	fmt.Println("exit")
}
