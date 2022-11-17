ARG BASE_IMAGE
FROM ${BASE_IMAGE}

RUN apk update
RUN apk add curl bash gettext make

RUN curl -sLO "https://dl.k8s.io/release/v1.22.0/bin/linux/amd64/kubectl"
RUN chmod +x kubectl
RUN install -m 0755 kubectl /usr/local/bin/kubectl

RUN curl -o kustomize.tar.gz -sL "https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2Fv4.4.0/kustomize_v4.4.0_linux_amd64.tar.gz"
RUN tar -xf kustomize.tar.gz
RUN chmod +x kustomize
RUN install -m 0755 kustomize /usr/local/bin/kustomize


