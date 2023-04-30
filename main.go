package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Thajun/test2/data"
	"github.com/Thajun/test2/handlers"
	httphandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const Port = ":9091"

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler path is incorrect")
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(data.MyApplicationRegistryResponse{Status: false, Result: "Incorrect path/values"})
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Go server is starting on Port: ", Port)

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	logTest2Deal := log.New(os.Stdout, "Test 1 Deals", log.LstdFlags|log.Lshortfile)
	test2Handler := handlers.NewTest2Handler(logTest2Deal)

	router.Methods(http.MethodPost).Subrouter().HandleFunc("/mangtas/test2", test2Handler.GetTop10UsedWords)

	ch := httphandler.CORS(
		httphandler.AllowedMethods([]string{"POST"}),
		httphandler.AllowedHeaders([]string{"Content-type"}),
		httphandler.AllowedOrigins([]string{"*", "localhost"}),
		httphandler.AllowCredentials(),
	)

	server := &http.Server{
		Addr:         Port,
		Handler:      ch(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  100 * time.Second,
		WriteTimeout: 100 * time.Second,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	//for gracefully shutdown
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Received request to terminate the server", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
