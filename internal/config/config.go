package config

import (
	"github.com/disturbing/github-app-k8s-secret-refresher/v2/internal/types"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	TokenProcessorType      types.TokenProcessorType
	GithubAppId             int
	GithubAppInstallationId int
	GithubAppPrivateKeyFile string

	KubeConfigPath      string
	KubeSecretNamespace string
)

func Load() {
	err := godotenv.Load()

	if err != nil {
		log.Panicf("Could not load godotenv utility with err: %v", err)
	}

	GithubAppId = getEnvAsInt("GITHUB_APP_ID")
	GithubAppInstallationId = getEnvAsInt("GITHUB_APP_INSTALLATION_ID")
	GithubAppPrivateKeyFile = os.Getenv("GITHUB_APP_PRIVATE_KEY_PATH")
	KubeConfigPath = os.Getenv("KUBE_CONFIG_PATH")
	KubeSecretNamespace = os.Getenv("KUBE_SECRET_NAMESPACE")
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
