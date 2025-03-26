package main

import (
	"fmt"
	"net"  // Add this import
	"net/http"
)

func main() {
	// Create a file server for serving static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Specify the port
	port := 3000
	
	// Create a listener to check port availability
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error listening on port %d: %v\n", port, err)
		return
	}
	defer listener.Close()

	fmt.Printf("Server is listening on port %d\n", port)
	
	// Serve using the listener
	err = http.Serve(listener, nil)
	if err != nil {
		fmt.Printf("Error serving: %v\n", err)
	}
}