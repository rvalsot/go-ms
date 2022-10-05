package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server up & running at port 9090")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello from your sever!")
		dataRead, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, "There was an Oopsy", http.StatusBadRequest)
			return
		}

		fmt.Printf("Did you say? %s \n", dataRead)

		if string(dataRead) == "Ping" {
			fmt.Println("Server says: Pong!")
		}
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Your severe says goodbye!")
	})

	http.ListenAndServe(":9090", nil)

}
