package controller

import (
	"context"
	"encoding/base64"
	"flag"
	"github.com/disturbing/github-app-k8s-secret-refresher/v2/internal/config"
	"github.com/disturbing/github-app-k8s-secret-refresher/v2/internal/types"
	k8sMeta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sCore "k8s.io/client-go/applyconfigurations/core/v1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

type kubernetesController struct {
	kubeClient                     kubernetes.Interface
	GithubAppAuthUsernameBase64Str string
}

func newKubernetesController() (Controller, error) {
	kubeConfig, err := loadKubeConfig(config.KubeConfigPath)

	if err != nil {
		return nil, err
	}

	controller := &kubernetesController{
		kubeClient:                     kubernetes.NewForConfigOrDie(kubeConfig),
		GithubAppAuthUsernameBase64Str: base64.StdEncoding.EncodeToString([]byte(types.GitHubAppAuthUsername)),
	}

	return controller, nil
}

func (controller *kubernetesController) ProcessNewToken(token string) error {
	var passwordBase64Bytes []byte

	var (
		k8sSecretKindName    = "Secret"
		k8sSecretKindVersion = "v1"
	)

	_, err := controller.kubeClient.
		CoreV1().
		Secrets(config.KubeSecretNamespace).
		Apply(context.Background(), &k8sCore.SecretApplyConfiguration{
			// This is required for the apply method... Need to find a cleaner way!
			TypeMetaApplyConfiguration: v1.TypeMetaApplyConfiguration{
				Kind:       &k8sSecretKindName,
				APIVersion: &k8sSecretKindVersion,
			},
			ObjectMetaApplyConfiguration: &v1.ObjectMetaApplyConfiguration{
				Name: &config.KubeSecretName,
			},
			Data: map[string][]byte{
				"username": []byte(controller.GithubAppAuthUsernameBase64Str),
				"password": passwordBase64Bytes,
			},
		}, k8sMeta.ApplyOptions{
			TypeMeta: k8sMeta.TypeMeta{
				APIVersion: "v1",
			},
			FieldManager: types.AppName,
		},
		)

	if err == nil {
		log.Printf("Successfully applied github credentials token to secret %s in namespace %s",
			config.KubeSecretName,
			config.KubeSecretNamespace,
		)
	}

	return err
}

func loadKubeConfig(kubeConfigPath string) (*rest.Config, error) {
	kubeConfigString := flag.String("kubeconfig", kubeConfigPath, "(optional) absolute path to the kubeconfig file")

	// use the current context in kubeconfig
	return clientcmd.BuildConfigFromFlags("", *kubeConfigString)
}
