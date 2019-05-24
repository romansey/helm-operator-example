// Main package of the example service.
//
// The example service is a basic web application that
// returns a text response based on a configuration value.
// The configuration will be read from a config.json file
// in the working directory.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type config struct {
	Response string
}

var cfg *config

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, cfg.Response)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func loadConfig() error {
	fileContent, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return err
	}
	cfg = &config{}
	err = json.Unmarshal(fileContent, cfg)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := loadConfig()
	if err != nil {
		log.Println("Could not read configuration file. Using default values.")
		cfg = &config{"Default Response"}
	}
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/", rootHandler)
	log.Println("Starting service on port 8080.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
