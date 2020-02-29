package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/critterjohnson/alexatoolbox/handlers"
	"github.com/critterjohnson/alexatoolbox/request"
	"github.com/critterjohnson/alexatoolbox/response"
)

func SonnetHandler(request request.Request) (response.Response, error) {
	return response.NewBuilder().
		WithTextOutputSpeech(request.RequestBody.Intent.Slots["SONNET_NUMBER"].Value).
		Build(), nil
}

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

	requestHandler := handlers.NewRequestHandler()
	requestHandler.AddIntentRequestHandler("readSonnet", SonnetHandler)
	response := requestHandler.Handle(request)

	responseJSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON", err)
		return
	}
	fmt.Println(string(responseJSON))
}
