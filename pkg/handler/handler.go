package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// map handler function to search the paths and return the corresponding url
// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathToURLs map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// extract the path from the incoming url from the request
		path := r.URL.Path

		// search for the path
		if dest, ok := pathToURLs[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		// if not found, fallback
		fallback.ServeHTTP(w, r)
	}
}

// json handler to handle the json data
func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {

	// parse the json data into a slice of structs (path, url)
	pathURLs, err := parseJSON(jsonBytes)
	if err != nil {
		log.Fatal(err)
	}

	// build a map from the slice of structs
	pathToURLs := buildMap(pathURLs)

	return MapHandler(pathToURLs, fallback), nil
}

func parseJSON(jsonBytes []byte) ([]pathURL, error) {

	// create a slice of struct
	var pathURLs []pathURL

	// unmarshal the data into the slice
	err := json.Unmarshal(jsonBytes, &pathURLs)
	if err != nil {
		log.Fatal(err)
	}

	// return the slice
	return pathURLs, nil
}

// this function can be used for any type of incoming data ex: json, yaml etc
func buildMap(pathURLs []pathURL) map[string]string {

	// make a map
	pathToURLs := make(map[string]string)

	// insert values into the map
	for _, pu := range pathURLs {
		pathToURLs[pu.Path] = pu.URL
	}

	// return map
	return pathToURLs
}

type pathURL struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}
