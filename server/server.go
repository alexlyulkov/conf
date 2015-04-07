package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/alexlyulkov/conf/conf"
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
		fmt.Fprintf(w, "Node name is not specified")
		return
	}*/
	if !conf.NameIsValid(name) {
		http.Error(w, "Name should consist only of English letters and numbers separated by dots.", 400)
		return
	}
	valueJSON := r.FormValue("value")
	if len(valueJSON) == 0 {
		http.Error(w, "Node value is not specified", 400)
		return
	}
	var value interface{}
	err := json.Unmarshal(([]byte)(valueJSON), &value)
	if err != nil {
		http.Error(w, "Node value should be proper json. Can't parse node value: "+err.Error(), 400)
		return
	}
	err = conf.CheckInterfaceConsistsOfMapsAndStrings(value)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, err.Error())
		return
	}

	path := conf.NameToPath(name)
	err = conf.CreateNode(path, value)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, "")
}

func Read(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if len(name) == 0 {
		http.Error(w, "Node name is not specified", 400)
		return
	}
	if !conf.NameIsValid(name) {
		http.Error(w, "Name should consist only of English letters and numbers separated by dots.", 400)
		return
	}

	depth := 1
	depthStr := r.FormValue("depth")
	if len(depthStr) != 0 {
		var err error
		depth, err = strconv.Atoi(depthStr)
		if err != nil {
			http.Error(w, "Depth should be a integer. "+err.Error(), 400)
			return
		}
	}
	if depth < 0 {
		depth = 2000000000
	}

	path := conf.NameToPath(name)
	value, err := conf.GetNode(path, 0, depth)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	valueJSON, err := json.Marshal(value)
	if err != nil {
		log.Panic(err)
	}

	fmt.Fprintf(w, (string)(valueJSON))
}

func Update(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	/*if len(name) == 0 {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Node name is not specified")
		return
	}*/
	if !conf.NameIsValid(name) {
		http.Error(w, "Name should consist only of English letters and numbers separated by dots.", 400)
		return
	}
	valueJSON := r.FormValue("value")
	if len(valueJSON) == 0 {
		http.Error(w, "Node value is not specified", 400)
		return
	}
	var value interface{}
	err := json.Unmarshal(([]byte)(valueJSON), &value)
	if err != nil {
		http.Error(w, "Node value should be proper json. Can't parse node value: "+err.Error(), 400)
		return
	}
	err = conf.CheckInterfaceConsistsOfMapsAndStrings(value)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	path := conf.NameToPath(name)
	err = conf.CheckSubtreeMatchesValueStructure(path, value)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = conf.UpdateNode(path, value)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Fprintf(w, "")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if len(name) == 0 {
		http.Error(w, "Node name is not specified", 400)
		return
	}
	if !conf.NameIsValid(name) {
		http.Error(w, "Name should consist only of English letters and numbers separated by dots.", 400)
		return
	}

	path := conf.NameToPath(name)
	err := conf.DeleteNode(path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Fprintf(w, "")
}
