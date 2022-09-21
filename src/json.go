package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	result struct {
		// Notice that we are using go tags here which are used
	        // by the json marshaller.
		UserId int `json:"userId"` 
		Id int `json:"lower_case_id"`
		Title string `json:"title"`
		Competed bool `json:"content"`
		Missing string `json:"missing,omitempty"`
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

	prefix := ""
	indent := "    "
	pretty, err := json.MarshalIndent(r, prefix, indent)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(string(pretty))
}
