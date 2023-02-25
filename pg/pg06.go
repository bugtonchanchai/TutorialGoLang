package pg

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type Person struct {
	Name string
	Age  int
}

type MyError struct {
	When time.Time
	What string
}

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

	//assertions
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//switch assertion
	do("Hello")
	do(105)
	do(10.10)

	// Stringers
	a := Person{"Smith", 24}
	b := Person{"Emonda", 40}
	fmt.Println(a, b)

	// Error
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// Reader
	r := strings.NewReader("Hello Reader!")
	bb := make([]byte, 8)
	for {
		n, err := r.Read(bb)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", bb[:n])
		if err == io.EOF {
			break
		}
	}
}

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func run() error {
	return &MyError{
		time.Now(),
		"It did't work"}
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}
