package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://random-data-api.com/api/v2/users")
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]string)

	json.Unmarshal(body, &data)

	fmt.Println(data["uid"])
	fmt.Println(data["first_name"])
	fmt.Println(data["last_name"])
	fmt.Println(data["username"])
	fmt.Println(data["email"])
}
