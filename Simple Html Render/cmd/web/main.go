package main

import (
	"SHR/pkg/handle"
	"fmt"
	"net/http"
)

const PortNumber = ":4400"

func main() {
	http.HandleFunc("/", handle.Home)
	http.HandleFunc("/about", handle.About)

	fmt.Printf("Server running and listening on port: %s\n", PortNumber)
	http.ListenAndServe(PortNumber, nil)
}
