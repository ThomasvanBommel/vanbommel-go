/*
 * File: RequestHandler.go
 * Date: 2022-08-26
 * Author: Thomas vanBommel
 */

package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
)

type RequestHandler struct {
	ResponseCount int64
	TotalResponseDuration int64
}

func router(response http.ResponseWriter, request *http.Request, handler *RequestHandler) (string, string) {
	data := ""
	contentType := "auto"
	
	switch request.URL.Path {
	case "/stats":
		average := "N/A"

		if handler.ResponseCount > 0 {
			average = time.Duration(handler.TotalResponseDuration / handler.ResponseCount).String()
		}
		
		contentType = "application/json"
		data = fmt.Sprintf(
			"{ \"ResponseCount\": %d, \"TotalResponseDuration\": %d, \"AverageResponseTime\": \"%s\" }",
			handler.ResponseCount,
			handler.TotalResponseDuration,
			average,
		)

	default:
		data = "Hi there ;)"
	}

	return data, contentType
}

func (handler *RequestHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// Get request time
	start := time.Now()
	
	// Set response headers
	response.Header().Set("X-Powered-By", "Sagittarius A*")

	// Get response data
	data, contentType := router(response, request, handler)

	// Set content type
	if contentType != "auto" {
		response.Header().Set("content-type", contentType)
	}

	// Send response
	_, err := response.Write([]byte(data))

	// Modify handler duration
	duration := time.Since(start)
	handler.ResponseCount++
	handler.TotalResponseDuration += duration.Nanoseconds()

	// Log
	str := fmt.Sprintf(
		"%s %s %s %s",
		request.RemoteAddr,
		request.Method,
		duration.String(),
		request.URL,
	)
	
	if err != nil {
		str += " WriteError: " + err.Error()
	}

	log.Println(str)
}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{
		ResponseCount: 0,
		TotalResponseDuration: 0,
	}
}
