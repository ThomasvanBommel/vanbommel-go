/*
 * File: main.go
 * Date: 2022-08-18
 * Author: Thomas vanBommel
 */
package main
 
import (
	"log"
	"time"
	"net/http"
)

func main() {
	router := RegexRouter{ }
	router.AddRoute(`^/$`, func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Welcome to the home page"))
	})
	
	server := &http.Server {
		Addr: ":4000",
		Handler: &router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,	
	}
	
	log.Fatal(server.ListenAndServe())
}
