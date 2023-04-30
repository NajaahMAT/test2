package services

import (
	"bytes"
	"encoding/json"
	"github.com/Thajun/test2/data"
	"log"
	"net/http"
)

func GetWordCount(req data.Test2Request) (res data.Test2Response, err error) {

	request, err := json.Marshal(req)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(data.MangtasTest1ServiceUrl+data.MangtasTest1WordCountEndpoint, "application/json",
		bytes.NewBuffer(request))
	if err != nil {
		log.Println("Error occurred while calling MangtasTest1 service. Error: ", err)
		return res, err
	}

	json.NewDecoder(resp.Body).Decode(&res)

	return res, nil
}
