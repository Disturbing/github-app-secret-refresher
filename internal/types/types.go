package types

type TokenProcessorType string

const (
	KUBERNETES TokenProcessorType = "KUBERNETES"
	// In the future, this could be Webhook, etc.
)

const (
	AppName               = "github-app-secret-refresher"
	GitHubAppAuthUsername = "x-access-token"
)
