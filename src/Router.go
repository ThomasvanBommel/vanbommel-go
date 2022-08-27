/*
 * File: Router.go
 * Date: 2022-08-26
 * Author: Thomas vanBommel
 */

package main

import (
	"regexp"
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
	for _, route := range h.routes {
		if route.pattern.MatchString(request.URL.Path) {
			route.handler.ServeHTTP(writer, request)
			return
		}
	}

	http.NotFound(writer, request)
}
