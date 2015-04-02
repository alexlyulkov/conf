package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func StartHttpServer(address string) {

	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/read", Read)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

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

func Insert(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	/*if len(name) == 0 {
		w.WriteHeader(400)
		fmt.Println(w, "Node name is not specified")
		return
	}*/
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

	depth := 1
	depthStr := r.FormValue("depth")
	if len(depthStr) != 0 {
		var err error
		depth, err = strconv.Atoi(depthStr)
		if err != nil {
			w.WriteHeader(400)
			fmt.Println(w, "Depth should be a integer. "+err.Error())
			return
		}
	}
	if depth < 0 {
		depth = 2000000000
	}

	path := NameToPath(name)
	value, err := GetNode(path, 0, depth)
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
	name := r.FormValue("name")
	/*if len(name) == 0 {
		w.WriteHeader(400)
		fmt.Println(w, "Node name is not specified")
		return
	}*/
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
