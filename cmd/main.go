package main

import (
	"github.com/disturbing/github-app-k8s-secret-refresher/v2/internal/config"
	"github.com/disturbing/github-app-k8s-secret-refresher/v2/internal/controller"
	"github.com/disturbing/github-app-k8s-secret-refresher/v2/internal/github"
	"log"
)

func main() {
	config.Load()

	controller, err := controller.New()

	if err != nil {
		log.Panicf("Could not create token processer with err: %v", err.Error())
	}

	token, err := github.GenerateInstallationToken()

	if err != nil {
		log.Panicf("Could not generate github installation token with err: %v", err)
	}

	log.Printf("Successfully generated installation token: %s", token)

	err = controller.ProcessNewToken(token)

	if err == nil {
		log.Printf("Successfully processed token!")
	} else {
		log.Printf("An error occured while processing the token: %v", err.Error())
	}
}
