package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	serverPort = ":8000"
)

type limitsResetBody struct {
	Addresses []string `json:"addresses"`
}

// StartServer ...
func StartServer() {
	log.Printf("Listening on %s\n", serverPort)
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRequest)
	mux.HandleFunc("/limits/reset", handleLimitsReset)
	http.ListenAndServe(serverPort, mux)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	addresses := r.Header.Get("X-Forwarded-For")
	IPAddr := strings.Split(addresses, ",")
	if IPAddr[0] == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Header 'X-Forwarded-For' error\n"))
		return
	}
	// subnet, err := utils.ExtractSubnet(IPAddr[0])
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte(err.Error()))
	// }

	// p := utils.Limit(strings.Split(subnet, "/")[0])
	// if !p {
	// 	http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
	// 	return
	// }
	w.Write([]byte("OK\n"))
	return
}

func handleLimitsReset(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Unmarshal
	var msg limitsResetBody
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
