package AppStats

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"

	l "github.com/acky666/ackyLog"
)

func UpdateInflux(AppConfig Config) {

	out := ""
	l.INFO("Starting GoRoutine to Update Influx")

	// curl -v -XPOST 'http://my.influx.server:8086/write?db=db_name' --data-binary 'my_measurement,tag_name=stuff value=1'

	httpClient := http.Client{Timeout: 10 * time.Second}

	for {

		collectedStats := 0
		out = "Metrics,host=" + AppConfig.Hostname + ",Project=" + AppConfig.Project + ",Version=" + AppConfig.Version + ",EnvType=" + AppConfig.EnvironmentType + ",Env=" + AppConfig.Environment + " "
		for _, v := range Stats.GetStatKeys() {
			out = out + v + "=" + fmt.Sprintf("%d", Stats.GetStat(v))
			out = out + ","
			collectedStats++
		}
		out = strings.TrimRight(out, ",")

		if collectedStats > 0 {

			request, _ := http.NewRequest("POST", AppConfig.InfluxURL, bytes.NewBufferString(out+"\n"))
			response, err := httpClient.Do(request)
			if err != nil {

				l.ERROR("Failed to Send Statistics to Influx")
				l.ERROR("Error: %v", err)

			} else {

				if response.StatusCode != 204 {
					l.ERROR("Bad response from Influx (Server Details and Labels Correct?)")
					l.SPEW(response)
				}
			}

		}
		time.Sleep(AppConfig.UpdateFrequency)
	}
}
