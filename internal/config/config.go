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
	KubeSecretName      string
	KubeSecretNamespace string
)

func Load() {
	godotenv.Load()

	TokenProcessorType = types.TokenProcessorType(os.Getenv("TOKEN_PROCESSOR_TYPE"))
	GithubAppId = getEnvAsInt("GITHUB_APP_ID")
	GithubAppInstallationId = getEnvAsInt("GITHUB_APP_INSTALLATION_ID")
	GithubAppPrivateKeyFile = os.Getenv("GITHUB_APP_PRIVATE_KEY_PATH")

	KubeConfigPath = os.Getenv("KUBE_CONFIG_PATH")
	KubeSecretName = os.Getenv("KUBE_SECRET_NAME")
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
