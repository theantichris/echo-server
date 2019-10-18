// EchoServer is an "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)

	port := getPort()
	log.Fatal(http.ListenAndServe(port, nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echos the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// getPort gets the port specified in the env or defaults to 5000.
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	return ":" + port
}
