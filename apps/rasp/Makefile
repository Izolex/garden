include ../.env

default:
	@echo "Grrr..."

build:
	docker run --rm \
		-v "$(shell pwd)/../rasp/src:/app" \
		-v "$(shell pwd)/../shared:/shared" \
		-v "$(shell pwd):/out" \
		-w /app/src \
		${BASE_IMAGE_GO} \
		env CC=arm-linux-gnueabihf-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 \
		go build -o /out /app

deploy.local:
	MYSQL_DSN="${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(localhost)/${MYSQL_DATABASE}" \
	SENTRY_DSN="" \
	API_SIGN_KEY="${API_SIGN_KEY}" \
	API_URL="https://todo" \
	go run ./src -ansible="$(shell pwd)/src/ansible.yaml" -bin="$(shell pwd)/main" -db="$(shell pwd)/../rasp/db.sql"
