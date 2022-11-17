.DEFAULT_GOAL=no_target

no_target:
	@echo "Bro..."

install:
	kubectl apply -f namespace.yaml
	kubectl apply -f storage.yaml

	helm repo add gitlab https://charts.gitlab.io/
	helm repo update
	helm upgrade --install --namespace gitlab -f values.yaml gitlab gitlab/gitlab --timeout 600s

upgrade:
	helm upgrade --namespace gitlab -f values.yaml gitlab gitlab/gitlab --timeout 600s

install.n:
	helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
	helm repo update

	helm install ingress-nginx ingress-nginx/ingress-nginx
