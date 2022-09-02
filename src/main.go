/*
 * File: main.go
 * Date: 2022-08-18
 * Author: Thomas vanBommel
 */

package main
 
import (
	"log"
	"time"
	"sync"
	"net/http"
)

var Stats Statistics
var Mutex *sync.RWMutex

type RouteStats struct {
	Requests              int64
	RequestDuration       int64
	AverageDuration       int64
	AverageDurationString string
}

type Statistics struct {
	RouteStats
	Routes map[string]RouteStats
}

func (s *Statistics) AddRequest(duration time.Duration, path string) {
	Mutex.Lock()
	
	s.Requests++
	s.RequestDuration += int64(duration)

	if s.Requests > 0 {
		s.AverageDuration = s.RequestDuration / s.Requests
		s.AverageDurationString = time.Duration(s.AverageDuration).String()
	}

	if _, exists := s.Routes[path]; !exists {
		s.Routes[path] = RouteStats{ }
	}

	route := s.Routes[path]

	route.Requests++
	route.RequestDuration += int64(duration)

	if route.Requests > 0 {
		route.AverageDuration = route.RequestDuration / route.Requests
		route.AverageDurationString = time.Duration(route.AverageDuration).String()
	}

	s.Routes[path] = route
	
	Mutex.Unlock()
}

func addLoggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			start := time.Now()
			
			handler.ServeHTTP(writer, request)

			duration := time.Since(start)

			log.Println(
				request.RemoteAddr,
				request.Method,
				duration.String(),
				request.URL.Path,
			);
			
			Stats.AddRequest(duration, request.URL.Path)
		},
	)
}

/*** main *********************************************************************
 * Entry point of the program                                                 *
 ******************************************************************************/
func main() {
	Mutex = &sync.RWMutex{ }
	Stats = Statistics{ RouteStats{ }, map[string]RouteStats{ } }

	// Add each route handler to our server, also attaching the logging middleware
	for _, route := range GetRoutes() {
		http.Handle(
			route.Pattern,
			addLoggingMiddleware(route.Handler),
		)
	}

	// Start listening for incomming connections
	log.Println(http.ListenAndServe(":4000", nil))
}
