install:
	helm repo add metallb https://metallb.github.io/metallb
	helm upgrade --install --namespace metallb metallb metallb/metallb -f values.yaml
