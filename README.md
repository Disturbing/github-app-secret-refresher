# GitHub App Secret Refresher

Where you can get [GitHub App Installation Access Tokens](https://docs.github.com/en/rest/reference/apps#create-an-installation-access-token-for-an-app) that are so FRESH, it hurts...

## Overview

When developing utilities for GitHub, several apps may require to use the GitHub API or require authentication
pull from a git repo on GitHub. GitHub App's private keys do not have access to pull code directly nor speak to
GitHub APIs directly and require a temporary [Installation Access Token](https://docs.github.com/en/rest/reference/apps#create-an-installation-access-token-for-an-app)
which expires every hour.

Some apps like [Flux](https://fluxcd.io/docs/) and [BuildKite](https://buildkite.com/) do not directly support
GitHub app authentication and only support SSH keys, Deploy Tokens or Personal Access Tokens to authenticate. These tools
can use a GitHub App Installation Access Token to authenticate if they can have a readily available token at all times.

Optionally, GitHub app installation tokens can be scoped to have specific permissions compared to other keys. This 
includes limiting the repositories that the token has access to and granted permissions to the repository. This provides
security so that if a service using the token is pwned, the token (1) expires within one hour and (2) is scoped to a
limited set of permissions versus an entire organization's source.

## How it works?

GitHub App Secret Refresher runs as a job on Kubernetes.

## Installation

We have a helm **v3** manifest that can be used to install the job on Kubernetes. In the root of this repository run
the following command after replacing the values specified:


```.shell
helm install -n <namespace> --name-template github-app-secret-refresher \
  --set githubAppPrivateKeyBase64=<base-64-github-app-private-key> \
  --set githubAppId=<app-id> \
  --set githubAppInstallationId=<installation-id> \
  --set jobSchedule="*/20 * * * *" 
```

The above command will schedule the job to create and update a secret called `github-credentials` within its namespace
every 20 minutes.

## Cleanup

```.shell
helm uninstall -n <namespace> github-app-secret-refresher
```

## Supported Processor Types

In the future, there may be multiple processor types such as webhooks, AWS Secret Manager, or more depending on how your
service may access the secret.

### Kubernetes Secret

Periodically creates and updates a Kubernetes Secret within the Job's namespace with the following spec.

* The **username** can be used to authenticate to a git repository on GitHub
* The **password** field is the Installation Access Token for authenticating to git or GitHub apis. 

```text
Name:         github-credentials
Namespace:    test
Labels:       <none>
Annotations:  <none>

Type:  Opaque

Data
====
password:  56 bytes
username:  20 bytes
```


# TODO List

* [ ] Remove GitHub App client and use raw API to generate a **scoped** token
* [ ] Create a **Target Namespace** so that the Job may run outside of the k8s namespace of which the secret or services live
* [ ] Be able to specify multiple target formats (oauth, custom, k8s image pull secret, etc)
* [ ] Introduce more processor types (WEBHOOK, AWS Secret Manager, etc)
