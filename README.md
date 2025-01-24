[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/cert-manager-webhook-bunny)](https://artifacthub.io/packages/helm/cert-manager-webhook-bunny/cert-manager-webhook-bunny)
[![Go Report Card](https://goreportcard.com/badge/github.com/davidhidvegi/cert-manager-webhook-bunny)](https://goreportcard.com/report/github.com/davidhidvegi/cert-manager-webhook-bunny)
[![License](https://img.shields.io/github/license/davidhidvegi/cert-manager-webhook-bunny)](https://github.com/davidhidvegi/cert-manager-webhook-bunny/blob/main/LICENSE)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/davidhidvegi/cert-manager-webhook-bunny)

cert-manager-webhook-bunny
===========================

[cert-manager](https://cert-manager.io) webhook implementation for use
with [Bunny](https://bunny.net) provider for solving [ACME DNS-01
challenges](https://cert-manager.io/docs/configuration/acme/dns01/).

Usage
-----

For the bunny-specific configuration, you will need to create a Kubernetes
secret, containing your customer number, API key and API password first.

You can do it like following, just place the correct values in the command:

```sh
kubectl create secret generic bunny-secret -n cert-manager --from-literal=api-key=<api-key-from-bunny-dashboard>
```
After creating the secret, configure the ``Issuer``/``ClusterIssuer`` of
yours to have the following configuration (as assumed, secret is
called "bunny-secret" and located in namespace "cert-manager"):

```yml
apiVersion: cert-manager.io/v1
kind: Issuer   # may also be a ClusterIssuer
...
spec:
    solvers:
    - dns01:
        webhook:
            groupName: com.bunny.webhook
            solverName: bunny
            config:
                secretRef: bunny-secret
                secretNamespace: cert-manager
```
For more details, please refer to https://cert-manager.io/docs/configuration/acme/dns01/#configuring-dns01-challenge-provider

Now, the actual webhook can be installed via Helm chart:
```
helm repo add cert-manager-webhook-bunny https://davidhidvegi.github.io/cert-manager-webhook-bunny/charts/

helm install my-cert-manager-webhook-bunny cert-manager-webhook-bunny/cert-manager-webhook-bunny --namespace cert-manager
```
From that point, the issuer configured above should be able to solve
the DNS01 challenges using ``cert-manager-webhook-bunny``.


Disclaimer
----------

I am in no way affiliated or associated with Bunny and this project
is done in my spare time.


License
-------

[Apache 2 License](./LICENSE)



