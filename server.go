package main

import (
	//"encoding/json"
	"log"
	"net/http"
	//	"strings"
	"time"
)

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
