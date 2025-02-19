package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(writer http.ResponseWriter, res *http.Request) {
	if res.URL.Path != "/hello" {
		http.Error(writer, fmt.Sprint("Path ", res.URL.Path, " Not Found"), http.StatusNotFound)
	}
	if res.Method != "GET" {
		http.Error(writer, "Method not Supported", http.StatusMethodNotAllowed)
	}

	fmt.Fprintf(writer, "Hello Buddy!")
}

func formHandler(writer http.ResponseWriter, res *http.Request) {
	if err := res.ParseForm(); err != nil {
		http.Error(writer, "Cannot parse form", http.StatusInternalServerError)
	}

	fmt.Fprintf(writer, "Form Data: %v", res.Form)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server is running on port 8890")
	if err := http.ListenAndServe(":8890", nil); err != nil {
		log.Fatal(err)
	}
}
