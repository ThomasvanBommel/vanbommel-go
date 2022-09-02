/*
 * File: routes.go
 * Date: 2022-09-01
 * Author: Thomas vanBommel
 */

package main

import (
	"net/http"
	"html/template"
	"encoding/json"
)

type Route struct {
	Pattern string
	Handler http.Handler
}

/*** GetRoutes ****************************************************************
 * Get all routes and load their templates if necessary                       *
 ******************************************************************************/
func GetRoutes() []Route {
	return []Route{
		home(),
		assets(),
		stats(),
	}
}

/*** loadTemplate *************************************************************
 * Load a template from the file system. If an error occurs the program will  *
 * panic.                                                                     *
 ******************************************************************************/
func loadTemplate(path string) *template.Template {
	return template.Must(template.ParseFiles(path))
}

/*** home *********************************************************************
 * Home page Route using a template                                           *
 ******************************************************************************/
func home() Route {
	tmpl := loadTemplate("website/templates/home.tmpl")
	
	return Route{ "/", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Reject 'catch-all' url paths
			if r.URL.Path != "/" {
				http.NotFound(w, r)
				return
			}
			
			tmpl.Execute(w, "oil!? we struck it rich!")
		}),
	}
}

/*** assets *******************************************************************
 * An asset file-server for static files                                      *
 ******************************************************************************/
func assets() Route {
	return Route{ "/assets/", http.StripPrefix("/assets/",
		http.FileServer(http.Dir("website/assets/")),
	)}
}

/*** stats ********************************************************************
 * JSON response of the website statistics                                    *
 ******************************************************************************/
func stats() Route {
	return Route{
		"/stats",
		http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
			data, err := json.Marshal(Stats)

			if err != nil {
				http.Error(w, "error parsing statistics", 500)
				return
			}
			
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		}),
	}
}
