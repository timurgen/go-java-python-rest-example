package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	router := mux.NewRouter()
	router.HandleFunc("/whosnext", ProcessRequest).Methods("POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 8080), router))
}

func ProcessRequest(w http.ResponseWriter, r *http.Request) {
	var inputDataArr []interface{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &inputDataArr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nextPerson := inputDataArr[rand.Intn(len(inputDataArr))].(map[string]interface{})
	w.Header().Set("Content-Type", "application/json")
	bytes, _ := json.Marshal(nextPerson)
	w.Write(bytes)
	fmt.Println("Magic 8-Ball says:", nextPerson["name"])
}
