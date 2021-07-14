package main

import (
	"github.com/disturbing/github-app-k8s-secret-refresher/v2/internal/config"
	"github.com/disturbing/github-app-k8s-secret-refresher/v2/internal/github"
	"log"
	"os"
	"strconv"
)

func main() {
	config.Load()
	token, err := github.GenerateInstallationToken()

	if err != nil {
		log.Panicf("Could not generate github installation token with err: %v", token)
	}

	log.Printf("Successfully generated installation token: %s", token)

	//controller, err := controller2.New()
	// TODO: Apply it to kubernetes

}

func getEnvAsInt(envVar string) int {
	if val := os.Getenv(envVar); val != "" {
		intVal, err := strconv.Atoi(val)

		if err == nil {
			return intVal
		}

		log.Panicf("Environment variable %s is not an int", envVar)
	}

	log.Panicf("Environment variable %s is not an int", envVar)
	return 0
}
