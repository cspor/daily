package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Go web app started on port: 3000")

	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutUs)

	http.ListenAndServe(":3000", nil)
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Homepage in daily")
}

func aboutUs(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "About us in daily")
}
