package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kedarnathpc/URL-Shortener/pkg/handler"
)

func main() {

	// create a new router
	r := mux.NewRouter()

	// map to search
	pathToURLs := map[string]string{
		"/linkedin": "https://www.linkedin.com/in/kedarnath-chavan-768a92226/",
		"/github":   "https://github.com/kedarnathpc",
	}

	// json file with paths and urls
	jsonFile := []byte(`
	[
  		{
    		"path": "/google",
    		"url": "https://www.google.com"
  		},
  		{
    		"path": "/go",
    		"url": "https://go.dev"
  		},
  		{	
    		"path": "/gophercises",
    		"url": "https://gophercises.com/"
  		},
  		{
    		"path": "/youtube",
    		"url": "https://www.youtube.com"
  		}
	]
	`)

	// create a handler function from the handler package
	// by passing the map and router
	mapHandler := handler.MapHandler(pathToURLs, r)

	// create a json handler
	jsonHandler, err := handler.JSONHandler(jsonFile, mapHandler)
	if err != nil {
		log.Fatal(err)
	}

	// default router
	r.HandleFunc("/", hello)

	// start the server
	fmt.Println("Starting the server at :8000...")
	http.ListenAndServe(":8000", jsonHandler)
}

// default router function
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is root")
}
