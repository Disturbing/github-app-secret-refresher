package github

import (
	"context"
	"github.com/bradleyfalzon/ghinstallation"
	"github.com/disturbing/github-app-k8s-secret-refresher/v2/internal/config"
	"net/http"
)

func GenerateInstallationToken() (string, error) {
	// Wrap the shared transport for use with the integration ID 1 authenticating with installation ID 99.
	itr, err := ghinstallation.NewKeyFromFile(
		http.DefaultTransport,
		int64(config.GithubAppId),
		int64(config.GithubAppInstallationId),
		config.GithubAppPrivateKeyFile,
	)

	if err != nil {
		return "", err
	}

	return itr.Token(context.Background())
}
