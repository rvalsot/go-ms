package main

import (
	"context"
	"fmt"
	"go-ms/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	loog := log.New(os.Stdout, "product-api", log.LstdFlags)

	helloHandler := handlers.NewHellow(loog)
	goodbyeHandler := handlers.NewGoodbye(loog)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)

	// http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
	// 	log.Println("Your severe says goodbye!")
	// })

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

	// })

	servidor := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	fmt.Println("Server up & running at port 9090")
	// http.ListenAndServe(":9090", serveMux)
	go func() {
		err := servidor.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)

	signal.Notify(sigChan, os.Kill)

	sigg := <-sigChan
	log.Println("Recibí una señal, graceful shutdown", sigg)

	// Timeout context
	toContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	servidor.Shutdown(toContext)
}
