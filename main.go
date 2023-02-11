package main

import (
	"fmt"
	"mygo/pg"
)

func main() {
	pg.Pg01(false)
	pg.Pg02(false)
	pg.Pg03(true)

	fmt.Println("exit")
}
