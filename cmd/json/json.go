package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	result struct {
		UserId int `json:"userId"`
		Id int `json:"id"`
		Title string `json:"title"`
		Competed bool `json:"content"`
	}

)

func main() {
	uri := "https://jsonplaceholder.typicode.com/todos/1"
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()

	var r result
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(r)

	pretty, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(string(pretty))
}
