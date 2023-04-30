package handlers

import (
	"encoding/json"
	"github.com/Thajun/test2/data"
	"github.com/Thajun/test2/services"
	"log"
	"net/http"
)

type Test2Handler struct {
	l *log.Logger
}

func NewTest2Handler(l *log.Logger) *Test2Handler {
	return &Test2Handler{l}
}

func (t *Test2Handler) GetTop10UsedWords(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var text data.Test2Request
	err := decoder.Decode(&text)
	if err != nil {
		log.Println("Error occurred while decoding request", err)
	}

	log.Println("Request Text for Testing: ", text.InputText)

	result, err := services.GetWordCount(text)
	if err != nil {
		log.Println("Error occurred while getting word count. error: ", err)
	}

	json.NewEncoder(w).Encode(result)
}
