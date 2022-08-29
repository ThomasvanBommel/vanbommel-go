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
	server := &http.Server {
		Addr: ":4000",
		Handler: &RegexRouter{ },
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,	
	}
	
	log.Fatal(server.ListenAndServe())
}
