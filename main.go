package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err :=r.ParseForm(); err != nil{
		fmt.Fprintf(w,"ParseForm() err: %v",err)
	}
	name := r.FormValue("name");
	address := r.FormValue("address");
	fmt.Fprintf(w,"Name = %s\n",name);
	fmt.Fprintf(w,"Address = %s\n",address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Page not found", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Hello World")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler);
	http.HandleFunc("/form",formHandler);
	fmt.Printf("Starting server on Port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
