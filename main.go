package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "POST method success")

	name := r.FormValue("Name")
	age := r.FormValue("Age")

	log.Println(name, age)

	fmt.Fprintf(w, "Name is %s\n", name)
	fmt.Fprintf(w, "Age is %s\n", age)

}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("server running at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
