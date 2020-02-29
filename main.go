package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/critterjohnson/alexatoolbox/request"
)

func main() {
	data, err := ioutil.ReadFile("sonnet138.json")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	request := request.Request{}
	if err = json.Unmarshal(data, &request); err != nil {
		fmt.Println("Error unmarshalling JSON", err)
	}
	fmt.Printf("%+v%v", request, request)
}
