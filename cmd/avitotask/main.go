package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
)

const (
	serverPort   = ":8000"
	maskBytesLen = "24"
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
	subnet, err := extractSubnet(IPAddr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	fmt.Println(subnet)
}

func extractSubnet(IPAddr string) (string, error) {
	i := bytes.IndexByte([]byte(IPAddr), '/')
	var addr string
	if i < 0 {
		addr = IPAddr
		IPAddr += "/" + maskBytesLen
	} else {
		if mask := IPAddr[i+1:]; mask != maskBytesLen {
			return "", fmt.Errorf("Provided mask: '/%s' does not match: '/%s'",
				mask, maskBytesLen)
		}
		addr = IPAddr[:i]
	}

	IP := net.ParseIP(addr)

	if IP == nil || IP.To4() == nil {
		return "", fmt.Errorf("IP address '%s' is not a valid IPv4 address", IPAddr)

	}

	_, subnet, err := net.ParseCIDR(IPAddr)
	if err != nil {
		log.Fatal(err)
	}
	return subnet.String(), nil
}
