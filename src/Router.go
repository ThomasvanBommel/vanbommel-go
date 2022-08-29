/*
 * File: Router.go
 * Date: 2022-08-26
 * Author: Thomas vanBommel
 */

package main

import (
	"os"
	"log"
	"time"
	"regexp"
	"strings"
	"net/http"
)

type route struct {
	pattern *regexp.Regexp
	handler  http.Handler
}

type RegexRouter struct {
	routes []*route
}

func (h *RegexRouter) AddRoute(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	h.routes = append(h.routes, &route{ regexp.MustCompile(pattern), http.HandlerFunc(handler) })
}

func (h *RegexRouter) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()
	
	callback := func(statusCode int) {
		duration := time.Since(start)
		
		log.Printf(
			"%s %s %d %s %s",
			request.RemoteAddr,
			request.Method,
			statusCode,
			duration.String(),
			request.RequestURI,
		)
	}
	
	for _, route := range h.routes {
		if route.pattern.MatchString(request.URL.Path) {
			route.handler.ServeHTTP(writer, request)
			return
		}
	}

	path := "public" + request.URL.Path

	// Add check for basenames with no extension (if none, add /index.html)

	if strings.HasSuffix(path, "/") {
		path += "index.html"
	}
	
	file, err := os.Open(path)

	defer file.Close()

	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			http.NotFound(writer, request)
			callback(404)
			return
		}

		http.Error(writer, err.Error(), 500)
		callback(500)
		return
	}
	
	http.ServeContent(writer, request, file.Name(), time.Time{ }, file)
	callback(200)
}
