package main

import (
	"encoding/json"
	"log"
	"net/http"
	//	"strings"
	"fmt"
	"time"
)

type Response struct {
	Value       string `json:"value"`
	IsSuccess   bool   `json:"is_success"`
	ErrorString string `json:"error_description"`
}

func StartHttpServer(address string) {

	http.HandleFunc("/messages", get_message_id)

	server := &http.Server{
		Addr:           address,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Starting http server")

	log.Panic(server.ListenAndServe())

}

func get_message_id(w http.ResponseWriter, r *http.Request) {
	/*
		var response Response
		message := r.FormValue("message")
		redirectingToMaster := r.FormValue("redirecting_to_master") == `true`
		if len(message) < 1 {
			response.ID = -1
			response.IsSuccess = false
			response.IsNew = false
			response.ErrorString = "Empty message"
		} else {
			response = GetMessageID(message, conf, m,
				cache, redirectingToMaster)
		}
		result, err := json.Marshal(response)
		if err != nil {
			fmt.Println("find_message error: ", err)
			panic("qweqw")
		}
		if conf.Machete.DebugLogs {
			log.Infof("get_message_id request returned: " + string(result))
		}
		fmt.Fprintf(w, string(result))

		WorkingRequestsMutex.Lock()
		WorkingRequests--
		WorkingRequestsMutex.Unlock()
		if rand.Int()%10 == 0 {
			Statsd.Gauge("main_db.working_requests", int64(WorkingRequests), 1.0)
		}
	*/
}

func Create(w http.ResponseWriter, r *http.Request) {
	var response Response
	name := r.FormValue("name")
	if len(name) == 0 {
		w.WriteHeader(400)
		fmt.Println(w, "Node name is not specified")
		return
	}
	if !NameIsValid(name) {
		w.WriteHeader(400)
		fmt.Println(w, "Name should consist only of English letters and numbers separated by dots.")
		return
	}
	valueJSON := r.FormValue("value")
	if len(valueJSON) == 0 {
		w.WriteHeader(400)
		fmt.Println(w, "Node value is not specified")
		return
	}
	var value interface{}
	err := json.Unmarshal(([]byte)(valueJSON), &value)
	if err != nil {
		w.WriteHeader(400)
		fmt.Println(w, "Node value should be proper json. Can't parse node value: "+err.Error())
		return
	}
	err = CheckInterfaceConsistsOfMapsAndStrings(value)
	if err != nil {
		w.WriteHeader(400)
		fmt.Println(w, err.Error())
		return
	}

	path := NameToPath(name)
	err = CreateNode(path, value)
	if err != nil {
		w.WriteHeader(400)
		fmt.Println(w, err.Error())
		return
	}

	fmt.Println(w, "")
}

func Read(w http.ResponseWriter, r *http.Request) {
	var response Response
	name := r.FormValue("name")
	if len(name) == 0 {
		w.WriteHeader(400)
		fmt.Println(w, "Node name is not specified")
		return
	}
	if !NameIsValid(name) {
		w.WriteHeader(400)
		fmt.Println(w, "Name should consist only of English letters and numbers separated by dots.")
		return
	}

	path := NameToPath(name)
	value, err := GetNode(path)
	if err != nil {
		w.WriteHeader(400)
		fmt.Println(w, err.Error())
		return
	}

	valueJSON, err := json.Marshal(value)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(w, (string)(valueJSON))
}

func Update(w http.ResponseWriter, r *http.Request) {
	var response Response
	name := r.FormValue("name")
	if len(name) == 0 {
		w.WriteHeader(400)
		fmt.Println(w, "Node name is not specified")
		return
	}
	if !NameIsValid(name) {
		w.WriteHeader(400)
		fmt.Println(w, "Name should consist only of English letters and numbers separated by dots.")
		return
	}
	valueJSON := r.FormValue("value")
	if len(valueJSON) == 0 {
		w.WriteHeader(400)
		fmt.Println(w, "Node value is not specified")
		return
	}
	var value interface{}
	err := json.Unmarshal(([]byte)(valueJSON), &value)
	if err != nil {
		w.WriteHeader(400)
		fmt.Println(w, "Node value should be proper json. Can't parse node value: "+err.Error())
		return
	}
	err = CheckInterfaceConsistsOfMapsAndStrings(value)
	if err != nil {
		w.WriteHeader(400)
		fmt.Println(w, err.Error())
		return
	}

	path := NameToPath(name)
	err = CheckSubtreeMatchesValueStructure(path, value)
	if err != nil {
		w.WriteHeader(400)
		fmt.Println(w, err.Error())
		return
	}

	err = UpdateNode(path, value)
	if err != nil {
		w.WriteHeader(400)
		fmt.Println(w, err.Error())
		return
	}

	fmt.Println(w, "")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var response Response
	name := r.FormValue("name")
	if len(name) == 0 {
		w.WriteHeader(400)
		fmt.Println(w, "Node name is not specified")
		return
	}
	if !NameIsValid(name) {
		w.WriteHeader(400)
		fmt.Println(w, "Name should consist only of English letters and numbers separated by dots.")
		return
	}

	path := NameToPath(name)
	err := DeleteNode(path)
	if err != nil {
		w.WriteHeader(400)
		fmt.Println(w, err.Error())
		return
	}

	fmt.Println(w, "")
}
