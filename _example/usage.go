package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/acky666/AppStats"
)

func WebHandler(w http.ResponseWriter, r *http.Request) {
	AppStats.Stats.IncStat("requests")
	w.Write([]byte("Hello World!"))
}

func main() {

	Hostname, _ := os.Hostname()

	AppStats.Go(AppStats.Config{Version: "v0.0.1",
		Hostname:        Hostname,
		Environment:     "development",
		EnvironmentType: "development",
		InfluxURL:       "http://influx.example.com/:8086/write?db=AppStats",
		Project:         "Example",
		UpdateFrequency: 15 * time.Second,
	})

	http.HandleFunc("/", WebHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
