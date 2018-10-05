package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tiltify"
)

func main() {

	url := "https://tiltify.com/api/v3"

	resp, err := http.Get(url + "/users")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	var userResp tiltify.UserResponse
	err = json.Unmarshal(b, &userResp)
	if err != nil {
		panic(err)
	}

	fmt.Println(*userResp.Data)

}
