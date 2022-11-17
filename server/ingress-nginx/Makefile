.DEFAULT_GOAL=no_target

no_target:
	@echo "Bro..."

install:
	helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
	helm repo update

	helm upgrade --install --namespace ingress-nginx ingress-nginx ingress-nginx/ingress-nginx

upgrade:
	helm upgrade --namespace ingress-nginx -f values.yaml ingress-nginx ingress-nginx/ingress-nginx
