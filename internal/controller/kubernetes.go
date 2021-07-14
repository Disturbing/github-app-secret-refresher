package controller

import (
	"context"
	"encoding/base64"
	"flag"
	"github.com/disturbing/github-app-k8s-secret-refresher/v2/internal/config"
	k8sMeta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sCore "k8s.io/client-go/applyconfigurations/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type kubernetesController struct {
	kubeClient                       kubernetes.Interface
	GithubAppAuthUsernameBase64Bytes []byte
}

func newKubernetesController() (Controller, error) {
	kubeConfig, err := loadKubeConfig(config.KubeConfigPath)

	if err != nil {
		return nil, err
	}

	controller := &kubernetesController{
		kubeClient: kubernetes.NewForConfigOrDie(kubeConfig),
	}

	base64.StdEncoding.Encode(controller.GithubAppAuthUsernameBase64Bytes, []byte("x-access-token"))

	return controller, nil
}

func (controller *kubernetesController) HandleNewToken(token string) {
	var passwordBase64Bytes []byte
	base64.StdEncoding.Encode(controller.GithubAppAuthUsernameBase64Bytes, []byte(token))

	controller.kubeClient.
		CoreV1().
		Secrets(config.KubeSecretNamespace).
		Apply(context.Background(), &k8sCore.SecretApplyConfiguration{
			Data: map[string][]byte{
				"username": controller.GithubAppAuthUsernameBase64Bytes,
				"password": passwordBase64Bytes,
			},
		}, k8sMeta.ApplyOptions{},
		)
}

func loadKubeConfig(kubeConfigPath string) (*rest.Config, error) {
	kubeConfigString := flag.String("kubeconfig", kubeConfigPath, "(optional) absolute path to the kubeconfig file")

	// use the current context in kubeconfig
	return clientcmd.BuildConfigFromFlags("", *kubeConfigString)
}
