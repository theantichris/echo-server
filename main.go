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

	port := getPort()
	log.Fatal(http.ListenAndServe(port, nil))
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// getPort gets the port specified in the env or defaults to 5000.
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	return ":" + port
}
