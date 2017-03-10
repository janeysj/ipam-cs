package main

import (
	"os"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	//"strconv"
)


type Message struct {
	Id      int    `json:"id"`
	Message string `json:"-"` // Won't show that element according the task, - mean that json will skip this field
}

//Simple storage
var messages []Message

//Id of the last added message
var lastId = 0

func apiHandler() http.Handler {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/ip", assignIP).Methods("GET")
	router.HandleFunc("/ip", messagesAdd).Methods("POST")
	/* simulate delete ip address action */
	router.HandleFunc("/ip/{id}", releaseIP).Methods("GET")

	return router

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Println("sj-------home---------")
	w.Write([]byte("Gorilla!\n"))
}

//Print out the messages list or error msg if none
func assignIP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sj-------get---------")
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	ret,_ := AssignIP()
	fmt.Fprintf(w, "%s", ret)
}

//Add new message
func messagesAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//Read message, 10240 charset limit
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 10240))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	//Call function which increment the ID and add new record to the messages array
	currentMessage, err := messageAddNew(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	//Convert result to the json to be able to print result in json format according the task
	testJson, err := json.Marshal(currentMessage)

	fmt.Fprintf(w, "\n%s\n", string(testJson))

}

//Increment the ID number an add new record to the messages array
func messageAddNew(msg string) (Message, error) {
	//Initialize temp structure to be able to use append function
	tmpMessage := Message{}

	lastId += 1
	tmpMessage.Id = lastId
	tmpMessage.Message = msg

	messages = append(messages, tmpMessage)

	return tmpMessage, nil
}

//Get message by ID => /messages/ID
func releaseIP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	//Read ID number
	vars := mux.Vars(r)
	//Convert id to the digit
	ipStr := vars["id"]
	fmt.Fprintf(os.Stderr, "id is %s\n", ipStr)
	err := ReleaseIP(ipStr)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "\nId %d not found\n", ipStr)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "\n%s deleted\n", ipStr)
}

//Find if ID exist and return message if so
func messageFindById(id int) string {

	for _, m := range messages {
		if id == m.Id {
			return m.Message
		}
	}

	return ""
}

func init(){
	simulateNet()
}

func main() {

	http.ListenAndServe(":8081", apiHandler())
}

