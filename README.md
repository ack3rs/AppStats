# AppStats

Small simple library to implment some application statistics and sending of these statistics to Influx for Graphing in Grafana.

Usage: 

```
	AppStats.Go(AppStats.Config{Version: "v0.0.1",
		Hostname:        "Host Name",
		Environment:     "Development",
		EnvironmentType: "Development",
		InfluxURL:       "http://my.influxdburl.com:8086/write?db=AppStats",
		Project:         "My Project Name",
		UpdateFrequency: 15 * time.Second,
	})
```

