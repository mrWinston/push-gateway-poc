package main

import (
	"fmt"
	"log"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

const (
	PUSHGW_URL_ENV_KEY = "PUSHGW_URL"
)

func main() {
	pushgatewayUrl, ok := os.LookupEnv(PUSHGW_URL_ENV_KEY)
	if !ok {
		log.Fatalf("Env var '%s' needs to be defined.", PUSHGW_URL_ENV_KEY)
	}

	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pushgw_poc_last_completed_time",
		Help: "The timestamp of the last successful completion of the pushgw poc cronjob",
	})
	completionTime.SetToCurrentTime()

	if err := push.New(pushgatewayUrl, "db_backup").
		Collector(completionTime).
		Grouping("db", "customers").
		Push(); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}
