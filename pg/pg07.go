package pg

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Pg07(isActive bool) {
	if !isActive {
		return
	}
	fmt.Println("Welcome to playground 07 (restfull api)")

	url := "https://jsonplaceholder.typicode.com/todos/1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
