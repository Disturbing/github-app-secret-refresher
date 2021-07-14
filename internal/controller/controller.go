package controller

import (
	"errors"
	"fmt"
	"github.com/disturbing/github-app-k8s-secret-refresher/v2/internal/config"
	"github.com/disturbing/github-app-k8s-secret-refresher/v2/internal/types"
)

type Controller interface {
	HandleNewToken(tokenId string)
}

func New() (Controller, error) {

	switch config.TokenProcessorType {
	case types.KUBERNETES:
		return newKubernetesController()
	default:
		return nil, errors.New(fmt.Sprintf("Unsupported token processor type %v", config.TokenProcessorType))
	}

}
