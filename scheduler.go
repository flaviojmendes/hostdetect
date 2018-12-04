package main

import (
	"fmt"

	"github.com/robfig/cron"
)

func startCron(config *Config) {
	cron := cron.New()
	cron.AddFunc(getInterval(config), func() { scanHosts(config) })
	cron.Start()
}

func getInterval(config *Config) string {
	return fmt.Sprintf("*/%v * * * * *", int(config.Interval.Seconds()))
}

func scanHosts(config *Config) {
	l, _ := scan(config)
	checks = l
}
