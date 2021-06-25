package AppStats

import "time"

type Config struct {
	Version         string
	Hostname        string
	Project         string
	Environment     string
	EnvironmentType string
	UpdateFrequency time.Duration
	InfluxURL       string
}

func Go(C Config) {

	// Do some Validation on C

	go UpdateInflux(C)

}
