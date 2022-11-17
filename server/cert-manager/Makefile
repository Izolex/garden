install:
	kubectl apply -f namespace.yaml
	kubectl apply -f issuer.yaml

	helm upgrade --install --namespace cert-manager -f values.yaml cert-manager jetstack/cert-manager

upgrade:
	helm upgrade --namespace cert-manager --version v1.5.3 -f values.yaml cert-manager jetstack/cert-manager
