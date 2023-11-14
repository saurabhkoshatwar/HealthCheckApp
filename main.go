package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"myhealthcheckapp/config"
	"myhealthcheckapp/dtos"
	"myhealthcheckapp/healthchecks"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <config_file_path>")
		return
	}

	configFilePath := os.Args[1]

	endpoints, err := config.ReadConfig(configFilePath)

	if err != nil {
		fmt.Println("Error reading config file: ", err)
		return
	}

	// stop channel to receive ctrl+C
	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	var wg sync.WaitGroup

	totalUrl := make(map[string]int)
	upCount := make(map[string]int)

	for _, endpoint := range endpoints {
		totalUrl[endpoint.Domain] += 1
	}

	cycleCount := 0

	for {
		select {
		case <-stopChannel:
			fmt.Println("Received CTRL+C Exiting...")
			return
		case <-ticker.C:
			cycleCount += 1

			for _, endpoint := range endpoints {
				wg.Add(1)
				go func(endpoint dtos.Endpoint) {
					defer wg.Done()
					result := healthchecks.PerformHealthCheck(endpoint)
					if result.IsUp {
						upCount[endpoint.Domain] += 1
					}
				}(endpoint)
			}

			// wait for all threads to finish execution
			wg.Wait()

			for domain, value := range upCount {
				availability := float64(value) / (float64(totalUrl[domain]) * float64(cycleCount)) * 100
				fmt.Printf("%s has %.0f%% availability percentage\n", domain, availability)
			}

		}
	}

}
