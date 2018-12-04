package main

import "log"

var checks []Check

func main() {
	config, err := getConfigEnv()

	if err != nil {
		log.Panicf("The config file doesn't exist.")
	}

	go startCron(config)
	Server()
}
