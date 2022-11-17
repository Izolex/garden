ARG BASE_IMAGE



FROM ${BASE_IMAGE} AS source

ARG APP

WORKDIR /usr/src
COPY shared shared
COPY ${APP}/src app/src
RUN cd app/src && go mod vendor



FROM ${BASE_IMAGE} AS builder

COPY --from=source /usr/src /usr/src
RUN cd /usr/src/app/src && CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/main



FROM scratch AS prod

COPY --from=builder /usr/local/bin/main /usr/local/bin/main
COPY --from=source /usr/share/zoneinfo /usr/share/zoneinfo

ENV GIN_MODE=release

EXPOSE 80
CMD [ "/usr/local/bin/main" ]
HEALTHCHECK --interval=5s --timeout=10s CMD curl http://localhost/monitoring/status -o /dev/null -w '%{http_code}\n' -s | grep -q '200'



FROM ${BASE_IMAGE} AS dev

RUN apt-get update
RUN apt-get install -y vim iputils-ping curl inotify-tools
RUN apt-get clean -y && apt-get autoremove -y && rm -rf /var/lib/apt/lists/*
RUN go install github.com/go-delve/delve/cmd/dlv@v1.7.3

COPY --from=source /usr/src /usr/src
COPY docker/go/debug.sh /debug.sh

ENV SRC=/usr/src
ENV SRC_APP=/usr/src/app/src

WORKDIR /usr/src/app/src
EXPOSE 80 40000
CMD sh /debug.sh