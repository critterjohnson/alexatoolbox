package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/critterjohnson/alexatoolbox/request"
	"github.com/critterjohnson/alexatoolbox/response"
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
		return
	}

	responseBuilder := response.NewBuilder()
	response := responseBuilder.
		WithTextOutputSpeech("hello world").
		WithSimpleCard("Cool Card", "This card is so cool!").
		AddAttribute("cool", "beans").
		Build()
	responseJSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON", err)
		return
	}
	fmt.Println(string(responseJSON))
}
