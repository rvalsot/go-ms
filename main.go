package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server up & running at port 9090")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		helloHandler := handlers.NewHellow()
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Your severe says goodbye!")
	})

	http.ListenAndServe(":9090", nil)

}
