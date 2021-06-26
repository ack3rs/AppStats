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

	// TODO: Validate C

	go UpdateInflux(C)

}
