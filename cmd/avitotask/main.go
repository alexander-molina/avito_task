package main

import (
	"fmt"
	"net/http"

	"github.com/alexander-molina/avito_task/internal/app/utils"
)

const (
	serverPort = ":8000"
)

func main() {
	fmt.Printf("server started at port %s\n", serverPort)
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(serverPort, nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	IPAddr := r.Header.Get("X-Forwarded-For")
	if IPAddr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Header 'X-Forwarded-For' is not set\n"))
		return
	}
	subnet, err := utils.ExtractSubnet(IPAddr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	fmt.Println(subnet)
}
