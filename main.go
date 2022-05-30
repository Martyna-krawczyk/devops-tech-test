package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = envWithDefault("PORT", "8080")
var gitCommit = "noData"
var version = "HEAD"

type MyApplication struct {
	Metadata []Metadata `json:"myapplication"`
}

type Metadata struct {
	Version       string `json:"version"`
	Description   string `json:"description"`
	LastCommitSha string `json:"lastcommitsha"`
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/metadata", metadataHandler)

	fmt.Printf("Starting server at port %v\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Hello World\n")
	} else {
		methodError(w, r)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "%v\n", http.StatusOK)
	} else {
		methodError(w, r)
	}
}

func metadataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		myApp := MyApplication{
			[]Metadata{{
				Version:       version,
				Description:   "pre-interview technical test",
				LastCommitSha: gitCommit,
			}},
		}
		data, err := json.Marshal(myApp)
		if err != nil {
			log.Printf("failed to marshal MyApplication struct: %s", err)
			return
		}

		fmt.Fprintf(w, "%s\n", string(data))
	} else {
		methodError(w, r)
	}
}

func methodError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, r.Method+" is not supported for this endpoint", http.StatusMethodNotAllowed)
}

func envWithDefault(key string, def string) string {
	s := os.Getenv(key)
	if s == "" {
		return def
	}
	return s
}
