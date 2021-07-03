package main

import (
	"os"
	"time"

	"github.com/acky666/AppStats"
)

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

	// TODO:   Web Handler.

	// Web Server

}
