default:
	@echo "Grrr..."

go.install:
	go install github.com/swaggo/swag/cmd/swag@v1.7.0
	go install github.com/go-delve/delve/cmd/dlv@v1.7.3

go.fmt:
	$(call goApps, go fmt ./...)

go.test:
	$(call goApps, go test ./...)

go.vet:
	$(call goApps, go vet ./...)

go.vendor:
	$(call goApps, go mod vendor)

go.tidy:
	$(call goApps, go mod tidy)

go.gen:
	$(call goApps, go generate ./...)

define goApps
	for app in api/src runner/src rasp/src rasp-deploy/src shared; do \
		cd "$(shell pwd)/$$app"; $(1) ; \
	done
endef
